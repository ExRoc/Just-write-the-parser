package exercise

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/Dmaxiya/Just-write-the-parser/step7/exercise/element"
	"github.com/Dmaxiya/Just-write-the-parser/util"
)

type parser struct {
	input string
	pos   int
	peek  *element.Element
}

func newParser(input string) *parser {
	p := &parser{
		input: input,
		peek:  &element.Element{},
	}
	p.read()
	return p
}

func (p *parser) indexOf(pos int) byte {
	if pos >= 0 && pos < len(p.input) {
		return p.input[pos]
	}
	return 0
}

func (p *parser) current() byte {
	return p.indexOf(p.pos)
}

func (p *parser) nextIs(str string) bool {
	for index := range str {
		if p.indexOf(p.pos+index) != str[index] {
			return false
		}
	}
	return true
}

func (p *parser) inLineCommentMark() bool {
	return p.nextIs("#") || p.nextIs("//")
}

func (p *parser) lineComments() {
	if !p.inLineCommentMark() {
		return
	}
	for p.pos < len(p.input) && p.current() != '\n' {
		p.pos++
	}
}

func (p *parser) inDelimitedCommentBeginMark() bool {
	return p.nextIs("/*")
}

func (p *parser) inDelimitedCommentEndMark() bool {
	return p.nextIs("*/")
}

func (p *parser) delimitedComments() {
	if !p.inDelimitedCommentBeginMark() {
		return
	}
	p.pos += 2
	for p.pos < len(p.input) && !p.inDelimitedCommentEndMark() {
		p.pos++
	}
	p.pos++
}

func (p *parser) comments() {
	p.lineComments()
	p.delimitedComments()
}

func (p *parser) whitespace() {
	for unicode.IsSpace(rune(p.current())) || p.inLineCommentMark() || p.inDelimitedCommentBeginMark() {
		p.comments()
		p.pos++
	}
}

func (p *parser) read() {
	p.whitespace()
	p.peek.Reset()

	if unicode.IsDigit(rune(p.current())) {
		p.peek.SetType(element.Number)
		for unicode.IsDigit(rune(p.current())) {
			p.peek.WriteByte(p.current())
			p.pos++
		}
		return
	}

	p.peek.SetType(element.ByteToElementTypeMap[p.current()])
	p.peek.WriteByte(p.current())
	p.pos++
}

func (p *parser) next() *element.Element {
	c := p.peek
	p.read()
	return c
}

func (p *parser) split(d *element.Element, f func()) {
	for {
		f()
		if p.peek.Equal(d) {
			p.next()
		} else {
			break
		}
	}
}

func (p *parser) number() int {
	util.Assert(p.peek.IsNumber(), fmt.Sprintf("expected a num: %s", p.peek.String()))
	n, _ := strconv.ParseInt(p.peek.String(), 10, 64)
	p.read()
	return int(n)
}

func (p *parser) factor() int {
	if p.peek.IsLeftBracket() {
		p.next()
		res := p.expr()
		util.Assert(p.peek.IsRightBracket(), fmt.Sprintf("expected ')': %s", p.peek.String()))
		p.next()
		return res
	}
	return p.number()
}

func (p *parser) binop(min int) int {
	res := p.factor()
	for element.Prec[p.peek.String()] >= min {
		nextMin := element.Prec[p.peek.String()] + element.Assoc[p.peek.String()]
		op := p.peek.String()
		p.next()
		res = element.Eval[op](res, p.binop(nextMin))
	}
	return res
}

func (p *parser) expr() int {
	return p.binop(1)
}

func ParsingByASinglePass(input string) int {
	p := newParser(input)
	res := p.expr()
	util.Assert(p.peek.IsNone(), fmt.Sprintf("unexpected input: %s", p.peek.String()))
	return res
}
