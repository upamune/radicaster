---
title: "eslint/prefer-rest-params | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-rest-params.md for this page in Markdown format

# eslint/prefer-rest-params Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#eslint-prefer-rest-params)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#what-it-does)

Disallows the use of the `arguments` object and instead enforces the use of rest parameters.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#why-is-this-bad)

The `arguments` object does not have methods from `Array.prototype`, making it inconvenient for array-like operations. Using rest parameters provides a more intuitive and efficient way to handle variadic arguments.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo() {
  console.log(arguments);
}

function foo(action) {
  var args = Array.prototype.slice.call(arguments, 1);
  action.apply(null, args);
}

function foo(action) {
  var args = [].slice.call(arguments, 1);
  action.apply(null, args);
}
```

Examples of **correct** code for this rule:

javascript

```
function foo(...args) {
  console.log(args);
}

function foo(action, ...args) {
  action.apply(null, args); // Or use `action(...args)` (related to `prefer-spread` rule).
}

// Note: Implicit `arguments` can be shadowed.
function foo(arguments) {
  console.log(arguments); // This refers to the first argument.
}
function foo() {
  var arguments = 0;
  console.log(arguments); // This is a local variable.
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-rest-params": "error"
  }
}
```

bash

```
oxlint --deny prefer-rest-params
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_rest_params.rs)
