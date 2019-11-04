package jsonrpc2

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (

	// CodeParseError means an invalid JSON was received by the server.
	// An error occurred on the server while parsing the JSON text.
	CodeParseError = -32700
	//CodeInvalidRequest means the JSON sent is not a valid Request object.
	CodeInvalidRequest = -32600
	// CodeMethodNotFound is generated when the method does not exist / is not available.
	CodeMethodNotFound = -32601
	// CodeInvalidParams means invalid method parameter(s) were found.
	CodeInvalidParams = -32602
	// CodeInternalError is used to denote internal JSON-RPC error
	CodeInternalError = -32603
)

//Request is a JSON RPC 2.0 Request. If an ID is not present then it is a notification
//see  https://www.jsonrpc.org/specification for details
type Request struct {
	Version VersionTag       `json:"jsonrpc"`
	ID      *ID              `json:"id,omitempty"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params,omitempty"`
}

//Response is a JSON RPC 2.0 Response
//see  https://www.jsonrpc.org/specification for details
type Response struct {
	Version VersionTag       `json:"jsonrpc"`
	ID      *ID              `json:"id,omitempty"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *Error           `json:"error,omitempty"`
}

//Error is a JSON RPC 2.0 Error
//see  https://www.jsonrpc.org/specification for details
type Error struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    *json.RawMessage `json:"data"`
}

//ID is a Request identifier, which is either a number or string
type ID struct {
	NumberID int64
	StringID string
}

func (id *ID) String() string {
	if id == nil {
		return ""
	}
	if id.StringID != "" {
		return strconv.Quote(id.StringID)
	}
	return "#" + strconv.FormatInt(id.NumberID, 10)
}

//MarshalJSON converts ID to JSON using the String representation first if it has a non-zero value,
//otherwise uses the Number value
func (id *ID) MarshalJSON() ([]byte, error) {
	if id.StringID != "" {
		return json.Marshal(id.StringID)
	}
	return json.Marshal(id.NumberID)
}

//UnmarshalJSON decodes JSON RPC 2.0 ID
func (id *ID) UnmarshalJSON(js []byte) error {
	*id = ID{}
	if err := json.Unmarshal(js, &id.NumberID); err != nil {
		return json.Unmarshal(js, &id.StringID)
	}
	return nil
}

//VersionTag encodes the JSON RPC 2.0 protocol version
type VersionTag struct {
}

//MarshalJSON encodes the JSON RPC 2.0 protocol number
func (VersionTag) MarshalJSON() ([]byte, error) {
	return json.Marshal("2.0")
}

//UnmarshalJSON decodes version, and errors if anything other than 2.0 is found
func (VersionTag) UnmarshalJSON(js []byte) error {
	var version string
	if err := json.Unmarshal(js, &version); err != nil {
		return err
	}

	if version != "2.0" {
		return fmt.Errorf("Expected JSON RPC version 2.0, but got %s", version)
	}

	return nil
}
