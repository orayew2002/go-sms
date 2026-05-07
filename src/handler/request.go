package handler

type SendMessageRequqest struct {
	From     string `json:"from"`
	To       string `json:"to" binding:"required"`
	TextType string `json:"text_type"`
	Text     string `json:"text"`
	APIKey   string `json:"api_key"`
}
