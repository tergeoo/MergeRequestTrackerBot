package model

type Config struct {
	BotToken string            `json:"bot_token"`
	ChaId    int               `json:"chat_id"`
	Devs     map[string]string `json:"developers"`
}
