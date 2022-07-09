package gchat

import "google.golang.org/api/chat/v1"

type Message interface{}

type Card struct {
	c *chat.Card
}

func NewCard(m Message) *Card {
	return &Card{
		c: &chat.Card{},
	}
}

func (c *Card) Render() error {
	return nil
}
