---
title: "eslint/no-plusplus | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-plusplus.md for this page in Markdown format

# eslint/no-plusplus Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#eslint-no-plusplus)

💡 A suggestion is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#what-it-does)

Disallow the unary operators `++` and `--`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#why-is-this-bad)

Because the unary `++` and `--` operators are subject to automatic semicolon insertion, differences in whitespace can change the semantics of source code. For example, these two code blocks are not equivalent:

js

```
var i = 10;
var j = 20;

i++;
j;
// => i = 11, j = 20
```

js

```
var i = 10;
var j = 20;

i;
++j;
// => i = 10, j = 21
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#examples)

Examples of **incorrect** code for this rule:

js

```
var x = 0;
x++;
var y = 0;
y--;
for (let i = 0; i < l; i++) {
  doSomething(i);
}
```

Examples of **correct** code for this rule:

js

```
var x = 0;
x += 1;
var y = 0;
y -= 1;
for (let i = 0; i < l; i += 1) {
  doSomething(i);
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#configuration)

This rule accepts a configuration object with the following properties:

### allowForLoopAfterthoughts [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#allowforloopafterthoughts)

type: `boolean`

default: `false`

Whether to allow `++` and `--` in for loop afterthoughts.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-plusplus": "error"
  }
}
```

bash

```
oxlint --deny no-plusplus
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_plusplus.rs)
