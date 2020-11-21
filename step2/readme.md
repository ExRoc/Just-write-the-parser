# [Step 2 - Parsing in a Single Pass](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

The first implementation above is inefficient: it traverses and copies the input string multiple times. This may not actually matter in practice, but sometimes it does. It also makes it a bit harder to add certain other features, some of which are important.

**Exercise:** think about how you’d add support for parentheses (”`(...)`”) above. Will your idea deal with `(1+2)+7*((3+4)*7)+5` correctly?

Let’s change the code to process the input strictly left to right. What’s the best way to guarantee this? We’ll use an iterator interface that only allows to inspect the current character and advance to the next.

```javascript
let input, pos, peek
function init(s) { input = s; pos = 0; peek = input[pos++] }
function next() { let c = peek; peek = input[pos++]; return c }
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

Instead of using JS’s built-in `String.split` function, we implement our own replacement. Since we don’t want to return a string (which would need to be traversed again) we take a callback function as argument:

```javascript
function split(d,f) {
  for (;;) { f(); if (peek == d) next(); else break }
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

The way we use `split` is as follows. It’s really similar to what we had before — the main difference is that the string to be split is now implicit.

```javascript
function expr() {
  let sum = 0
  split('+', () => {
    let prod = 1
    split('*', () => {
      prod *= number()
    })
    sum += prod
  })
  return sum
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

We’re modularizing things a little bit on the way. The core expression parsing logic has moved to its own function `expr`, and we add a driver function `parse` to run it.

```javascript
function parse(s) {
  init(s)
  let res = expr()
  assert(!peek, "unexpected input: "+peek)
  return res
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

The top-level function `parse` also checks that we’ve consumed the entire input. This is important, because:

> Internal routines are set up to stop consuming input when they can’t make sense of the next character.

The last missing piece is our own function to parse numbers:

```javascript
function number() {
  let isdigit = () => '0' <= peek && peek <= '9'
  assert(isdigit(), "expected a number: "+peek)
  let n = Number(next())
  while (isdigit())
    n = n * 10 + Number(next())
  return n
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

The `assert` serves to assure that there is at least one digit.

Now we’re ready to run examples:

```javascript
print(parse("2*6+4*5") + " == " + (2*6+4*5))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

32 == 32

**Exercise:** how much faster is this version? (a) asymptotically, in big-O terms (b) experimentally, in milliseconds on sufficiently large inputs.

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)