package lsp

import (
	"encoding/json"

	"github.com/adedayo/go-lsp/pkg/jsonrpc2"
)

//Server defines the contracts
type Server interface {
	Initialize(req *jsonrpc2.Request)
	Initialized(req *jsonrpc2.Request)
	Start()
	Default(req *jsonrpc2.Request)
}

//NewServer creates a new DefaultServer
func NewServer(io Stream) Server {
	return &DefaultServer{
		io: io,
	}
}

//DefaultServer is a default implementation of a Language Server that implements the LSP
//Compose your LSP server with this and override with specifics of your language server
type DefaultServer struct {
	io          Stream
	initialized bool
}

//Start starts the LSP server protocol by reading the stream in a loop and dispatching the data to
//well-defined methods, unknown RPC method are dispatched to the Default handler
func (s *DefaultServer) Start() {
	for {
		if buf, _, err := s.io.Read(); err == nil {
			req := &jsonrpc2.Request{}
			if err := json.Unmarshal(buf, req); err == nil {
				switch req.Method {
				case "initialize":
					s.Initialize(req)
				case "initialized":
					s.Initialized(req)
				default:
					s.Default(req)
				}
			} else {
				//TODO: determine what to do on receipt of a malformed request
			}
		}
	}
}

//Initialize the initialize request is sent as the first request from the client to the server. If the server receives a request or notification before the initialize request it should act as follows:
// * For a request the response should be an error with code: -32002. The message can be picked by the server.
// * Notifications should be dropped, except for the exit notification. This will allow the exit of a server without an initialize request.
//see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#initialize
func (s *DefaultServer) Initialize(req *jsonrpc2.Request) {

}

func (s *DefaultServer) Initialized(req *jsonrpc2.Request) {

}

func (s *DefaultServer) Default(req *jsonrpc2.Request) {

}

type securityServer struct {
	*DefaultServer
}

func (ss *securityServer) Initialized(req *jsonrpc2.Request) {
	ss.DefaultServer.Initialized(req)
}
