# [Step 4 - Whitespace and Basic Tokenization](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside6)

Wouldn’t it be nice to be able to add spaces anywhere?

```javascript
function whitespace() {
  while (input[pos] == ' ' || input[pos] == '\n') pos++
}
function init(s) { input = s; pos = 0; read() }
function read() { whitespace(); peek = input[pos++] }
function next() { let c = peek; read(); return c }
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside6)

Done! We no longer return every character but chose to skip some we want to ignore.

```javascript
print(parse("2 * ( 6 + 4 ) * 5") + " == " + (2*(6+4)*5))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside6)

100 == 100

Our handling of whitespace is an instance of an important pattern: lexical scanning or *tokenization*. Instead of blindly iterating over raw characters, iterate over meaningful chunks of characters (“tokens”).

**Exercise:** add line comments: skip to end of line when seeing `#` or `//`

**Exercise:** add delimited comments: skip parts enclosed between `/*` and `*/`. There is a design choice whether these should be nestable. Try both versions!

Note that the current implementation literally allows spaces *everywhere*, including as part of a number. This may or may not be what we want!

**Exercise:** treat numbers as atomic *tokens*. Hint: modify `number` so that it can be invoked from `read`.

[Solution >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside7)

Tokenization is not essential (“scannerless” parsers exist), but it is a useful design pattern and can be seen as an optimization. In particular it enables distinguishing between keywords (e.g., `if`,`else`, etc.) and identifiers (i.e., potential variable names) without backtracking.

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)