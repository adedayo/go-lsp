package lsp

import (
	"encoding/json"

	"github.com/adedayo/go-lsp/pkg/code"
)

//InitializeParams see https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/
type InitializeParams struct {
	ProcessID             *int64             `json:"processId,omitempty"`
	ClientInfo            *ClientInfo        `json:"clientInfo,omitempty"`
	RootPath              *string            `json:"rootPath,omitempty"`
	RootURI               *code.DocumentURI  `json:"documentUri,omitempty"`
	InitializationOptions *json.RawMessage   `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	Trace                 *string            `json:"trace,omitempty"`
	WorkspaceFolders      []WorkspaceFolder  `json:"workspaceFolders,omitempty"`
}

//ClientInfo Information about the client
type ClientInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

//ClientCapabilities The capabilities provided by the client (editor or tool)
type ClientCapabilities struct {
	WorkspaceCapabilities    *WorkspaceCapabilities          `json:"workspace,omitempty"`
	TextDocumentCapabilities *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	Experimental             *json.RawMessage                `json:"experimental,omitempty"`
}

//WorkspaceCapabilities are workspace-specific client capabilities.
type WorkspaceCapabilities struct {
	ApplyEdit              *bool                                     `json:"applyEdit,omitempty"`
	WorkspaceEdit          *WorkspaceEditClientCapabilities          `json:"workspaceEdit,omitempty"`
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`
	DidChangeWatchedFiles  *DidChangeWatchedFilesClientCapabilities  `json:"didChangeWatchedFiles,omitempty"`
	Symbol                 *WorkspaceSymbolClientCapabilities        `json:"symbol,omitempty"`
	ExecuteCommand         *ExecuteCommandClientCapabilities         `json:"executeCommand,omitempty"`
}

