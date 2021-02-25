package chat

import "testing"

func TestChatSendFirstMessage(t *testing.T) {
	chat := NewChat(8)
	msg := NewMessage("john_doe67", "hi")

	chat.Send(msg)

	got := chat.numMessages
	want := 1

	if got != want {
		t.Errorf("Chat.Send(Message) did not increase the message number. Got: '%d' but wanted '%d'", got, want)
	}

	got = chat.end
	want = 0

	if got != want {
		t.Errorf("Chat.Send(Message) did not increase 'end' pointer. Got: '%d' but wanted '%d'", got, want)
	}
}

func TestChatSendMessageDeletionWhenLimitIsReached(t *testing.T) {
	chat := NewChat(1)

	message1 := NewMessage("john_doe67", "hi")
	message2 := NewMessage("john_doe67", "bye")

	chat.Send(message1)
	chat.Send(message2)

	got := chat.Messages[0]
	want := message2

	if got != want {
		t.Errorf("Chat.Send(Message) did not put the message at the appropriate place. Got: '%v' but wanted '%v'", got, want)
	}
}

func TestChatSendStartAndEndPointerWhenLimitIsReached(t *testing.T) {
	chat := NewChat(4)

	messages := []Message{
		NewMessage("john_doe67", "hi"),
		NewMessage("JohnDoeZX9", "hi, how are you?"),
		NewMessage("john_doe67", "actualy not very good. I will logout now."),
		NewMessage("JohnDoeZX9", "ok, see you later"),
		NewMessage("john_doe67", "see ya"),
	}

	for _, msg := range messages {
		chat.Send(msg)
	}

	got := chat.start
	want := 1

	if got != want {
		t.Errorf("Chat.Start do not have the expected value. Got: '%d' but wanted '%d'", got, want)
	}

	got = chat.end
	want = 0

	if got != want {
		t.Errorf("Chat.End do not have the expected value. Got: '%d' but wanted '%d'", got, want)
	}
}

func TestChatGetMessages(t *testing.T) {
	chat := NewChat(4)

	messagesSource := []Message{
		NewMessage("john_doe67", "hi"),
		NewMessage("JohnDoeZX9", "hi, how are you?"),
		NewMessage("john_doe67", "actualy not very good. I will logout now."),
		NewMessage("JohnDoeZX9", "ok, see you later"),
		NewMessage("john_doe67", "see ya"),
	}

	for _, msg := range messagesSource {
		chat.Send(msg)
	}

	messages := chat.GetMessages()

	isSomeMessageDifferent := false
	for idx, msg := range messages {
		isSomeMessageDifferent = isSomeMessageDifferent || msg != messagesSource[5-idx-1]
	}

	if isSomeMessageDifferent {
		correctMessages := []Message{
			messagesSource[4],
			messagesSource[3],
			messagesSource[2],
			messagesSource[1],
		}
		t.Errorf("Chat.GetMessages() returns messages in wrong order. Got: %v but wanted: %v", messages, correctMessages)
	}
}
