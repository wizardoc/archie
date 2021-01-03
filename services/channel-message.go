package services

import (
	"archie/utils"
	"errors"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	// type of send way
	BROADCAST = iota
	DIRECTIONAL
)

const (
	// type of message
	CHAT = iota
	NOTIFY
)

const (
	SYSTEM = iota
	PERSONAL
	INVITE
)

type ChannelMessageMain struct {
	Title   string      `json:"title"`
	Body    string      `json:"body"`
	Payload interface{} `json:"payload"`
}

type ChannelMessage struct {
	ID            string              `json:"id"`
	Owner         string              `json:"owner"`
	Type          int                 `json:"type"`
	Tag           int                 `json:"tag"`
	From          string              `json:"from"`
	To            []string            `json:"users"`
	SendTime      string              `json:"sendTime"`
	IsRead        bool                `json:"isRead"`
	IsDelete      bool                `json:"isDelete"`
	MessageType   int                 `json:"messageType"`
	Main          *ChannelMessageMain `json:"main"`
	DirectionType int
}

func NewChannelMessage(owner string, from string, to []string, sendType int, messageType int, tag int, title string, body string, payload interface{}) (*ChannelMessage, error) {
	msgID := uuid.NewV4().String()

	if err := validType(sendTypes(), sendType, "This sendType is invalid"); err != nil {
		return nil, err
	}

	if err := validType(msgTypes(), messageType, "This msgType is invalid"); err != nil {
		return nil, err
	}

	return &ChannelMessage{
		ID:          msgID,
		Owner:       owner,
		From:        from,
		To:          to,
		Tag:         tag,
		Type:        sendType,
		MessageType: messageType,
		SendTime:    time.Now().String(),
		IsRead:      false,
		Main: &ChannelMessageMain{
			Title:   title,
			Body:    body,
			Payload: payload,
		},
	}, nil
}

func (message *ChannelMessage) valid() (err error) {
	if message.Main == nil {
		err = errors.New("The body of message is not be a nil")
	} else if message.ID == "" {
		err = errors.New("The ID of message is not be a empty str")
	} else if message.From == "" {
		err = errors.New("You have to specify a sender")
	}

	return
}

func validType(arr []int, t int, errMsg string) (err error) {
	if utils.ArrayIncludes(arr, t) {
		err = errors.New(errMsg)
	}

	return nil
}

func msgTypes() []int {
	return []int{CHAT, NOTIFY}
}

func sendTypes() []int {
	return []int{BROADCAST, DIRECTIONAL}
}
