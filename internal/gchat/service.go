package gchat

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/api/chat/v1"
)

var (
	ErrHTTPClientIsNil = errors.New("http client is nil")
	ErrWebhookURLIsNil = errors.New("webhookURL is nil")
	ErrCardIsNil       = errors.New("card is nil")
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type Service struct {
	client     HTTPClient
	webhookURL *WebhookURL
	card       *chat.Card
	threaded   bool
}

func NewService(client HTTPClient, webhookURL *WebhookURL, card *chat.Card, threaded bool) (*Service, error) {
	if client == nil {
		log.Error("using default http.Client")
		return nil, ErrHTTPClientIsNil
	}
	if webhookURL == nil {
		log.Fatalf("webhookURL is required")
		return nil, ErrWebhookURLIsNil
	}
	if card == nil {
		log.Error("using default card")
		return nil, ErrCardIsNil
	}

	return &Service{
		client:     client,
		webhookURL: webhookURL,
		card:       card,
		threaded:   threaded,
	}, nil
}

func (s *Service) SendCard() error {
	if s.threaded {
		s.webhookURL.SetThreadKey(s.card.Name)
	}

	payload, err := s.card.MarshalJSON()
	if err != nil {
		return err
	}

	resp, err := s.client.Post(s.webhookURL.String(), "application/json", bytes.NewReader(payload))
	if err != nil {
		log.Errorf("cannot send card: %s", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
