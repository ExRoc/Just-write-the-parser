package version1

import (
	"fmt"
	"unicode"

	"github.com/Dmaxiya/Just-write-the-parser/util"
)

type parser struct {
	input string
	pos   int
	peek  byte
}

func newParser(input string) *parser {
	p := &parser{
		input: input,
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
		if p.inDelimitedCommentBeginMark() {
			p.delimitedComments()
		}
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
	p.peek = p.current()
	p.pos++
}

func (p *parser) next() byte {
	c := p.peek
	p.read()
	return c
}

func (p *parser) split(d byte, f func()) {
	for {
		f()
		if p.peek == d {
			p.next()
		} else {
			break
		}
	}
}

func (p *parser) number() int {
	util.Assert(unicode.IsDigit(rune(p.peek)), fmt.Sprintf("expected a num: %s", string(p.peek)))
	n := util.Number(p.next())
	for unicode.IsDigit(rune(p.peek)) {
		n = n*10 + util.Number(p.next())
	}
	return n
}

func (p *parser) expr() int {
	sum := 0
	p.split('+', func() {
		prod := 1
		p.split('*', func() {
			if p.peek == '(' {
				p.next()
				prod *= p.expr()
				util.Assert(p.peek == ')', fmt.Sprintf("expected ')': %s", string(p.peek)))
				p.next()
			} else {
				prod *= p.number()
			}
		})
		sum += prod
	})
	return sum
}

func ParsingByASinglePass(input string) int {
	p := newParser(input)
	res := p.expr()
	util.Assert(p.peek == 0, fmt.Sprintf("unexpected input: %s", string(p.peek)))
	return res
}
