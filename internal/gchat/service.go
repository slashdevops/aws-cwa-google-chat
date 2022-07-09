package gchat

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type Service struct {
	client     HTTPClient
	webhookURL *WebhookURL
	card       *Card
	threaded   bool
}

func NewService(client HTTPClient, webhookURL *WebhookURL, card *Card, threaded bool) *Service {
	if client == nil {
		log.Info("using default http.Client")
		client = &http.Client{}
	}
	if card == nil {
		log.Info("using default card")
		card = NewCard(nil)
	}
	if webhookURL == nil {
		log.Fatalf("webhookURL is required")
	}

	return &Service{
		client:     client,
		webhookURL: webhookURL,
		card:       card,
		threaded:   threaded,
	}
}

func (s *Service) SendCard() error {
	if s.threaded {
		s.webhookURL.SetThreadKey(s.card.GetName())
	}

	resp, err := s.client.Post(s.webhookURL.String(), "application/json", s.card.Render())
	if err != nil {
		log.Errorf("cannot send card: %s", err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
