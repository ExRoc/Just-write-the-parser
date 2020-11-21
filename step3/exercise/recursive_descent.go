package exercise

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
	p.peek = p.current()
	p.pos++
	return p
}

func (p *parser) current() byte {
	if p.pos >= 0 && p.pos < len(p.input) {
		return p.input[p.pos]
	}
	return 0
}

func (p *parser) next() byte {
	c := p.peek
	p.peek = p.current()
	p.pos++
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
