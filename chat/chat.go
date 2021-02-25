package chat

//Chat struct holds data about the server chat
type Chat struct {
	MaxMessages int
	Messages    []Message
	start       int
	numMessages int
	end         int
}

//NewChat creates and returns a Chat object with maxMessages capacity
func NewChat(maxMessages int) Chat {
	var chat Chat
	chat.MaxMessages = maxMessages

	chat.Messages = make([]Message, chat.MaxMessages)
	chat.start = 0
	chat.end = -1
	chat.numMessages = 0

	return chat
}

//Send adds a message in the chat and, if the message limit is reached, deletes the last message.
func (c *Chat) Send(message Message) {
	c.end = (c.end + 1) % c.MaxMessages
	c.Messages[c.end] = message
	c.numMessages++
	if c.numMessages > c.MaxMessages {
		c.numMessages = c.MaxMessages
		c.start = (c.start + 1) % c.MaxMessages
	}

}

//GetMessages returns a slice containing the chat messages in order
func (c *Chat) GetMessages() []Message {
	messageList := make([]Message, c.numMessages)

	for i := 0; i < c.numMessages; i++ {
		idx := (c.start + i) % c.MaxMessages
		messageList[c.numMessages-i-1] = c.Messages[idx]
	}

	return messageList
}
