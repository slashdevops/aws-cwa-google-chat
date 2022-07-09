package gchat

import (
	"fmt"
	"net/url"
	"regexp"
)

type WebhookURL struct {
	URL        *url.URL
	SpaceID    string
	APIVersion string
	Key        string
	Token      string
}

func NewWebhookURL(u *url.URL) (*WebhookURL, error) {
	// https://regex101.com/r/ifAT7C/1
	pathRegex := regexp.MustCompile(`(\w{2,})`)
	regexGroups := pathRegex.FindAllString(u.Path, -1)

	if len(regexGroups) < 4 || len(regexGroups) > 5 {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, URL Path invalid")
	}

	r := regexp.MustCompile(`v\d{1,2}`)
	apiVersion := regexGroups[0]
	if !r.MatchString(apiVersion) {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, unsupported API Version: %s", apiVersion)
	}

	if regexGroups[1] != "spaces" {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, unsupported space: %s", regexGroups[1])
	}

	r = regexp.MustCompile(`/messages/$`)
	if !r.MatchString(u.Path) {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, url doesn't contains 'messages': %s", u.Path)
	}

	r = regexp.MustCompile(`spaces/(.*?)/messages`)
	spaceIDGroups := r.FindStringSubmatch(u.Path)
	if len(spaceIDGroups) < 2 {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, invalid spaceID")
	}
	spaceID := spaceIDGroups[1]
	if spaceID == "" || spaceID != regexGroups[2] || len(spaceID) < 5 || len(spaceID) > 20 {
		return nil, fmt.Errorf("invalid Google Chat Webhook URL, spaceID is empty or invalid, spaceID: %s", spaceID)
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

func (w *WebhookURL) String() string {
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
