package models

type SendWhatsappRequest struct {
	ClientID int       `json:"client_id" validate:"required"`
	To       string    `json:"to"`
	Message  WAMessage `json:"message" validate:"required"`
}

type WAMessage struct {
	Type    string          `json:"type"`
	Content WhatsappContent `json:"content"`
}

type SendInstagramRequest struct {
	ClientID int       `json:"client_id" validate:"required"`
	To       string    `json:"to"`
	Message  WAMessage `json:"message" validate:"required"`
}

type WhatsappContent struct {
	Hsm       *Hsm         `json:"hsm,omitempty"`
	Text      string       `json:"text,omitempty"`
	Audio     *OutAudio    `json:"audio,omitempty"`
	Image     *OutImage    `json:"image,omitempty"`
	Document  *OutDocument `json:"document,omitempty"`
	Video     *OutVideo    `json:"video,omitempty"`
	Sticker   *OutSticker  `json:"sticker,omitempty"`
	TextEmail string       `json:"text_email,omitempty"`
	Name      string       `json:"name,omitempty"`
}

type BroadcastDirectRequest struct {
	Messages OutMessage `json:"messages" validate:"required"`
	ClientID int        `json:"client_id" validate:"required"`
	Channel  string     `json:"channel" validate:"required"`
}

type BroadcastDirectBulkRequest struct {
	Messages []OutMessage `json:"messages" validate:"required"`
	ClientID int          `json:"client_id" validate:"required"`
	Channel  string       `json:"channel" validate:"required"`
}
