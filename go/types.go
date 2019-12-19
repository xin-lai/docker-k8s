package main

//文本消息
type TextWebhook struct {
	Msgtype string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}

type Text struct {
	Content string `json:"content"`
}

//@，支持@手机号码和所有人
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool     `json:"isAtAll"`
}

//markdown消息
type MarkdownWebHook struct {
	Msgtype  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}

type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
