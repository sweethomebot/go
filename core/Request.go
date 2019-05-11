package core

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
	"io"
	"strings"
	"encoding/json"
)

//go:generate go get github.com/google/uuid

type RequestMessage struct {
	FromPlugin string
	ToPlugin   string
	SendAnswer bool
	UsrId      string
	Cmd        string
	Data       MessageData
	Answer     chan MessageData
}

type EventMessage struct {
	FromPlugin string
	EventName  string
	Data       MessageData
}

type Channel struct {
	EventIn  chan EventMessage
	EventOut chan EventMessage

	// Anfragen von anderen
	RequestIn chan RequestMessage

	// selbst gestellte Anfragen
	RequestOut chan RequestMessage
}

func NewChannel() Channel {
	var ch = Channel{}
	ch.EventIn = make(chan EventMessage)
	ch.EventOut = make(chan EventMessage)
	ch.RequestIn = make(chan RequestMessage)
	ch.RequestOut = make(chan RequestMessage)
	return ch
}

func UUID() string {
	u, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("UUID Error: ", err)
	}
	return u.String()
}

func MD5(data string) string {
	h := md5.New()
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func JsonEncode(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JsonEncode Error:", err)
	}
	return string(jsonStr)
}
func JsonDecode(data string, v interface{})  {
	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		fmt.Println("JsonDecode Error:", err, data)
	}
}

func EscapeString(data string) string {
	// http://php.net/manual/de/function.mysql-real-escape-string.php
	data = strings.Replace(data, "\\", "\\\\", -1)
	data = strings.Replace(data, "\x00", "\\\x00", -1)
	data = strings.Replace(data, "\x1a", "\\\x1a", -1)
	data = strings.Replace(data, "\n", "\\\n", -1)
	data = strings.Replace(data, "\r", "\\\r", -1)
	data = strings.Replace(data, "'", "''", -1)
	//data = strings.Replace(data, "\"", "\\\"", -1)
	return data
}
func EscapeStringMap(data map[string]string) map[string]string {
	var m = make(map[string]string)
	for k, v := range data {
		m[EscapeString(k)] = EscapeString(v)
	}
	return m
}

func SendRequest(ch Channel, toPlugin string, cmd string, data MessageData) {
	var msg = RequestMessage{}
	msg.ToPlugin = toPlugin
	msg.Cmd = cmd
	msg.Data = data
	msg.SendAnswer = false
	ch.RequestOut <- msg
}

func SendRequestAndWait(ch Channel, toPlugin string, cmd string, data MessageData) MessageData {
	return SendRequestAndWaitUsrId(ch, toPlugin, cmd, data, "")
}

func SendRequestAndWaitUsrId(ch Channel, toPlugin string, cmd string, data MessageData, usrId string) MessageData {
	var msg = RequestMessage{}
	var answer = make(chan MessageData)
	msg.ToPlugin = toPlugin
	msg.Cmd = cmd
	msg.UsrId = usrId
	msg.Data = data
	msg.SendAnswer = true
	msg.Answer = answer
	ch.RequestOut <- msg

	// wait for answer
	buffer := <-answer
	close(answer)
	return buffer
}

func SendEvent(ch Channel, eventName string, data MessageData) {
	var event = EventMessage{}
	event.EventName = eventName
	event.Data = data
	ch.EventOut <- event
}

func ListenForAllEvents(ch Channel, pluginName string) {
	SendRequest(ch, "system", "listenForAllEvents", CreText(pluginName))
}

func ListenForEvent(ch Channel, pluginName, eventName string) {
	SendRequest(ch, "system", "listenForEvent", CreMap(map[string]string{"pluginName":pluginName, "eventName":eventName}))
}

func ClearEventListeners(ch Channel, pluginName string) {
	SendRequest(ch, "system", "clearEventListeners", CreText(pluginName))
}
