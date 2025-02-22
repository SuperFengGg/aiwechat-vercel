package config

import "os"

const (
	Grok_Welcome_Reply_Key = "gptWelcomeReply"
	Grok_Token             = "GROK_TOKEN"
)

func GetGrokWelcomeReply() (r string) {
	r = os.Getenv(Grok_Welcome_Reply_Key)
	if r == "" {
		r = "我是grok，开始聊天吧！"
	}
	return
}

func GetGroktToken() string {
	return os.Getenv(Grok_Token)
}
