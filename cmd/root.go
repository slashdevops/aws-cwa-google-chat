/*
Copyright © 2022 Christian González Di Antonio <christiangda@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/slashdevops/aws-cwa-google-chat/internal/config"
	"github.com/slashdevops/aws-cwa-google-chat/internal/event"
	"github.com/slashdevops/aws-cwa-google-chat/internal/gchat"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg config.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-cwa-google-chat",
	Short: "Send AWS CloudWatch Alarms to Google Chat using webhooks",
	Long: `This application could be used to send AWS CloudWatch Alarms messages
read from AWS SNS Topic to a Google Chat room using incoming webhooks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return handlerRequest(context.Background(), json.RawMessage(args[0]))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if cfg.IsLambda {
		lambda.Start(rootCmd.Execute)
	}
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cfg = config.New()
	cfg.IsLambda = len(os.Getenv("_LAMBDA_SERVER_PORT")) > 0

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfg.ConfigFile, "config-file", "c", config.DefaultConfigFile, "configuration file")
	rootCmd.PersistentFlags().BoolVarP(&cfg.Debug, "debug", "d", config.DefaultDebug, "fast way to set the log-level to debug")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogFormat, "log-format", "f", config.DefaultLogFormat, "set the log format")
	rootCmd.PersistentFlags().StringVarP(&cfg.LogLevel, "log-level", "l", config.DefaultLogLevel, "set the log level [panic|fatal|error|warn|info|debug|trace]")

	rootCmd.PersistentFlags().StringVarP(&cfg.WebhookURL, "webhook-url", "u", config.DefaultWebhookURL, "incoming Webhook URL from Google Chat Space")

	rootCmd.PersistentFlags().BoolVarP(
		&cfg.UseChatThreads,
		"use-chat-threads",
		"t",
		config.DefaultUseChatThreads,
		"create a thread for each alarm, you will see the alarms status in the same thread",
	)
}

// initConfig reads in config file and ENV variables.
func initConfig() {
	viper.SetEnvPrefix("cwagc") // allow to read in from environment

	envVars := []string{
		"log_level",
		"log_format",
		"debug",
		"webhook_url",
		"use_chat_threads",
	}

	for _, e := range envVars {
		if err := viper.BindEnv(e); err != nil {
			log.Panicf(errors.Wrap(err, "cannot bind environment variable").Error())
		}
	}

	// when use a lambda, we need to read the config from the environment variables only
	if !cfg.IsLambda {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)

		currentDir, err := os.Getwd()
		cobra.CheckErr(err)
		viper.AddConfigPath(currentDir)

		fileDir := filepath.Dir(cfg.ConfigFile)
		viper.AddConfigPath(fileDir)

		// Search config in home directory with name "downloader" (without extension).
		fileExtension := filepath.Ext(cfg.ConfigFile)
		fileExtensionName := fileExtension[1:]
		viper.SetConfigType(fileExtensionName)

		fileNameExt := filepath.Base(cfg.ConfigFile)
		fileName := fileNameExt[0 : len(fileNameExt)-len(fileExtension)]
		viper.SetConfigName(fileName)

		log.WithFields(log.Fields{
			"directory": fileDir,
			"file":      fileName,
			"extension": fileExtensionName,
		}).Debug("configuration file")

		if err := viper.ReadInConfig(); err == nil {
			log.WithFields(log.Fields{"file": viper.ConfigFileUsed()}).Info("using config file")
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Panicf(errors.Wrap(err, "cannot unmarshal config").Error())
	}

	switch strings.ToLower(cfg.LogFormat) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	default:
		log.Warnf("unknown log format: %s, using text", cfg.LogFormat)
		log.SetFormatter(&log.TextFormatter{})
	}

	if cfg.Debug {
		cfg.LogLevel = "debug"
	}

	// set the configured log level
	if level, err := log.ParseLevel(strings.ToLower(cfg.LogLevel)); err == nil {
		log.SetLevel(level)
	}

	if cfg.WebhookURL == "" {
		log.Panic("webhook-url is required")
	}

	u, err := url.Parse(cfg.WebhookURL)
	if err != nil {
		log.Panicf("invalid webhook-url: %s", err)
	}
	cfg.ChatWebhookURL = u
}

func handlerRequest(ctx context.Context, raw json.RawMessage) error {
	var err error
	var snsEvent events.SNSEvent
	var cwaEvent events.CloudWatchEvent

	log.WithFields(
		log.Fields{
			"functionName":    ctx.Value("FunctionName"),
			"functionVersion": ctx.Value("FunctionVersion"),
			"request":         string(raw),
		},
	).Debug("handlerRequest")

	if err = json.Unmarshal(raw, &snsEvent); err == nil {
		ev, e := event.NewSNSAlarm(&snsEvent)
		if e != nil {
			log.Errorf("error creating SNS Alarm event: %s", e)
			return e
		}
		return handleEventRequest(ev)
	}

	if err = json.Unmarshal(raw, &cwaEvent); err == nil {
		ev, e := event.NewSNSCloudWatchEvent(&cwaEvent)
		if e != nil {
			log.Errorf("error creating SNS CloudWatch Event: %s", e)
			return e
		}
		return handleEventRequest(ev)
	}

	if err != nil {
		log.Errorf("unable to marshall event: %s, error: %s", string(raw), err)
	}

	return err
}

func handleEventRequest(e event.Eventer) error {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 5
	retryClient.RetryWaitMin = time.Second * 1
	retryClient.RetryWaitMax = time.Second * 4

	if cfg.Debug {
		retryClient.Logger = log.StandardLogger()
	} else {
		retryClient.Logger = nil
	}

	httpClient := retryClient.StandardClient()

	cardHeader := gchat.CardHeaderBuilder().
		WithTitle("CloudWatch Alarm").
		WithSubtitle("CloudWatch Alarm").
		WithImageURL("https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png").
		WithImageStyle("IMAGE").
		Build()

	cardSections := gchat.CardSectionBuilder().
		WithHeader("This is a section header").
		Build()

	card := gchat.CardBuilder(e.GetSource()).
		WithHeader(cardHeader).
		WithSections(cardSections).
		Build()

	w, err := gchat.NewWebhookURL(cfg.ChatWebhookURL)
	if err != nil {
		log.Errorf("cannot create webhook: %s", err)
		return err
	}

	s, err := gchat.NewService(httpClient, w, card, cfg.UseChatThreads)
	if err != nil {
		log.Errorf("cannot create service: %s", err)
		return err
	}

	return s.SendCard()
}
