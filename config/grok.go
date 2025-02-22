package config

import "os"

const (
	Gpt_Welcome_Reply_Key = "gptWelcomeReply"
	Gpt_Token             = "GROK_TOKEN"
)

func GetGptWelcomeReply() (r string) {
	r = os.Getenv(Gpt_Welcome_Reply_Key)
	if r == "" {
		r = "我是grok，开始聊天吧！"
	}
	return
}

func GetGptToken() string {
	return os.Getenv(Gpt_Token)
}
