---
title: "eslint/no-param-reassign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-param-reassign.md for this page in Markdown format

# eslint/no-param-reassign Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#eslint-no-param-reassign)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#what-it-does)

Disallow reassigning function parameters or, optionally, their properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#why-is-this-bad)

Reassigning parameters can lead to unexpected behavior, especially when relying on the original arguments passed into the function. Mutating parameter properties can be similarly surprising and harder to reason about.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#examples)

javascript

```
function foo(bar) {
  bar = 1;
}

function baz(qux) {
  qux.prop = 2; // when `props` option is enabled
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#configuration)

This rule accepts a configuration object with the following properties:

### ignorePropertyModificationsFor [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#ignorepropertymodificationsfor)

type: `string[]`

default: `[]`

An array of parameter names whose property modifications should be ignored.

### ignorePropertyModificationsForRegex [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#ignorepropertymodificationsforregex)

type: `string[]`

An array of regex patterns (as strings) for parameter names whose property modifications should be ignored. Note that this uses [Rust regex syntax](https://docs.rs/regex/latest/regex/) and so may not have all features available to JavaScript regexes.

### props [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#props)

type: `boolean`

default: `false`

When true, also check for modifications to properties of parameters.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-param-reassign": "error"
  }
}
```

bash

```
oxlint --deny no-param-reassign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_param_reassign.rs)
