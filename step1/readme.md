# [Step 1 - Arithmetic Expressions and Operator Precedence](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside3)

One thing that’s often thought to be a challenge for hand-written parsers is parsing arithmetic expressions with the proper operator precedence, e.g., parsing `2*6+4*5` as `(2*6)+(4*5)`.

But faster than you can say recursive descent, LL(1), factoring, left recursion, Pratt parsing, etc., I’ll show you how to do it in 10 lines of JavaScript.

Ready? Here you go (imperative style):

```javascript
let input = "2*6+4*5"
let sum = 0
for (let term of input.split("+")) {
  let prod = 1
  for (let factor of term.split("*")) {
    prod *= Number.parseInt(factor)
  }
  sum += prod
}
print(sum + " == " + (2*6+4*5))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside3)

32 == 32

Oh, you prefer a functional style? Bam, 2 lines.

```javascript
let prod = s => s.split("*").map(Number).reduce((x,y)=>x*y)
let sum  = s => s.split("+").map(prod).reduce((x,y)=>x+y)
print(sum("2*6+4*5") + " == " + (2*6+4*5))
```

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside3)

32 == 32

The code snippets above are editable, so feel free to play around!

Clearly this is not the end of the road. But if parsing simple expressions with the right precedence (”`*`” before ”`+`”) is all you wanted to do *you can stop right here!* Nothing else is needed.

Even if you want to go further, it’s important to realize that the basic idea behind operator precedence and in fact most other parsing tasks is *right there*:

> Repeatedly group the input based on delimiter characters

**Exercise:** the code above is editable - change it to build an AST instead of computing the result directly. Hint: use JS arrays to build S-expressions such as `["+", ["*", 2, 6], ["*", 4, 5]]`.

**Exercise:** prove that the input language is a regular language (corollary: no context-free grammar or parsing approach is necessary so far, a deterministic finite automaton (DFA) is sufficient). Hint: provide a regular expression that describes the input language or a DFA that parses it.

Backlinks:

[<< Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)