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
	httpClient HTTPClient
	webhookURL string
	card       *Card
}

func NewService(httpClient HTTPClient, webhookURL string, card *Card) *Service {
	if httpClient == nil {
		log.Info("using default http.Client")
		httpClient = &http.Client{}
	}
	if card == nil {
		log.Info("using default card")
		card = NewCard(nil)
	}
	if webhookURL == "" {
		log.Fatalf("webhookURL is required")
	}

	return &Service{
		httpClient: httpClient,
		webhookURL: webhookURL,
		card:       card,
	}
}

func (s *Service) Send() error {
	return nil
}
