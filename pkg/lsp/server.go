package lsp

import (
	"encoding/json"
	"io"
	"os"

	"github.com/adedayo/go-lsp/pkg/jsonrpc2"
)

// var (
// f, _ = os.Create("debug_from_server.txt")
// )

//Server defines the contracts
type Server interface {
	Initialize(req *jsonrpc2.Request)
	Initialized(req *jsonrpc2.Request)
	Start(inputStream io.Reader, outputStream io.Writer)
	Shutdown(req *jsonrpc2.Request)
	Exit(req *jsonrpc2.Request)
}

//DefaultMethodProvider interface allows the embedding struct of `DefaultServer` to be able to provide
// custom implementations for LSP RPC calls that are not implemented by the DefaultServer
type DefaultMethodProvider interface {
	Default(req *jsonrpc2.Request)
}

//NewServer creates a new DefaultServer
func NewServer() Server {
	return &DefaultServer{}
}

//DefaultServer is a default implementation of a Language Server that implements the LSP
//Compose your LSP server with this and override with specifics of your language server
type DefaultServer struct {
	// io jsonrpc2.Stream
	*jsonrpc2.DefaultTransport
	initialized             bool
	receivedShutdownRequest bool
	embeddingServer         *DefaultMethodProvider
}

//Init passes in a reference to the embedding struct to allow calling its `Default` method
func (s *DefaultServer) Init(composedServer DefaultMethodProvider) {
	s.embeddingServer = &composedServer
}

func (s *DefaultServer) sendResponse(id *jsonrpc2.ID, data []byte) {
	rawMessage := json.RawMessage(data)
	response := jsonrpc2.Response{
		Version: jsonrpc2.VersionTag{},
		ID:      id,
		Result:  &rawMessage,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		s.Write(outBytes)
	}
}

func (s *DefaultServer) sendNotification(data []byte) {
	rawMessage := json.RawMessage(data)
	response := jsonrpc2.Response{
		Version: jsonrpc2.VersionTag{},
		Result:  &rawMessage,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		s.Write(outBytes)
	}
}

func (s *DefaultServer) sendErrorResponse(id *jsonrpc2.ID, err *jsonrpc2.Error) {
	response := jsonrpc2.Response{
		Version: jsonrpc2.VersionTag{},
		ID:      id,
		Error:   err,
	}
	if outBytes, err := json.Marshal(response); err == nil {
		s.Write(outBytes)
	}
}

//Start starts the LSP server protocol by reading the stream in a loop and dispatching the data to
//well-defined methods, unknown RPC method are dispatched to the Default handler
func (s *DefaultServer) Start(in io.Reader, out io.Writer) {
	s.DefaultTransport = jsonrpc2.MakeTransport(jsonrpc2.NewStream(in, out))
	// s.io = jsonrpc2.NewStream(in, out)
	for {
		if buf, _, err := s.Read(); err == nil {
			req := &jsonrpc2.Request{}
			if err := json.Unmarshal(buf, req); err == nil {
				switch req.Method {
				case "initialize":
					s.Initialize(req)
				case "initialized":
					s.Initialized(req)
				case "shutdown":
					go s.Shutdown(req)
				default:
					s.forward(req)
				}
			} else {
				f.WriteString("Got Error " + err.Error() + "\nMessage: " + string(buf))

				//TODO: determine what to do on receipt of a malformed request
			}
		} else {
			if err == io.EOF {
				break
			}
		}
	}
}

//Stop gives the server an opportunity to do any clean up as may be required
func (s *DefaultServer) Stop() {
}

//forward forwards the request to the embedding server's default handler
func (s *DefaultServer) forward(req *jsonrpc2.Request) {
	// f.WriteString(fmt.Sprintf("\n\n%s, %s\n\n", req.Method, req.Params))
	if s.embeddingServer != nil {
		(*s.embeddingServer).Default(req)
	}
}

//Initialize the initialize request is sent as the first request from the client to the server. If the server receives a request or notification before the initialize request it should act as follows:
// * For a request the response should be an error with code: -32002. The message can be picked by the server.
// * Notifications should be dropped, except for the exit notification. This will allow the exit of a server without an initialize request.
//see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#initialize
func (s *DefaultServer) Initialize(req *jsonrpc2.Request) {
	supported := true
	syncKind := textDocumentSyncKind(1)
	result := InitializeResult{
		Capabilities: ServerCapabilities{
			TextDocumentSync: &textDocSyncOptionsOrLegacy{
				Options: TextDocumentSyncOptions{
					OpenClose: &supported,
					Change:    &syncKind},
			},
			Workspace: &workspaceServerCapabilities{
				WorkspaceFolders: WorkspaceFolderServerCapabilities{
					Supported: &supported,
				},
			},
		},
	}

	if err := s.SendResponse(req.ID, result); err != nil {
		e := jsonrpc2.Error{
			Message: err.Error(),
			Code:    jsonrpc2.CodeInternalError,
		}
		s.SendErrorResponse(req.ID, &e)
	}

	s.forward(req)
}

//Initialized is called when the initialized notification is sent from the client to the server
//after the client received the result of the initialize request but before the client is sending
// any other request or notification to the server.
// see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#initialized
func (s *DefaultServer) Initialized(req *jsonrpc2.Request) {
	s.initialized = true
	s.forward(req)
}

//Shutdown : the shutdown request is sent from the client to the server. It asks the server to shut down, but to not exit
// see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#shutdown
func (s *DefaultServer) Shutdown(req *jsonrpc2.Request) {
	s.receivedShutdownRequest = true
	s.forward(req)
}

//Exit is a notification to ask the server to exit its process. The server should exit with success code 0 if the shutdown request has been received before; otherwise with error code 1.
// see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#shutdown
func (s *DefaultServer) Exit(req *jsonrpc2.Request) {
	s.forward(req)
	if s.receivedShutdownRequest {
		os.Exit(0)
	}
	os.Exit(1)
}
