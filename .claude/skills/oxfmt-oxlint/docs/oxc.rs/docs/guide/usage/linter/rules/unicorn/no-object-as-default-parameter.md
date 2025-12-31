---
title: "unicorn/no-object-as-default-parameter | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.md for this page in Markdown format

# unicorn/no-object-as-default-parameter Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#unicorn-no-object-as-default-parameter)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#what-it-does)

Disallow the use of an object literal as a default value for a parameter.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#why-is-this-bad)

Default parameters should not be passed to a function through an object literal. The `foo = {a: false}` parameter works fine if only used with one option. As soon as additional options are added, you risk replacing the whole `foo = {a: false, b: true}` object when passing only one option: `{a: true}`. For this reason, object destructuring should be used instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo(foo = { a: false }) {}
```

Examples of **correct** code for this rule:

javascript

```
function foo({ a = false } = {}) {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-object-as-default-parameter": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-object-as-default-parameter
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_object_as_default_parameter.rs)
