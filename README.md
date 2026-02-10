# minterpreter
a math intepreter written in Go with common interpreter/compiler design principles

## why?
i have been researching compilers on and off for a little while, and i figured i would implement something akin to one, but a bit simpler.
i figured a mathematical intepreter would be a good choice, since it's easier to evaluate math than to generate assembly (i hope).
i also ran into this [excellent blog-series](https://ruslanspivak.com/lsbasi-part2/) by Ruslan Spivak.
this isn't meant to necessarily be production-ready, but more of a learning experience for me.

## design principles
with minterpreter, i want to focus on adhering to compiler/interpreter design. 
specifically, i want to utilize certain algorithms and data structures to make my interpreter robust and extensible in the future. 
minterpreter will make use of the following:
- lexical analysis/tokens
- pratt parsing
- abstract syntax tree

## todo
- [x] support add, sub, mul, and div
- [x] support numbers of arbitrary length
- [x] write tests
- [ ] support parenthesis/operator precedence
- [ ] add exponents
- [ ] add trig functions
- [ ] add numerical constants (pi, e, etc.)
- [ ] create stable and acessible api

## license
this project is distributed under the Lesser GNU Public License v3. see `LICENSE`
