package chat

//Message represents a specific message in the chat
type Message struct {
	Author  string
	Content string
}

//NewMessage creates and returns a new Message object based in the provided arguments
func NewMessage(author string, content string) Message {
	var message Message
	message.Author = author
	message.Content = content

	return message
}
