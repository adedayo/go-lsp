package jsonrpc2

import (
	"encoding/json"
)

//Transport is a protocol-agnostic transport interface for JSON RPC 2.0 messages
type Transport interface {
	//SendResponse constructs a JSON RPC 2.0 Response using the `id` and `data` (converted to JSON raw message)
	//over a Stream returning an error as may be necessary
	SendResponse(id *ID, data interface{}) error
	//SendNotification constructs a JSON RPC 2.0 Notification using the `data` (converted to JSON raw message)
	//over a Stream returning an error as may be necessary
	SendNotification(data interface{}) error
	//SendErrorResponse sends an error `err` over some Stream in response to an error associated with the response identified by `id`
	SendErrorResponse(id *ID, err *Error) error
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

//SendResponse constructs a JSON RPC 2.0 Response using the `id` and `data` (converted to JSON raw message)
//over a Stream returning an error as may be necessary
func (dt *DefaultTransport) SendResponse(id *ID, data interface{}) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rawMessage := json.RawMessage(raw)
	response := Response{
		Version: VersionTag{},
		ID:      id,
		Result:  &rawMessage,
	}
	outBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	dt.io.Write(outBytes)
	return nil

}

//SendNotification constructs a JSON RPC 2.0 Notification using the `data` (converted to JSON raw message)
//over a Stream returning an error as may be necessary
func (dt *DefaultTransport) SendNotification(method string, data interface{}) error {
	raw, err := json.Marshal(data)
	if err != nil {
		return err
	}
	rawMessage := json.RawMessage(raw)
	response := Request{
		Version: VersionTag{},
		Method:  method,
		Params:  &rawMessage,
	}
	outBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	dt.io.Write(outBytes)
	return nil
}

//SendErrorResponse sends an error `errX` over some Stream in response to an error associated with the response identified by `id`
func (dt *DefaultTransport) SendErrorResponse(id *ID, errX *Error) error {
	response := Response{
		Version: VersionTag{},
		ID:      id,
		Error:   errX,
	}

	outBytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	dt.io.Write(outBytes)
	return nil
}