//TextDocumentClientCapabilities Text document specific client capabilities
type TextDocumentClientCapabilities struct {
	Synchronization    *TextDocumentSyncClientCapabilities         `json:"synchronization,omitempty"`
	Completion         *CompletionClientCapabilities               `json:"completion,omitempty"`
	Hover              *HoverClientCapabilities                    `json:"hover,omitempty"`
	SignatureHelp      *SignatureHelpClientCapabilities            `json:"signatureHelp,omitempty"`
	Declaration        *DeclarationClientCapabilities              `json:"declaration,omitempty"`
	Definition         *DefinitionClientCapabilities               `json:"definition,omitempty"`
	TypeDefinition     *TypeDefinitionClientCapabilities           `json:"typeDefinition,omitempty"`
	Implementation     *ImplementationClientCapabilities           `json:"implementation,omitempty"`
	References         *ReferenceClientCapabilities                `json:"references,omitempty"`
	DocumentHighlight  *DocumentHighlightClientCapabilities        `json:"documentHighlight,omitempty"`
	DocumentSymbol     *DocumentSymbolClientCapabilities           `json:"documentSymbol,omitempty"`
	CodeAction         *CodeActionClientCapabilities               `json:"codeAction,omitempty"`
	CodeLens           *CodeLensClientCapabilities                 `json:"codeLens,omitempty"`
	DocumentLink       *DocumentLinkClientCapabilities             `json:"documentLink,omitempty"`
	ColorProvider      *DocumentColorClientCapabilities            `json:"colorProvider,omitempty"`
	Formatting         *DocumentFormattingClientCapabilities       `json:"formatting,omitempty"`
	RangeFormatting    *DocumentRangeFormattingClientCapabilities  `json:"rangeFormatting,omitempty"`
	OnTypeFormatting   *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`
	Rename             *RenameClientCapabilities                   `json:"rename,omitempty"`
	PublishDiagnostics *PublishDiagnosticsClientCapabilities       `json:"publishDiagnostics,omitempty"`
	FoldingRange       *FoldingRangeClientCapabilities             `json:"foldingRange,omitempty"`
}

//WorkspaceFolder a workspace folder
type WorkspaceFolder struct {
	URI  code.DocumentURI `json:"uri"`
	Name code.DocumentURI `json:"name"`
}

//WorkspaceEditClientCapabilities Capabilities specific to `WorkspaceEdit`s
type WorkspaceEditClientCapabilities struct {
	DocumentChanges    *bool                   `json:"DocumentChanges,omitempty"`
	ResourceOperations []resourceOperationKind `json:"resourceOperations,omitempty"`
	FailureHandling    *failureHandlingKind    `json:"failureHandling,omitempty"`
}

//DidChangeConfigurationClientCapabilities is a notification sent from the client to the server to signal the change of configuration settings
type DidChangeConfigurationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DidChangeWatchedFilesClientCapabilities is the watched files notification sent from the client to the server when the client detects changes to files watched by the language client
type DidChangeWatchedFilesClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//ExecuteCommandClientCapabilities are capabilities specific to the `workspace/executeCommand` request
type ExecuteCommandClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//WorkspaceSymbolClientCapabilities is the workspace symbol request sent from the client to the server to list project-wide symbols matching the query string.
type WorkspaceSymbolClientCapabilities struct {
	DynamicRegistration *bool             `json:"dynamicRegistration,omitempty"`
	SymbolKind          *symbolKindValues `json:"symbolKind,omitempty"`
}

type symbolKindValues struct {
	ValueSet []symbolKind `json:"valueSet,omitempty"`
}

//TextDocumentSyncClientCapabilities client capabilities for syncing text documents ;-)
type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	WillSave            *bool `json:"willSave,omitempty"`
	WillSaveWaitUntil   *bool `json:"willSaveWaitUntil,omitempty"`
	DidSave             *bool `json:"didSave,omitempty"`
}

//CompletionClientCapabilities are capabilities specific to the `textDocument/completion`
type CompletionClientCapabilities struct {
	DynamicRegistration *bool                     `json:"dynamicRegistration,omitempty"`
	CompletionItem      *completionItem           `json:"completionItem,omitempty"`
	CompletionItemKind  *completionItemKindValues `json:"completionItemKind,omitempty"`
	ContextSupport      *bool                     `json:"contextSupport,omitempty"`
}

type completionItemKindValues struct {
	ValueSet []completionItemKind `json:"valueSet,omitempty"`
}

type completionItem struct {
	SnippetSupport          *bool              `json:"snippetSupport,omitempty"`
	CommitCharactersSupport *bool              `json:"commitCharactersSupport,omitempty"`
	DocumentationFormat     []markupKind       `json:"documentationFormat,omitempty"`
	DeprecatedSupport       *bool              `json:"deprecatedSupport,omitempty"`
	PreselectSupport        *bool              `json:"preselectSupport,omitempty"`
	TagSupport              *completionItemTag `json:"tagSupport,omitempty"`
}

type markupKind string
type codeActionKind string
type symbolKind int
type completionItemTag int
type diagnosticTag int
type completionItemKind int
type resourceOperationKind string
type failureHandlingKind string

//DiagnosticRelatedInformation Represents a related message and source code location for a diagnostic. This should be
// used to point to code locations that cause or are related to a diagnostics, e.g when duplicating
// a symbol in a scope.
type DiagnosticRelatedInformation struct {
	Location code.Location `json:"location"`
	Message  string        `json:"message"`
}

//HoverClientCapabilities describes client capabilities specific to the `textDocument/hover`
type HoverClientCapabilities struct {
	DynamicRegistration *bool        `json:"dynamicRegistration,omitempty"`
	ContentFormat       []markupKind `json:"contentFormat,omitempty"`
}

//SignatureHelpClientCapabilities describes client capabilities specific to the `textDocument/signatureHelp`
type SignatureHelpClientCapabilities struct {
	DynamicRegistration  *bool                 `json:"dynamicRegistration,omitempty"`
	SignatureInformation *signatureInformation `json:"signatureInformation,omitempty"`
}

type signatureInformation struct {
	DocumentFormat       []markupKind          `json:"documentFormat,omitempty"`
	ParameterInformation *parameterInformation `json:"parameterInformation,omitempty"`
}

type parameterInformation struct {
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
}

//DeclarationClientCapabilities describes client capabilities specific to the `textDocument/declaration`
type DeclarationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         *bool `json:"linkSupport,omitempty"`
}

//DefinitionClientCapabilities describes client capabilities specific to the `textDocument/definition`.
type DefinitionClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         *bool `json:"linkSupport,omitempty"`
}

//TypeDefinitionClientCapabilities describes client capabilities specific to the `textDocument/typeDefinition`.
type TypeDefinitionClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         *bool `json:"linkSupport,omitempty"`
}

//ImplementationClientCapabilities describes client capabilities specific to the `textDocument/implementation`.
type ImplementationClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	LinkSupport         *bool `json:"linkSupport,omitempty"`
}

//ReferenceClientCapabilities describes client capabilities specific to the `textDocument/references`.
type ReferenceClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentHighlightClientCapabilities describes client capabilities specific to the `textDocument/documentHighlight`.
type DocumentHighlightClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentSymbolClientCapabilities describes client capabilities specific to the `textDocument/documentSymbol`.
type DocumentSymbolClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//CodeActionClientCapabilities describes client capabilities specific to the `textDocument/codeAction`.
type CodeActionClientCapabilities struct {
	DynamicRegistration      *bool                     `json:"dynamicRegistration,omitempty"`
	CodeActionLiteralSupport *codeActionLiteralSupport `json:"codeActionLiteralSupport,omitempty"`
	IsPreferredSupport       *bool                     `json:"isPreferredSupport,omitempty"`
}

type codeActionLiteralSupport struct {
	CodeActionKind codeAction `json:"codeActionKind"`
}

type codeAction struct {
	ValueSet []codeActionKind `json:"valueSet"`
}

//CodeLensClientCapabilities describes client capabilities specific to the `textDocument/codeLens`.
type CodeLensClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentLinkClientCapabilities describes client capabilities specific to the `textDocument/documentLink`.
type DocumentLinkClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	TooltipSupport      *bool `json:"tooltipSupport,omitempty"`
}

//DocumentColorClientCapabilities describes client capabilities specific to the `textDocument/documentColor` and the `textDocument/colorPresentation` request.
type DocumentColorClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentFormattingClientCapabilities describes client capabilities specific to the `textDocument/formatting`.
type DocumentFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentRangeFormattingClientCapabilities describes client capabilities specific to the `textDocument/rangeformatting`.
type DocumentRangeFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//DocumentOnTypeFormattingClientCapabilities describes client capabilities specific to the `textDocument/onTypeformatting`.
type DocumentOnTypeFormattingClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

//RenameClientCapabilities describes client capabilities specific to the `textDocument/rename`.
type RenameClientCapabilities struct {
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
	PrepareSupport      *bool `json:"prepareSupport,omitempty"`
}

//PublishDiagnosticsClientCapabilities describes client capabilities specific to the `textDocument/publishDiagnostics`.
type PublishDiagnosticsClientCapabilities struct {
	RelatedInformation *bool                 `json:"relatedInformation,omitempty"`
	TagSupport         *diagnosticTagSupport `json:"tagSupport,omitempty"`
}

type diagnosticTagSupport struct {
	ValueSet []diagnosticTag `json:"valueSet"`
}

//FoldingRangeClientCapabilities describes client capabilities specific to `textDocument/foldingRange requests`.
type FoldingRangeClientCapabilities struct {
	DynamicRegistration *bool  `json:"dynamicRegistration,omitempty"`
	RangeLimit          *int64 `json:"RangeLimit,omitempty"`
	LineFoldingOnly     *bool  `json:"lineFoldingOnly,omitempty"`
}
