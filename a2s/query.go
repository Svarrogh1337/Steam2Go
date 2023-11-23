package a2s

import "bytes"

func RequestInfo() *Request {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, infoRequest})
	msg.WriteString("Source Engine Query")
	msg.WriteByte(0)
	return &Request{
		msg:         msg,
		requestType: infoRequest,
	}
}

func RequestPlayer() *Request {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, playerRequest, 0xFF, 0xFF, 0xFF, 0xFF})
	return &Request{
		msg:         msg,
		requestType: playerRequest,
	}
}

func RequestRules() *Request {
	var msg bytes.Buffer
	msg.Write([]byte{0xFF, 0xFF, 0xFF, 0xFF, rulesRequest, 0xFF, 0xFF, 0xFF, 0xFF})
	return &Request{
		msg:         msg,
		requestType: rulesRequest,
	}
}
