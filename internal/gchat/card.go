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
	c.c.Header.Title = "Title"
	return nil
}

func (c *Card) SetHeader(title, subtitle string) {
	c.c.Header.Title = title
	c.c.Header.Subtitle = subtitle
}

func (c *Card) GetName() string {
	return "NotImplements"
}
