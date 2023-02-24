package a2s

import (
	"encoding/binary"
)

type Response struct {
	raw      []byte
	position int
	t        uint8
	decoded  interface{}
	error    error
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
func (r *Response) Decode() error {
	switch r.t {
	case response:
		r.decoded = *r.decodeInfoResponse()
	case player_request:

	}

	return r.error
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

func (r *Response) decodePlayerResponse() *InfoResponse {
	var data InfoResponse
	return &data
}
