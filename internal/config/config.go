package config

const (
	// DefaultIsLambda is the program execute as a lambda function?
	DefaultIsLambda = false

	// DefaultLogLevel is the default logging level.
	// possible values: "debug", "info", "warn", "error", "fatal", "panic"
	DefaultLogLevel = "info"

	// DefaultLogFormat is the default format of the logger
	// possible values: "text", "json"
	DefaultLogFormat = "text"

	// DefaultDebug is the default debug status.
	DefaultDebug = false

	// DefaultConfigFile is the default config file name.
	DefaultConfigFile = ".aws-cwa-google-chat.yaml"

	// DefaultWebhookURL is the default incoming webhook url.
	DefaultWebhookURL = ""
)

// Config represents the configuration of the application.
type Config struct {
	ConfigFile string `mapstructure:"config-file"`
	IsLambda   bool
	Debug      bool

	LogLevel  string `mapstructure:"log_level" json:"log_level" yaml:"log_level"`
	LogFormat string `mapstructure:"log_format" json:"log_format" yaml:"log_format"`

	WebhookURL string `mapstructure:"webhook_url" json:"webhook_url" yaml:"webhook_url"`
}

// New returns a new Config
func New() Config {
	return Config{
		ConfigFile: DefaultConfigFile,
		IsLambda:   DefaultIsLambda,
		Debug:      DefaultDebug,
		LogLevel:   DefaultLogLevel,
		LogFormat:  DefaultLogFormat,
		WebhookURL: DefaultWebhookURL,
	}
}
