# [Just write the #!%/* parser](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)

Writing parsers from scratch. Why simpler is better and why you don't need a parser generator.

This is a whirlwind tour of writing parsers by hand. Why would you want to do that, when tools like Yacc exist to do it for you?

- It’s highly instructive, in a way that using a parser generator is not. To quote Feynman: “What I cannot create, I do not understand”
- It’s an important skill: most real-world compilers use hand-written parsers because they provide more control over error handling, significant whitespace, etc.
- It’s not actually difficult!

If you’re an educator, here are some additional thoughts:

[The role of parsing in compiler classes >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside1)

### The Result

We focus on arithmetic expressions, but the concepts and techniques immediately generalize to richer languages.

Enter an expression on the left and see the parse tree change!

[code >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside2)

[done](https://tiarkrompf.github.io/notes/?/just-write-the-parser/)

```
3+4*5*(1+2)
```

```
+ "3+4*5*(1+2)"
	num "3"
	* "4*5*(1+2)"
		num "4"
		num "5"
		+ "1+2"
			num "1"
			num "2"
```

### Step by Step

[Step 1 - Arithmetic Expressions and Operator Precedence >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside3)

[Step 2 - Parsing in a Single Pass >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside4)

[Step 3 - Parentheses and Recursive Grouping >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside5)

[Step 4 - Whitespace and Basic Tokenization >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside6)

[Step 5 - Significant Whitespace >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside8)

[Step 6 - Error Handling and Recovery >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside9)

[Step 7 - Scaling Up >>](https://tiarkrompf.github.io/notes/?/just-write-the-parser/aside10)

Backlinks:

[<< Notes - Start Page](https://tiarkrompf.github.io/notes/?/public/)