# [Step 6 - Error Handling and Recovery](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

Right now our parser immediately throws an exception and aborts when encountering an error in the program.

This can be a reasonable strategy, but it’s often more useful to see if there are additional errors in a program so that the programmer can decide which one to fix first (clearly, there are trade-offs!).

A useful strategy is to skip input until the next known delimiter whenever we hit an error:

```javascript
function error(msg) {
  while (peek && peek != '\n' && peek != '*' && peek != '+' && peek != ')')
    next()
  throw new Error(msg)
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

We now want to unwind the stack as much as necessary to pick up the right case (`*` or `+` or `)` etc.)

To achieve this, the code that deals with delimiters is modified to try proceeding normally after an error:

```javascript
function split(d,f) {
  for (;;) {
    try { f() } catch (ex) { }
    // we know that we stopped at a delimiter,
    // let's see if it was possibly ours
    if (peek == d) next(); else break
  }
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

Note how this strategy maps very closely to standard exception handling patterns in a recursive-descent context.

For cases where we expect a specific token such as a closing parenthesis, we can implement a helper function `expect` that tries to skip ahead if the token is not found immediately and either continues if the token is found in this way or propagates the error if a non-matching delimiter was found instead:

```javascript
function expect(d) {
  if (peek == d) {
    next()
  } else {
    try { error(d+" expected") }
    catch (ex) {
      if (peek == d) next(); else throw ex
    }
  }
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

Another missing piece is that we might hit an error after parsing a perfectly valid prefix of the program. This means that we terminate early with some input left over. To deal with such situations, we need to insert pieces of retry logic in critical places, including at the very top:

```javascript
function parse(text) {
  init(text)
  expr(indent)
  let max = 20
  while (peek && max--) {
    try { error("unexpected input:"+peek) } catch(ex) {};
    if (isDelimiter(peek)) next()
    expr(indent)
  }
  if (peek) error("unexpected input:"+peek) // exceeded max, skip rest
}
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

In general, it’s a good idea to impose an upper bound on the number of such retries.

These forms of error handling are all included in the example on the front page of this post.

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)