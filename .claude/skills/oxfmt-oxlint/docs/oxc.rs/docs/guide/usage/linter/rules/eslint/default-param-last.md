---
title: "eslint/default-param-last | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/default-param-last.md for this page in Markdown format

# eslint/default-param-last Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#eslint-default-param-last)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#what-it-does)

Requires default parameters in functions to be the last ones.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#why-is-this-bad)

Placing default parameters last allows function calls to omit optional trailing arguments, which improves readability and consistency. This rule applies equally to JavaScript and TypeScript functions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#examples)

Examples of **incorrect** code for this rule:

js

```
/* default-param-last: "error" */

function f(a = 0, b) {}
function f(a, b = 0, c) {}
function createUser(isAdmin = false, id) {}
createUser(undefined, "tabby");
```

Examples of **correct** code for this rule:

js

```
/* default-param-last: "error" */

function f(a, b = 0) {}
function f(a = 0, b = 0) {}
function createUser(id, isAdmin = false) {}
createUser("tabby");
```

Examples of **incorrect** TypeScript code for this rule:

ts

```
/* default-param-last: "error" */

function greet(message: string = "Hello", name: string) {}
function combine(a: number = 1, b: number, c: number) {}
function combine(a: number, b: number = 2, c: number) {}
function combine(a: number = 1, b?: number, c: number) {}
```

Examples of **correct** TypeScript code for this rule:

ts

```
/* default-param-last: "error" */

function greet(name: string, message: string = "Hello") {}
function combine(a: number, b: number = 2, c: number = 3) {}
function combine(a: number, b?: number, c: number = 3) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "default-param-last": "error"
  }
}
```

bash

```
oxlint --deny default-param-last
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/default_param_last.rs)
