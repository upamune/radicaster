---
title: "unicorn/catch-error-name | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/catch-error-name.md for this page in Markdown format

# unicorn/catch-error-name Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#unicorn-catch-error-name)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#what-it-does)

This rule enforces consistent and descriptive naming for error variables in `catch` statements, preventing the use of vague names like `badName` or `_` when the error is used.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#why-is-this-bad)

Using non-descriptive names like `badName` or `_` makes the code harder to read and understand, especially when debugging. It's important to use clear, consistent names to represent errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
try {
} catch (badName) {}

// `_` is not allowed if it's used
try {
} catch (_) {
  console.log(_);
}

promise.catch((badName) => {});

promise.then(undefined, (badName) => {});
```

Examples of **correct** code for this rule:

javascript

```
try {
} catch (error) {}

// `_` is allowed if it's not used
try {
} catch (_) {
  console.log(123);
}

promise.catch((error) => {});

promise.then(undefined, (error) => {});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#configuration)

This rule accepts a configuration object with the following properties:

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#ignore)

type: `string[]`

A list of patterns to ignore when checking `catch` variable names. The pattern can be a string or regular expression.

### name [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#name)

type: `string`

default: `"error"`

The name to use for error variables in `catch` blocks. You can customize it to something other than `'error'` (e.g., `'exception'`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/catch-error-name": "error"
  }
}
```

bash

```
oxlint --deny unicorn/catch-error-name
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/catch_error_name.rs)
