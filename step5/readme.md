# [Step 5 - Significant Whitespace](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside8)

Some languages like Python attach meaning to whitespace and use indentation to indicate grouping, much like braces or parantheses in other languages.

How can we achieve this in our parser? We need two small changes. First, we no longer skip newlines (only spaces) and we keep track of the number of spaces at the beginning of a line:

```javascript
function whitespace() {
  let start = pos
  while (input[pos] == ' ') pos++
  return pos - start
}
let indent
function init(s) { input = s; pos = 0; indent = whitespace(); read() }
function next() { let c = peek; read(); return c }
function read() {
  whitespace(); peek = input[pos++]
  if (peek == '\n') indent = whitespace()
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside8)

Second, we keep track of the current indentation of each expression and proceed much like we do for parentheses when we detect a newline with increased indentation:

```javascript
function expr(ind) {
  let sum = 0
  split('+', () => {
    let prod = 1
    split('*', () => {
      if (peek == '\n' && indent > ind) {
        next()
        prod *= expr(indent)
        if (peek == '\n' && indent == ind) next()
      } else if (peek == '(') {
        next()
        prod *= expr(ind)
        assert(peek == ')', "expected ')': "+peek); next()
      } else
        prod *= number()
    })
    sum += prod
  })
  return sum
}
function parse(s) {
  init(s)
  let res = expr(indent)
  assert(!peek, "unexpected input: "+peek)
  return res
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside8)

Now we can parse indented blocks as an alternative grouping mechanism to parentheses, or both in combination. An example:

```javascript
let inp = `3*(
 1+4*
  2
+2)`
print(parse(inp) + " == " + (3*((1+4*(2))+2)))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside8)

33 == 33

From a formal language perspective, our input language is now *context-sensitive*: how exactly an expression is parsed depends on the context it occurs in. Note how this is easily and efficiently implemented by parameterizing the recursive function `expr` with a context abstraction `ind`.

**Exercise**: observe that our parser always requires a closing paranthesis but leaves “exdentation” optional. Why is this a reasonable choice, and what would change if exdentation were required? Try it out!

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)