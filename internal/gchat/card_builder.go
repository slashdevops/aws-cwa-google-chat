package gchat

import "google.golang.org/api/chat/v1"

type cardSectionBuilder struct {
	sections *chat.Section
}

func CardSectionBuilder() *cardSectionBuilder {
	return &cardSectionBuilder{
		sections: &chat.Section{},
	}
}

func (s *cardSectionBuilder) WithHeader(header string) *cardSectionBuilder {
	s.sections.Header = header
	return s
}

func (s *cardSectionBuilder) WithWidgets(widgets ...*chat.WidgetMarkup) *cardSectionBuilder {
	s.sections.Widgets = widgets
	return s
}

func (s *cardSectionBuilder) Build() *chat.Section {
	return s.sections
}

type cardActionsBuilder struct {
	actions []*chat.CardAction
}

func CardActionsBuilder() *cardActionsBuilder {
	return &cardActionsBuilder{
		actions: []*chat.CardAction{},
	}
}

func (a *cardActionsBuilder) WithAction(action *chat.CardAction) *cardActionsBuilder {
	a.actions = append(a.actions, action)
	return a
}

func (a *cardActionsBuilder) Build() []*chat.CardAction {
	return a.actions
}

type cardHeaderBuilder struct {
	header *chat.CardHeader
}

func CardHeaderBuilder() *cardHeaderBuilder {
	return &cardHeaderBuilder{
		header: &chat.CardHeader{},
	}
}

func (h *cardHeaderBuilder) WithTitle(title string) *cardHeaderBuilder {
	h.header.Title = title
	return h
}

func (h *cardHeaderBuilder) WithSubtitle(subtitle string) *cardHeaderBuilder {
	h.header.Subtitle = subtitle
	return h
}

func (h *cardHeaderBuilder) WithImageStyle(imageStyle string) *cardHeaderBuilder {
	h.header.ImageStyle = imageStyle
	return h
}

func (h *cardHeaderBuilder) WithImageURL(imageURL string) *cardHeaderBuilder {
	h.header.ImageUrl = imageURL
	return h
}

func (h *cardHeaderBuilder) Build() *chat.CardHeader {
	return h.header
}

type cardBuilder struct {
	card *chat.Card
}

func CardBuilder(name string) *cardBuilder {
	return &cardBuilder{
		card: &chat.Card{
			Name: name,
		},
	}
}

func (c *cardBuilder) WithHeader(header *chat.CardHeader) *cardBuilder {
	c.card.Header = header
	return c
}

func (c *cardBuilder) WithSections(sections ...*chat.Section) *cardBuilder {
	c.card.Sections = sections
	return c
}

func (c *cardBuilder) Build() *chat.Card {
	return c.card
}
