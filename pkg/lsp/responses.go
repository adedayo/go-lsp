package lsp

import "encoding/json"

//InitializeResult see: https://microsoft.github.io/language-server-protocol/specifications/specification-3-15/#initialize
type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *serverInfo        `json:"serverInfo,omitempty"`
}

type serverInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}

//ServerCapabilities represent the capabilities the language server provides.
type ServerCapabilities struct {
	TextDocumentSync          *textDocSyncOptionsOrLegacy  `json:"textDocumentSync,omitempty"`
	CompletionProvider        *CompletionOptions           `json:"completionProvider,omitempty"`
	HoverProvider             *hoverUnion                  `json:"hoverProvider,omitempty"`
	SignatureHelpProvider     *SignatureHelpOptions        `json:"signatureHelpProvider,omitempty"`
	DeclarationProvider       *declarationUnion            `json:"declarationProvider,omitempty"`
	DefinitionProvider        *definitionUnion             `json:"definitionProvider,omitempty"`
	TypeDefinitionProvider    *typeDefinitionUnion         `json:"typeDefinitionProvider,omitempty"`
	ImplementationProvider    *implementationProviderUnion `json:"implementationProvider,omitempty"`
	ReferencesProvider        *referencesUnion             `json:"referencesProvider,omitempty"`
	DocumentHighlightProvider *documentHighlightUnion      `json:"documentHighlightProvider,omitempty"`
	//TODO: Complete the rest
	// DocumentSymbolProvider           *documentSymbolUnion           `json:"documentSymbolProvider,omitempty"`
	// CodeActionProvider               *codeActionUnion               `json:"codeActionProvider,omitempty"`
	// CodeLensProvider                 *codeLensUnion                 `json:"codeLensProvider,omitempty"`
	// DocumentLinkProvider             *DocumentLinkOptions           `json:"documentLinkProvider,omitempty"`
	// ColorProvider                    *colorUnion                    `json:"colorProvider,omitempty"`
	// DocumentFormattingProvider       *documentFormattingUnion       `json:"documentFormattingProvider,omitempty"`
	// DocumentRangeFormattingProvider  *documentRangeFormattingUnion  `json:"documentRangeFormattingProvider,omitempty"`
	// DocumentOnTypeFormattingProvider *documentOnTypeFormattingUnion `json:"documentOnTypeFormattingProvider,omitempty"`
	// RenameProvider                   *renameUnion                   `json:"renameProvider,omitempty"`
	// FoldingRangeProvider             *foldingRangeUnion             `json:"foldingRangeProvider,omitempty"`
	// ExecuteCommandProvider           *ExecuteCommandOptions         `json:"executeCommandProvider,omitempty"`
	WorkspaceSymbolProvider *bool `json:"workspaceSymbolProvider,omitempty"`
	// Workspace                        *workspaceServerCapabilities   `json:"workspace,omitempty"`
	Experimental *json.RawMessage `json:"experimental,omitempty"`
}

type textDocumentSyncKind int
type textDocSyncOptionsOrLegacy struct {
	Kind    *textDocumentSyncKind
	Options TextDocumentSyncOptions
}

func (textSync *textDocSyncOptionsOrLegacy) MarshalJSON() ([]byte, error) {
	if textSync.Kind != nil {
		return json.Marshal(*textSync.Kind)
	}
	return json.Marshal(textSync.Options)
}

func (textSync *textDocSyncOptionsOrLegacy) UnmarshalJSON(js []byte) error {
	*textSync = textDocSyncOptionsOrLegacy{}
	if err := json.Unmarshal(js, textSync.Kind); err != nil {
		return json.Unmarshal(js, &textSync.Options)
	}
	return nil
}

//TextDocumentSyncOptions options
type TextDocumentSyncOptions struct {
	OpenClose *bool                 `json:"openClose,omitempty"`
	Change    *textDocumentSyncKind `json:"change,omitempty"`
}

