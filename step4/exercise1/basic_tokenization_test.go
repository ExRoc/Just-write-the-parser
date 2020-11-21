package exercise1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicTokenization(t *testing.T) {
	a := assert.New(t)

	a.Equal(2*(6+4)*5+5, ParsingByASinglePass("2 * ( 6 + 4 ) * 5 + 5 "))
	a.Equal((10+97)*(25*5), ParsingByASinglePass("(   10 +  97  ) * (  25 * 5 )"))
	a.Equal((10+97)*(25*5)+55, ParsingByASinglePass("(   10 +  97  	) * (  25 * 5 ) + 55"))
	a.Equal(92*(12)*(1+(1)), ParsingByASinglePass("92 * ( 12 ) * ( 1 + ( 1 ))"))
	a.Equal((10+97)*(25*5)+5, ParsingByASinglePass(` (   10 +  97  )
			* (  25 * 5 ) +
			5  `))
	a.Equal((36)*(74)*(34+33), ParsingByASinglePass(`// hello world
			( 36 # this is a number 36
			// and bracket
			# emmmmmm.... and some meaningless comments	// 
			) * ( 74     	)
			*  (
			# hello world // hello world again 
			// hello world again # and again
			// 5 + 
			34+
			33)
			// over`))
}
