package gchat

import (
	"io"

	"google.golang.org/api/chat/v1"
)

type Event interface{}

type Card struct {
	c *chat.Card
}

func NewCard(e Event) *Card {
	return &Card{
		c: &chat.Card{},
	}
}

func (c *Card) Render() io.Reader {
	return nil
}
