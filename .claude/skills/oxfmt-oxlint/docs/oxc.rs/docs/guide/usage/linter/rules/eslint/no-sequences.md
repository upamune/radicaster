---
title: "eslint/no-sequences | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-sequences.md for this page in Markdown format

# eslint/no-sequences Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#eslint-no-sequences)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#what-it-does)

Disallows the use of the comma operator.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#why-is-this-bad)

The comma operator evaluates each of its operands (from left to right) and returns the value of the last operand. However, this frequently obscures side effects, and its use is often an accident.

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#options)

* `allowInParentheses` (default: `true`): If set to `false`, disallows the comma operator even when wrapped in parentheses.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
((foo = doSomething()), val);

(0, eval("doSomething();"));

// Arrow function body needs double parentheses
const fn = () => (doSomething(), val);

// with allowInParentheses: false
foo = (doSomething(), val);
```

Examples of **correct** code for this rule:

javascript

```
foo = (doSomething(), val);

(0, eval)("doSomething();");

// Single extra parentheses is enough for conditions
do {} while ((doSomething(), !!test));

for (i = 0, j = 10; i < j; i++, j--) {}

// Arrow function body needs double parentheses
const fn = () => (doSomething(), val);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#configuration)

This rule accepts a configuration object with the following properties:

### allowInParentheses [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#allowinparentheses)

type: `boolean`

default: `true`

If this option is set to `false`, this rule disallows the comma operator even when the expression sequence is explicitly wrapped in parentheses. Default is `true`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-sequences": "error"
  }
}
```

bash

```
oxlint --deny no-sequences
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_sequences.rs)