//WorkDoneProgressOptions Options to signal work done progress support in server capabilities.
type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`
}

//CompletionOptions indicates whether the server provides completion support
type CompletionOptions struct {
	*WorkDoneProgressOptions
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`
	ResolveProvider     *bool    `json:"resolveProvider,omitempty"`
}

type hoverUnion struct {
	Boolean *bool
	Options HoverOptions
}

func (hu *hoverUnion) MarshalJSON() ([]byte, error) {
	if hu.Boolean != nil {
		return json.Marshal(*hu.Boolean)
	}
	return json.Marshal(hu.Options)
}

func (hu *hoverUnion) UnmarshalJSON(js []byte) error {
	*hu = hoverUnion{}
	if err := json.Unmarshal(js, hu.Boolean); err != nil {
		return json.Unmarshal(js, &hu.Options)
	}
	return nil
}

//HoverOptions indicates whether the server provides hover support
type HoverOptions struct {
	*WorkDoneProgressOptions
}

//SignatureHelpOptions indicates whether the server provides signature help support
type SignatureHelpOptions struct {
	*WorkDoneProgressOptions
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
	ResolveProvider     *bool    `json:"resolveProvider,omitempty"`
}

//DeclarationOptions indicates whether the server provides declaration support
type DeclarationOptions struct {
	*WorkDoneProgressOptions
}

//DeclarationRegistrationOptions indicates whether the server provides declaration registration support
type DeclarationRegistrationOptions struct {
	*DeclarationOptions
	*TextDocumentRegistrationOptions
	*StaticRegistrationOptions
}

//TextDocumentRegistrationOptions used to dynamically register for requests for a set of text documents.
type TextDocumentRegistrationOptions struct {
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}

//DocumentSelector is a combination of one or more document filters
type DocumentSelector []DocumentFilter

//DocumentFilter denotes a document through properties like language, scheme or pattern
type DocumentFilter struct {
	Language *string `json:"language,omitempty"`
	Scheme   *string `json:"scheme,omitempty"`
	Pattern  *string `json:"pattern,omitempty"`
}

//StaticRegistrationOptions used to register a feature in the initialize result with a given server control ID to be able to un-register the feature later on.
type StaticRegistrationOptions struct {
	ID *string `json:"id,omitempty"`
}

type declarationUnion struct {
	Boolean             *bool
	Options             *DeclarationOptions
	RegistrationOptions *DeclarationRegistrationOptions
}

func (du *declarationUnion) MarshalJSON() ([]byte, error) {
	if du.Boolean != nil {
		return json.Marshal(*du.Boolean)
	}
	if du.Options != nil {
		return json.Marshal(*du.Options)
	}
	return json.Marshal(*du.RegistrationOptions)
}

func (du *declarationUnion) UnmarshalJSON(js []byte) error {
	*du = declarationUnion{}
	if err := json.Unmarshal(js, *du.Boolean); err != nil {
		if err = json.Unmarshal(js, *du.Options); err != nil {
			return json.Unmarshal(js, *du.RegistrationOptions)
		}
	}
	return nil
}

type definitionUnion struct {
	Boolean *bool
	Options DefinitionOptions
}

//DefinitionOptions indicates whether server provides goto definition support
type DefinitionOptions struct {
	*WorkDoneProgressOptions
}

func (du *definitionUnion) MarshalJSON() ([]byte, error) {
	if du.Boolean != nil {
		return json.Marshal(*du.Boolean)
	}
	return json.Marshal(du.Options)
}

func (du *definitionUnion) UnmarshalJSON(js []byte) error {
	*du = definitionUnion{}
	if err := json.Unmarshal(js, du.Boolean); err != nil {
		return json.Unmarshal(js, &du.Options)
	}
	return nil
}

type typeDefinitionUnion struct {
	Boolean             *bool
	Options             *TypeDefinitionOptions
	RegistrationOptions *TypeDefinitionRegistrationOptions
}

