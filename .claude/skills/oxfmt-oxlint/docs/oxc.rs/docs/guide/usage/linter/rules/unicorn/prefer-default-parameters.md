---
title: "unicorn/prefer-default-parameters | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.md for this page in Markdown format

# unicorn/prefer-default-parameters Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#unicorn-prefer-default-parameters)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#what-it-does)

Instead of reassigning a function parameter, default parameters should be used. The `foo = foo || 123` statement evaluates to `123` when `foo` is falsy, possibly leading to confusing behavior, whereas default parameters only apply when passed an `undefined` value. This rule only reports reassignments to literal values.

You should disable this rule if you want your functions to deal with `null` and other falsy values the same way as `undefined`. Default parameters are exclusively applied [when `undefined` is received.](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/Default_parameters#passing_undefined_vs._other_falsy_values). However, we recommend [moving away from `null`](https://github.com/sindresorhus/meta/discussions/7).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#why-is-this-bad)

Using default parameters makes it clear that a parameter has a default value, improving code readability and maintainability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#examples)

Examples of **incorrect** code for this rule:

js

```
function abc(foo) {
  foo = foo || "bar";
}

function abc(foo) {
  const bar = foo || "bar";
}
```

Examples of **correct** code for this rule:

js

```
function abc(foo = "bar") {}

function abc(bar = "bar") {}

function abc(foo) {
  foo = foo || bar();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-default-parameters": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-default-parameters
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_default_parameters.rs)
