package code

//Position in a text document is a zero-based line and zero-based character offset. A position is between two characters like an ‘insert’ cursor in a editor.
type Position struct {

	//Line position in a document (zero-based).
	Line int

	//Character offset on a line in a document (zero-based). Assuming that the line is
	//represented as a string, the `character` value represents the gap between the
	//`character` and `character + 1`.
	//If the character value is greater than the line length it defaults back to the
	//line length.
	Character int
}

//Range in a text document expressed as (zero-based) start and end positions. A range is comparable to a selection in an editor. Therefore the end position is exclusive. If you want to specify a range that contains a line including the line ending character(s) then use an end position denoting the start of the next line.
type Range struct {
	Start Position
	End   Position
}

//Diagnostic represents some diagnostic information such as compiler error or warning
type Diagnostic struct {
	Range Range
	//TBC
}
