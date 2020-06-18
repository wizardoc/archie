package services

import (
	"archie/robust"
	"fmt"
)

type Message struct {
	From  string
	To    string
	Title string
	Body  string
}

type InviteMessagePayload struct {
	OrganizeName string `json:"organizeName"`
	InviteToken  string `json:"inviteToken"`
}

func wrapError(err error) error {
	if err != nil {
		return robust.MESSAGE_SEND_FAILURE
	}

	return nil
}

func publish(channelMsg *ChannelMessage) error {
	return wrapError(NewPublisher(channelMsg).Publish())
}

func (msg *Message) SendPersonalMessage() error {
	channelMsg, err := NewChannelMessage(msg.From, msg.From, []string{msg.To}, DIRECTIONAL, NOTIFY, PERSONAL, msg.Title, msg.Body, nil)

	if err != nil {
		return err
	}

	return publish(channelMsg)
}

func (msg *Message) SendInviteMessage(inviteToken string, inviteUsername string, organizeName string) error {
	msg.Title = "邀请通知"
	msg.Body = fmt.Sprintf("你好，%s 邀请您加入 %s 组织，点击下方按钮接受邀请，1 小时内有效哦。", inviteUsername, organizeName)
	channelMsg, err := NewChannelMessage(msg.From, msg.From, []string{msg.To}, DIRECTIONAL, NOTIFY, INVITE, msg.Title, msg.Body, InviteMessagePayload{
		OrganizeName: organizeName,
		InviteToken:  inviteToken,
	})

	if err != nil {
		return err
	}

	return publish(channelMsg)
}
