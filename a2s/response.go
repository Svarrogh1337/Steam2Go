package a2s

import (
	"encoding/binary"
	"log"
)

const (
	theShip   = 2400
	Challenge = 100
	Info      = 200
	Player    = 300
	Rules     = 400
	Ping      = 500
)

type Response struct {
	raw      []byte
	position int
	decoded  interface{}
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
	start := r.position
	for {
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
	response := &EDFResponse{}
	switch EDF = r.readUint8(); EDF {
	case 0x01:
		response.GameID = r.readUint64()
	case 0x10:
		response.SteamID = r.readUint64()
	case 0x20:
		response.Keywords = r.readString()
	case 0x40:
		SourceTV := &SourceTVResponse{}
		SourceTV.Port = r.readUint16()
		SourceTV.Name = r.readString()
		response.SourceTV = SourceTV
	case 0x80:
		response.Port = r.readUint16()
	}
	return response
}
func (r *Response) Decode() {
	var data InfoResponse
	data.PacketHeader = r.readUint32()
	log.Println(data.PacketHeader)
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
	r.decoded = data
}
