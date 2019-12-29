package jsonrpc2

import (
	"encoding/json"
)

//Transport is a protocol-agnostic transport interface for JSON RPC 2.0 messages
type Transport interface {
	SendResponse(id *ID, data []byte)
	SendNotification(data []byte)
	SendErrorResponse(id *ID, err *Error)
}

//DefaultTransport is a default implementation of the `Transport` interface
type DefaultTransport struct {
	io Stream
}

func (dt *DefaultTransport) Read() ([]byte, int64, error) {
	return dt.io.Read()
}

func (dt *DefaultTransport) Write(data []byte) (int64, error) {
	return dt.io.Write(data)
}

//MakeTransport creates a default transport mechanism over a `Stream`
func MakeTransport(io Stream) *DefaultTransport {
	dt := DefaultTransport{io: io}
	return &dt
}

//SendResponse sends
func (dt *DefaultTransport) SendResponse(id *ID, data []byte) {
	rawMessage := json.RawMessage(data)
	response := Response{
		Version: VersionTag{},
		ID:      id,
		Result:  &rawMessage,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		dt.io.Write(outBytes)
	}
}

func (dt *DefaultTransport) SendNotification(method string, data []byte) {
	rawMessage := json.RawMessage(data)
	response := Request{
		Version: VersionTag{},
		Method:  method,
		Params:  &rawMessage,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		dt.io.Write(outBytes)
	}
}

func (dt *DefaultTransport) SendErrorResponse(id *ID, err *Error) {
	response := Response{
		Version: VersionTag{},
		ID:      id,
		Error:   err,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		dt.io.Write(outBytes)
	}
}
