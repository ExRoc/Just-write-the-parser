# [Step 7 - Scaling Up](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

The language we’ve considered has remained simple on purpose, but the techniques we discussed scale to much more complicated scenarios.

To conclude, let’s discuss just one extension, namely a general design pattern to support many more operators with varying precedence levels and left- or right associativity (commonly attributed to Vaughan Pratt).

Recall our initial implementation of `expr`:

```javascript
function expr() {
  let sum = 0
  split('+', () => {
    let prod = 1
    split('*', () => {
      if (peek == '(') {
        next()
        prod *= expr()
        assert(peek == ')', "expected ')': "+peek)
        next()
      } else
        prod *= number()
    })
    sum += prod
  })
  return sum
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

How should we add support for `-`, `/`, and potentially many others? Clearly we can extend `split` to look for sets of operators but we still have to list all possible precedence levels explicitly and implement the right evaluation logic in one place. This could get hairy!

Let’s refactor:

```javascript
let eval = {
  '+': (x,y) => x+y,
  '-': (x,y) => x-y,
  '*': (x,y) => x*y,
  '/': (x,y) => x/y
}
function factor() {
  if (peek == '(') {
    next()
    let res = expr()
    assert(peek == ')', "expected ')': "+peek)
    next()
    return res
  } else
    return number()
}
function term() {
  let res = factor()
  while (peek == '*' || peek == '/')
    res = eval[next()](res, factor())
  return res
}
function expr() {
  let res = term()
  while (peek == '+' || peek == '-')
    res = eval[next()](res, term())
  return res
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

On the surface we haven’t gained much as we still have only two levels of precedence — but it turns out that we can actually use a single function if we extract the current level of precedence as a parameter:

```javascript
// precedence: higher binds tighter
let prec = {
  '+': 100,
  '-': 100,
  '*': 200,
  '/': 200
}
// associativity: 1 for left, 0 for right
let assoc = {
  '+': 1,
  '-': 1,
  '*': 1,
  '/': 1
}
function binop(min) {
  let res = factor()
  while (peek in prec && prec[peek] >= min) {
    let nextMin = prec[peek] + assoc[peek] // + 1 for left assoc
    res = eval[next()](res, binop(nextMin))
  }
  return res
}
function expr() {
  return binop(0)
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

Voilà, now we can support an arbitrary number of operators with varying precedence levels and associativity behavior.

```javascript
print(parse("8-12/3") + " == " + (8-12/3))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

4 == 4

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)