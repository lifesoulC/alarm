package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log"
	"net/smtp"
	"strings"
)

type Message struct {
	Receiver string
	Msg      []Req
}

type MessageByHost map[string][]AlarmMsg

const ContentType = "Content-Type: text/html; charset=UTF-8"

/*********************************************************************
* Function: 将处理好的数据以邮件形式发出
* Description:
*	syntax --
*          (s *SenderType) Send(Packet map[string]string)
* 	input :
*				Packet map[string]string   存放数据信息的map
* 	return value :
* Designer:
**************************************************************************/
func (s *SenderType) Send(Packet map[string][]Req) {
	var msg Message
	fmt.Println("--------------------------------------------")
	for msg.Receiver, msg.Msg = range Packet {
		fmt.Printf("Email: %v\nMsg: %v\n", msg.Receiver, msg.Msg)
	}
	for msg.Receiver, msg.Msg = range Packet {
		mailServer := strings.Split(s.MailHost, ":")
		auth := smtp.PlainAuth("", s.User, s.Password, mailServer[0])

		MsgToSend := s.MakeMessage(msg.Msg)
		sendto := strings.Split(msg.Receiver, ";")
		err := smtp.SendMail(s.MailHost, auth, s.User, sendto, MsgToSend)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("send success")
		fmt.Println("--------------------------------------------")
	}
}

/*及时发送邮件*/
func (sender *SenderType) SendNow(r *Req) {
	sendmap := make(PacketByName)
	rslice := []Req{*r}
	sendmap[r.Hoster.Email] = rslice
	go sender.Send(sendmap)
}

/*********************************************************************
* Function: 读取json配置数据（发信人信息）
* Description:
*	syntax --
*          eadSenderInfo() (*SenderType, error)
* 	input :
*
* 	return value : 发送人信息结构
* Designer:
**************************************************************************/
func ReadSenderInfo() (*SenderType, error) {
	sender := &SenderType{}

	bytes, err := ioutil.ReadFile("sender.json")
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, sender); err != nil {
		return nil, err
	}
	return sender, nil
}

func (s *SenderType) MakeMessage(msg []Req) []byte {
	hostmap := make(MessageByHost)
	for _, v := range msg {
		for _, b := range v.Content {
			hostmap[v.Hoster.HostName] = append(hostmap[v.Hoster.HostName], b)
		}
	}

	var hostlist string
	var MsgBody string
	MsgBody += "每个机房的报警信息：<br/>"
	for host, alarmMsg := range hostmap {
		hostlist += host + ","
		MsgBody += host + "<br/>"
		for _, s := range alarmMsg {
			MsgBody += s.MsgName + " : " + s.MsgValue + "<br/>"
		}
		MsgBody += "<br/>"
	}

	tmphostlist := []byte(hostlist)
	tmphostlist = tmphostlist[:len(tmphostlist)-1]
	hostlist = string(tmphostlist)

	receiver := "To: " + msg[0].Hoster.Email + "\r\n"
	sender := "From: " + s.User + "\r\n"
	subject := "Subject: " + hostlist + "主机报警！" + "\r\n"

	fmt.Println(receiver + sender + subject + ContentType + "\r\n\r\n" + MsgBody)

	MsgToSend := []byte(receiver + sender + subject + ContentType + "\r\n\r\n" + MsgBody)
	return MsgToSend
}
