---
title: "eslint/no-warning-comments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-warning-comments.md for this page in Markdown format

# eslint/no-warning-comments Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#eslint-no-warning-comments)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#what-it-does)

Disallows warning comments such as TODO, FIXME, XXX in code.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#why-is-this-bad)

Developers often add comments like TODO or FIXME to mark incomplete work or areas that need attention. While useful during development, these comments can indicate unfinished code that shouldn't be shipped to production. This rule helps catch such comments before they make it into production code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// TODO: implement this feature
function doSomething() {}

// FIXME: this is broken
const x = 1;

/* XXX: hack */
let y = 2;
```

Examples of **correct** code for this rule:

javascript

```
// This is a regular comment
function doSomething() {}

// Note: This explains something
const x = 1;
```

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#options)

This rule has an options object with the following defaults:

json

```
{
  "terms": ["todo", "fixme", "xxx"],
  "location": "start",
  "decoration": []
}
```

#### `terms` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#terms)

An array of terms to match. The matching is case-insensitive.

#### `location` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#location)

Where to check for the terms:

* `"start"` (default): Terms must appear at the start of the comment (after any decoration)
* `"anywhere"`: Terms can appear anywhere in the comment

#### `decoration` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#decoration)

An array of characters to ignore at the start of comments when `location` is `"start"`. Useful for ignoring common comment decorations like `*` in JSDoc-style comments.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-warning-comments": "error"
  }
}
```

bash

```
oxlint --deny no-warning-comments
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_warning_comments.rs)
