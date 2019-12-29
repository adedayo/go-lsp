package lsp

import "github.com/adedayo/go-lsp/pkg/code"

//TextDocumentItem is an item to transfer a text document from the client to the server.
type TextDocumentItem struct {
	/**
	 * The text document's URI.
	 */
	URI code.DocumentURI `json:"uri"`

	/**
	 * The text document's language identifier.
	 */
	LanguageID string `json:"languageId"`

	/**
	 * The version number of this document (it will increase after each
	 * change, including undo/redo).
	 */
	Version int64 `json:"version"`

	/**
	 * The content of the opened text document.
	 */
	Text string `json:"text"`
}

//DidOpenTextDocumentParams is the parameter sent during document open notification is sent from the client to the server
//to signal newly opened text documents. The document’s content is now managed by the client and
//the server must not try to read the document’s content using the document’s Uri. Open in this
//sense means it is managed by the client. It doesn’t necessarily mean that its content is presented
//in an editor.
type DidOpenTextDocumentParams struct {
	//The document that was opened.
	TextDocument TextDocumentItem `json:"textDocument"`
}

//DidChangeTextDocumentParams is the parameters sent during document change notification is sent from the client to the server to signal changes to a text document.
type DidChangeTextDocumentParams struct {
	/**
	 * The document that did change. The version number points
	 * to the version after all provided content changes have
	 * been applied.
	 */
	TextDocument VersionedTextDocumentIdentifier `json:"textDocument"`

	/**
	 * The actual content changes. The content changes describe single state changes
	 * to the document. So if there are two content changes c1 and c2 for a document
	 * in state S then c1 move the document to S' and c2 to S''.
	 */
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

//TextDocumentIdentifier identifies a text document
type TextDocumentIdentifier struct {
	URI code.DocumentURI `json:"uri"`
}

//VersionedTextDocumentIdentifier is The version number of this document. If a versioned text document identifier
/** is sent from the server to the client and the file is not open in the editor
 * (the server has not received an open notification before) the server can send
 * `null` to indicate that the version is known and the content on disk is the
 * master (as speced with document content ownership).
 *
 * The version number of a document will increase after each change, including
 * undo/redo. The number doesn't need to be consecutive
 */
type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int64 `json:"version"`
}

//TextDocumentContentChangeEvent An event describing a change to a text document. If range and rangeLength are omitted
//the new text is considered to be the full content of the document.
type TextDocumentContentChangeEvent struct {
	Range code.Range `json:"range,omitempty"`
	Text  string     `json:"text"`
}
