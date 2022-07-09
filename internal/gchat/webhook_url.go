package gchat

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

type WebhookURL struct {
	URL        *url.URL
	SpaceID    string
	APIVersion string
	Key        string
	Token      string
}

func NewWebhookURL(u *url.URL) (*WebhookURL, error) {
	s := strings.Split(u.Path, "/")

	r := regexp.MustCompile(`v\d{1,2}`)
	apiVersion := s[1]
	if !r.MatchString(apiVersion) {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, unsupported API Version: %s", apiVersion)
	}

	if len(s) < 3 {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, missing spaces")
	}
	if s[2] != "spaces" {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, unsupported space: %s", s[2])
	}

	r = regexp.MustCompile(`spaces/(.*?)/messages`)
	groups := r.FindStringSubmatch(u.Path)
	if len(groups) < 1 {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, invalid spaceID")
	}
	spaceID := groups[1]
	if spaceID == "" {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, spaceID is empty")
	}

	r = regexp.MustCompile(`/messages/$`)
	if !r.MatchString(u.Path) {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, url path doesn't contains 'messages': %s", u.Path)
	}

	key := u.Query().Get("key")
	if key == "" {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, key is empty")
	}

	token := u.Query().Get("token")
	if token == "" {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, token is empty")
	}

	return &WebhookURL{
		URL:        u,
		APIVersion: apiVersion,
		SpaceID:    spaceID,
		Key:        key,
		Token:      token,
	}, nil
}

func (w *WebhookURL) GetURL() string {
	return w.URL.String()
}

func (w *WebhookURL) GetQuery() string {
	return w.URL.RawQuery
}

func (w *WebhookURL) GetSpaceID() string {
	return w.SpaceID
}

func (w *WebhookURL) GetKey() string {
	return w.Key
}

func (w *WebhookURL) GetToken() string {
	return w.Token
}
