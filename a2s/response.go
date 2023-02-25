package a2s

import (
	"encoding/binary"
	"log"
)

type Response struct {
	raw      []byte
	position int
	t        uint8
	decoded  interface{}
}

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
type PlayerResponse struct {
	PacketHeader   uint32
	PayloadHeader  byte
	PlayerCount    byte
	Players        []Player
	TheShipPlayers []TheShipPlayer
}
type Player struct {
	Index    byte
	Name     string
	Score    uint32
	Duration uint32
}

type TheShipPlayer struct {
	Index    byte
	Name     string
	Score    uint32
	Duration uint32
	Deaths   uint32
	Money    uint32
}

type RulesResponse struct {
	PacketHeader  uint32
	PayloadHeader byte
	RulesCount    byte
	Rules         []Rule
}

type Rule struct {
	Name  string
	Value string
}

func (r *Response) readRawUint32() []byte {
	v := r.raw[r.position : r.position+4]
	r.position += 4
	return v
}

func (r *Response) readUint64() uint64 {
	v := binary.LittleEndian.Uint64(r.raw[r.position:])
	r.position += 8
	return v
}

func (r *Response) readUint32() uint32 {
	v := binary.LittleEndian.Uint32(r.raw[r.position:])
	r.position += 4
	return v
}

func (r *Response) readUint16() uint16 {
	v := binary.LittleEndian.Uint16(r.raw[r.position:])
	r.position += 2
	return v
}

func (r *Response) readUint8() uint8 {
	v := r.raw[r.position]
	r.position += 1
	return v
}

func (r *Response) readString() string {
	log.Println(r.position)
	start := r.position
	for {
		if r.position == 1400 {
			return "Truncated response"
		}
		if r.raw[r.position] == 0 {
			r.position++
			break
		}
		r.position++
	}

	return string(r.raw[start : r.position-1])
}

func (r *Response) TheShipResponse() *TheShipResponse {
	return &TheShipResponse{
		Mode:      r.readUint8(),
		Witnesses: r.readUint8(),
		Duration:  r.readUint8(),
	}
}

func (r *Response) EDFResponse(EDF byte) *EDFResponse {
	res := &EDFResponse{}
	switch EDF = r.readUint8(); EDF {
	case gameID:
		res.GameID = r.readUint64()
	case steamID:
		res.SteamID = r.readUint64()
	case keywords:
		res.Keywords = r.readString()
	case sourceTV:
		SourceTV := &SourceTVResponse{}
		SourceTV.Port = r.readUint16()
		SourceTV.Name = r.readString()
		res.SourceTV = SourceTV
	case port:
		res.Port = r.readUint16()
	}
	return res
}

func (r *Response) Decode() {
	switch r.t {
	case info:
		r.decoded = *r.decodeInfoResponse()
	case player:
		r.decoded = *r.decodePlayerResponse()
	case rules:
		r.decoded = *r.decodeRulesResponse()
	}
}

func (r *Response) decodeInfoResponse() *InfoResponse {
	var data InfoResponse
	data.PacketHeader = r.readUint32()
	data.PayloadHeader = r.readUint8()
	data.Protocol = r.readUint8()
	data.Name = r.readString()
	data.Map = r.readString()
	data.Folder = r.readString()
	data.Game = r.readString()
	data.ID = r.readUint16()
	data.Players = r.readUint8()
	data.MaxPlayers = r.readUint8()
	data.ServerType = r.readUint8()
	data.Environment = r.readUint8()
	data.Visibility = r.readUint8()
	data.VAC = r.readUint8()
	if data.ID == theShip {
		data.TheShipResponse = r.TheShipResponse()
	}
	data.Version = r.readString()
	return &data
}

func (r *Response) decodePlayerResponse() *PlayerResponse {
	var data PlayerResponse
	data.PacketHeader = r.readUint32()
	data.PayloadHeader = r.readUint8()
	data.PlayerCount = r.readUint8()
	for i := 0; i <= int(data.PlayerCount); i++ {
		if r.position == len(r.raw) {
			log.Printf("Steam2Go A2S: Truncated response. To recieve the full response increase the Client size.")
			break
		}
		p := &Player{
			Index:    r.readUint8(),
			Name:     r.readString(),
			Score:    r.readUint32(),
			Duration: r.readUint32(),
		}
		data.Players = append(data.Players, *p)
	}
	return &data
}

func (r *Response) decodeRulesResponse() *RulesResponse {
	var data RulesResponse
	data.PacketHeader = r.readUint32()
	data.PayloadHeader = r.readUint8()
	data.RulesCount = r.readUint8()
	log.Println(data.RulesCount)
	for i := 0; i <= int(data.RulesCount); i++ {
		if r.position == len(r.raw) {
			log.Printf("Steam2Go A2S: Truncated response. To recieve the full response increase the Client size.")
			break
		}
		rule := &Rule{
			Name:  r.readString(),
			Value: r.readString(),
		}
		data.Rules = append(data.Rules, *rule)
	}
	return &data
}
