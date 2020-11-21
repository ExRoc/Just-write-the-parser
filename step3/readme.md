# [Step 3 - Parentheses and Recursive Grouping](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside5)

Let’s add parentheses! These can be arbitrarily nested, so our parser becomes recursive. Hence, the approach is called “recursive descent”:

```javascript
function expr() {
  let sum = 0
  split('+', () => {
    let prod = 1
    split('*', () => {
      if (peek == '(') {
        next()           // consume opening paren
        prod *= expr()   // parse nested expression
        assert(peek == ')', "expected ')': "+peek)
        next()           // consume closing paren
      } else
        prod *= number()
    })
    sum += prod
  })
  return sum
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside5)

Done!

```
print(parse("2*(6+4)*5") + " == " + (2*(6+4)*5))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside5)

100 == 100

Looking at the formal language side of things, recursion without parameters means that the grammar is *context-free*: an expression is always parsed the same way, independent of the context it appears in.

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)