---
title: "eslint/no-inner-declarations | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-inner-declarations.md for this page in Markdown format

# eslint/no-inner-declarations Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#eslint-no-inner-declarations)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#what-it-does)

Disallow variable or function declarations in nested blocks

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#why-is-this-bad)

A variable declaration is permitted anywhere a statement can go, even nested deeply inside other blocks. This is often undesirable due to variable hoisting, and moving declarations to the root of the program or function body can increase clarity. Note that block bindings (let, const) are not hoisted and therefore they are not affected by this rule.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (test) {
  function doSomethingElse() {}
}
```

Examples of **correct** code for this rule:

javascript

```
function doSomethingElse() {}
if (test) {
  // your code here
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#configuration)

This rule accepts a configuration object with the following properties:

### blockScopedFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#blockscopedfunctions)

type: `"allow" | "disallow"`

default: `null`

Controls whether function declarations in nested blocks are allowed in strict mode (ES6+ behavior).

#### `"allow"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#allow)

Allow function declarations in nested blocks in strict mode (ES6+ behavior).

#### `"disallow"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#disallow)

Disallow function declarations in nested blocks regardless of strict mode.

### config [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#config)

type: `"functions" | "both"`

default: `"functions"`

Determines what type of declarations to check.

#### `"functions"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#functions)

Disallows function declarations in nested blocks.

#### `"both"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#both)

Disallows function and var declarations in nested blocks.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-inner-declarations": "error"
  }
}
```

bash

```
oxlint --deny no-inner-declarations
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_inner_declarations.rs)