//TypeDefinitionOptions indicates whether server provides goto type definition support
type TypeDefinitionOptions struct {
	*WorkDoneProgressOptions
}

//TypeDefinitionRegistrationOptions indicates whether server provides goto type definition support
type TypeDefinitionRegistrationOptions struct {
	*TextDocumentRegistrationOptions
	*TypeDefinitionOptions
	*StaticRegistrationOptions
}

func (tdu *typeDefinitionUnion) MarshalJSON() ([]byte, error) {
	if tdu.Boolean != nil {
		return json.Marshal(*tdu.Boolean)
	}
	if tdu.Options != nil {
		return json.Marshal(*tdu.Options)
	}
	return json.Marshal(*tdu.RegistrationOptions)
}

func (tdu *typeDefinitionUnion) UnmarshalJSON(js []byte) error {
	*tdu = typeDefinitionUnion{}
	if err := json.Unmarshal(js, *tdu.Boolean); err != nil {
		if err = json.Unmarshal(js, *tdu.Options); err != nil {
			return json.Unmarshal(js, *tdu.RegistrationOptions)
		}
	}
	return nil
}

type implementationProviderUnion struct {
	Boolean             *bool
	Options             *ImplementationOptions
	RegistrationOptions *ImplementationRegistrationOptions
}

//ImplementationOptions indicates whether server provides goto implementation support
type ImplementationOptions struct {
	*WorkDoneProgressOptions
}

//ImplementationRegistrationOptions indicates whether server provides goto implementation support
type ImplementationRegistrationOptions struct {
	*TextDocumentRegistrationOptions
	*ImplementationOptions
	*StaticRegistrationOptions
}

func (ipu *implementationProviderUnion) MarshalJSON() ([]byte, error) {
	if ipu.Boolean != nil {
		return json.Marshal(*ipu.Boolean)
	}
	if ipu.Options != nil {
		return json.Marshal(*ipu.Options)
	}
	return json.Marshal(*ipu.RegistrationOptions)
}

func (ipu *implementationProviderUnion) UnmarshalJSON(js []byte) error {
	*ipu = implementationProviderUnion{}
	if err := json.Unmarshal(js, *ipu.Boolean); err != nil {
		if err = json.Unmarshal(js, *ipu.Options); err != nil {
			return json.Unmarshal(js, *ipu.RegistrationOptions)
		}
	}
	return nil
}

type referencesUnion struct {
	Boolean *bool
	Options ReferenceOptions
}

//ReferenceOptions indicates whether server provides find references support
type ReferenceOptions struct {
	*WorkDoneProgressOptions
}

func (ru *referencesUnion) MarshalJSON() ([]byte, error) {
	if ru.Boolean != nil {
		return json.Marshal(*ru.Boolean)
	}
	return json.Marshal(ru.Options)
}

func (ru *referencesUnion) UnmarshalJSON(js []byte) error {
	*ru = referencesUnion{}
	if err := json.Unmarshal(js, ru.Boolean); err != nil {
		return json.Unmarshal(js, &ru.Options)
	}
	return nil
}

type documentHighlightUnion struct {
	Boolean *bool
	Options DocumentHighlightOptions
}

//DocumentHighlightOptions indicates whether server provides document highlight support
type DocumentHighlightOptions struct {
	*WorkDoneProgressOptions
}

func (ru *documentHighlightUnion) MarshalJSON() ([]byte, error) {
	if ru.Boolean != nil {
		return json.Marshal(*ru.Boolean)
	}
	return json.Marshal(ru.Options)
}

func (ru *documentHighlightUnion) UnmarshalJSON(js []byte) error {
	*ru = documentHighlightUnion{}
	if err := json.Unmarshal(js, ru.Boolean); err != nil {
		return json.Unmarshal(js, &ru.Options)
	}
	return nil
}
