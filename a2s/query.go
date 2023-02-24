package a2s

import "bytes"

type InfoResponse struct {
	PacketHeader    uint32
	PayloadHeader   byte
	Protocol        byte
	Name            string
	Map             string
	Folder          string
	Game            string
	ID              uint16
	Players         byte
	MaxPlayers      byte
	Bots            byte
	ServerType      byte
	Environment     byte
	Visibility      byte
	VAC             byte
	TheShipResponse *TheShipResponse
	Version         string
	EDF             byte
	EDFResponse     *EDFResponse
}

type TheShipResponse struct {
	Mode      byte
	Witnesses byte
	Duration  byte
}
type EDFResponse struct {
	Port     uint16
	SteamID  uint64
	SourceTV *SourceTVResponse
	Keywords string
	GameID   uint64
}
type SourceTVResponse struct {
	Port uint16
	Name string
}

func RequestInfo() *Request {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x54})
	msg.WriteString("Source Engine Query")
	msg.WriteByte(0)
	return &Request{
		msg:         msg,
		response:    nil,
		requestType: 0x54,
	}
}

func RequestPlayer() *Request {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0x55, 0xFF, 0xFF, 0xFF, 0xFF})
	return &Request{
		msg:         msg,
		response:    nil,
		requestType: 0x55,
	}
}
