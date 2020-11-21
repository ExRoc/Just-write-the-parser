package exercise

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
	a.Equal((3 * ((1 + 4*(2)) + 2)), ParsingByASinglePass(`3*(
	1+4*
		2	
+2)`))
	a.Equal((10+97)*((25*5)+(5)), ParsingByASinglePass(` (   10 +  97  )*
	(  25 * 5 ) +
		5  `))
	a.Equal((36)*((74)*(34+33)), ParsingByASinglePass(`( 36 ) *
	( 74       ) *
			34+ # hello world // hello world again
					/* hello world again and again
		hahaha */
				33`))
	a.Equal(8, ParsingByASinglePass(`2 /* first comment /* second comment */ * 2 * 2`))
	a.Equal(64*(10+(72+44*86)), ParsingByASinglePass(`64 * // see you
			/*hello /*hello /*hello // world3 world2 world1*/
			# /*    //
			/* /* */ 10+ // hello world */
			(72+44* /* there is a delimited comment between '*' and 86 // */ 86)`))
}
