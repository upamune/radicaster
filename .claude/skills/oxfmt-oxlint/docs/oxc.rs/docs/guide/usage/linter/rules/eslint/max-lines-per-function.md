---
title: "eslint/max-lines-per-function | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-lines-per-function.md for this page in Markdown format

# eslint/max-lines-per-function Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#eslint-max-lines-per-function)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#what-it-does)

Enforce a maximum number of lines of code in a function. This rule ensures that functions do not exceed a specified line count, promoting smaller, more focused functions that are easier to maintain and understand.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#why-is-this-bad)

Some people consider large functions a code smell. Large functions tend to do a lot of things and can make it hard to follow what’s going on. Many coding style guides dictate a limit to the number of lines that a function can comprise of. This rule can help enforce that style.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#examples)

Examples of **incorrect** code for this rule with a particular max value:

js

```
/* { "eslint/max-lines-per-function": ["error", 2] } */
function foo() {
  const x = 0;
}

/* { "eslint/max-lines-per-function": ["error", 4] } */
function foo() {
  // a comment followed by a blank line

  const x = 0;
}
```

Examples of **correct** code for this rule with a particular max value:

js

```
/* { "eslint/max-lines-per-function": ["error", 3] } */
function foo() {
  const x = 0;
}

/* { "eslint/max-lines-per-function": ["error", 5] } */
function foo() {
  // a comment followed by a blank line

  const x = 0;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#configuration)

This rule accepts a configuration object with the following properties:

### IIFEs [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#iifes)

type: `boolean`

default: `false`

The `IIFEs` option controls whether IIFEs are included in the line count. By default, IIFEs are not considered, but when set to `true`, they will be included in the line count for the function.

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#max)

type: `integer`

default: `50`

Maximum number of lines allowed in a function.

### skipBlankLines [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#skipblanklines)

type: `boolean`

default: `false`

Skip lines made up purely of whitespace.

### skipComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#skipcomments)

type: `boolean`

default: `false`

Skip lines containing just comments.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-lines-per-function": "error"
  }
}
```

bash

```
oxlint --deny max-lines-per-function
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_lines_per_function.rs)
