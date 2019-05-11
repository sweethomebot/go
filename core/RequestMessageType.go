package core

import (
	"encoding/json"
	"fmt"
)

type MessageType int32

const (
	MessageTypeNull      = MessageType(0)
	MessageTypeText      = MessageType(1)
	MessageTypeBool      = MessageType(2)
	MessageTypeMap       = MessageType(3)
	MessageTypeArrMap    = MessageType(4)
	MessageTypeArrString = MessageType(5)
	MessageTypeHttp      = MessageType(6)
)

type MessageData struct {
	Type MessageType
	Data interface{}
}
type MessageDataHttp struct {
	Args []string
	Data map[string]string
}

func CreNull() MessageData {
	return MessageData{Type: MessageTypeNull, Data: nil}
}
func CreText(text string) MessageData {
	return MessageData{Type: MessageTypeText, Data: text}
}
func CreBool(text bool) MessageData {
	var val = "0"
	if text {
		val = "1"
	}
	return MessageData{Type: MessageTypeBool, Data: val}
}
func CreMap(mapData map[string]string) MessageData {
	return MessageData{Type: MessageTypeMap, Data: mapData}
}
func CreArrMap(mapData []map[string]string) MessageData {
	return MessageData{Type: MessageTypeArrMap, Data: mapData}
}

func CreArrString(data []string) MessageData {
	return MessageData{Type: MessageTypeArrString, Data: data}
}

func CreHttp(args []string, data map[string]string) MessageData {
	return MessageData{Type: MessageTypeHttp, Data: MessageDataHttp{args,data} }
}

func CreJson(data interface{}) MessageData {
	jsonData, _ := json.Marshal(data)
	return CreText(string(jsonData))
}

func GetText(data MessageData) string {
	if data.Type != MessageTypeText {
		fmt.Println("Core getText Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return ""
	}
	return data.Data.(string)
}
func GetBool(data MessageData) bool {
	if data.Type != MessageTypeBool {
		fmt.Println("Core GetBool Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return false
	}
	return data.Data.(string) == "1"
}
func GetMap(data MessageData) map[string]string {
	if data.Type != MessageTypeMap {
		fmt.Println("Core getMap Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return nil
	}
	return data.Data.(map[string]string)
}
func GetArrMap(data MessageData) []map[string]string {
	if data.Type != MessageTypeArrMap {
		fmt.Println("Core GetArrMap Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return nil
	}
	return data.Data.([]map[string]string)
}
func GetArrString(data MessageData) []string {
	if data.Type != MessageTypeArrString {
		fmt.Println("Core GetArrString Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return nil
	}
	return data.Data.([]string)
}
func GetHttp(data MessageData) MessageDataHttp {
	if data.Type != MessageTypeHttp {
		fmt.Println("Core GetArrString Error: "+JsonEncode(data)+" DataType: ", data.Type)
		return MessageDataHttp{}
	}

	var http = data.Data.(MessageDataHttp)
	http.Data = EscapeStringMap(http.Data)
	return http
}

func GetJson(data MessageData, v interface{}) {
	JsonDecode(GetText(data), v)
}
