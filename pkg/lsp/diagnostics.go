package lsp

import (
	"encoding/json"
	"strconv"

	"github.com/adedayo/go-lsp/pkg/code"
)

//DiagnosticSeverity indicates the severity of reported diagnostic
/**
//Reports an error.
export const Error: 1 = 1;

 //Reports a warning.
export const Warning: 2 = 2;
/**
//Reports an information.

export const Information: 3 = 3;
/**
//Reports a hint.

export const Hint: 4 = 4;
*/
type DiagnosticSeverity int

// DiagnosticCode is a code which might appear in the user interface relating to a diagnostic
type DiagnosticCode struct {
	NumberID int64
	StringID string
}

func (id *DiagnosticCode) String() string {
	if id == nil {
		return ""
	}
	if id.StringID != "" {
		return strconv.Quote(id.StringID)
	}
	return "#" + strconv.FormatInt(id.NumberID, 10)
}

// MarshalJSON marshals diagnostic code as expected
func (id *DiagnosticCode) MarshalJSON() ([]byte, error) {
	if id.StringID != "" {
		return json.Marshal(id.StringID)
	}
	return json.Marshal(id.NumberID)
}

// UnmarshalJSON unmarshals byte slice into DiagnosticCode
func (id *DiagnosticCode) UnmarshalJSON(js []byte) error {
	*id = DiagnosticCode{}
	if err := json.Unmarshal(js, &id.NumberID); err != nil {
		return json.Unmarshal(js, &id.StringID)
	}
	return nil
}

// Diagnostic represents some diagnostic information such as compiler error or warning
type Diagnostic struct {
	Range              code.Range                      `json:"range"`
	Severity           *DiagnosticSeverity             `json:"severity,omitempty"`
	Code               *DiagnosticCode                 `json:"code,omitempty"`
	Source             *string                         `json:"source,omitempty"`
	Message            string                          `json:"message"`
	Tags               *[]diagnosticTag                `json:"tags,omitempty"`
	RelatedInformation *[]DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
}
