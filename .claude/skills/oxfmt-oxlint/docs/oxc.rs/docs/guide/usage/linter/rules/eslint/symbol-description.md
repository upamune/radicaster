---
title: "eslint/symbol-description | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/symbol-description.md for this page in Markdown format

# eslint/symbol-description Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#eslint-symbol-description)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#what-it-does)

Require symbol descriptions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#why-is-this-bad)

The Symbol function may have an optional description.

js

```
var foo = Symbol("some description");

var someString = "some description";
var bar = Symbol(someString);
```

Using `description` promotes easier debugging: when a symbol is logged the description is used:

js

```
var foo = Symbol("some description");

console.log(foo);
// prints - Symbol(some description)
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = Symbol();
```

Examples of **correct** code for this rule:

javascript

```
var foo = Symbol("some description");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "symbol-description": "error"
  }
}
```

bash

```
oxlint --deny symbol-description
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/symbol_description.rs)
