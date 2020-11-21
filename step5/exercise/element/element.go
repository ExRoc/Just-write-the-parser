package element

import "bytes"

type elementType int

const (
	None elementType = iota
	AddOperation
	MultOperation
	LeftBracket
	RightBracket
	Indent
	Wrap
	Number
)

var (
	ByteToElementTypeMap = map[byte]elementType{
		0:    None,
		'+':  AddOperation,
		'*':  MultOperation,
		'(':  LeftBracket,
		')':  RightBracket,
		'\t': Indent,
		'\n': Wrap,
	}

	AddElement          *Element
	MultElement         *Element
	LeftBracketElement  *Element
	RightBracketElement *Element
	IndentElement       *Element
	WrapElement         *Element
)

func init() {
	AddElement = &Element{eType: AddOperation}
	AddElement.value.WriteByte('+')

	MultElement = &Element{eType: MultOperation}
	MultElement.value.WriteByte('*')

	LeftBracketElement = &Element{eType: LeftBracket}
	LeftBracketElement.value.WriteByte('(')

	RightBracketElement = &Element{eType: RightBracket}
	RightBracketElement.value.WriteByte(')')

	IndentElement = &Element{eType: Indent}
	IndentElement.value.WriteByte('\t')

	WrapElement = &Element{eType: Wrap}
	WrapElement.value.WriteByte('\n')
}

type Element struct {
	value bytes.Buffer
	eType elementType
}

func (e *Element) Reset() {
	e.value.Reset()
	e.eType = None
}

func (e *Element) SetType(t elementType) {
	e.eType = t
}

func (e *Element) WriteByte(b byte) error {
	return e.value.WriteByte(b)
}

func (e *Element) String() string {
	return e.value.String()
}

func (e *Element) Equal(e2 *Element) bool {
	if e == e2 {
		return true
	}
	if e.eType != e2.eType {
		return false
	}
	if e.eType != Number {
		return true
	}
	return e.value.String() == e2.value.String()
}

func (e *Element) IsNone() bool {
	return e.eType == None
}

func (e *Element) IsAddOperation() bool {
	return e.eType == AddOperation
}

func (e *Element) IsMultOperation() bool {
	return e.eType == MultOperation
}

func (e *Element) IsLeftBracket() bool {
	return e.eType == LeftBracket
}

func (e *Element) IsRightBracket() bool {
	return e.eType == RightBracket
}

func (e *Element) IsIndent() bool {
	return e.eType == Indent
}

func (e *Element) IsWrap() bool {
	return e.eType == Wrap
}

func (e *Element) IsNumber() bool {
	return e.eType == Number
}
