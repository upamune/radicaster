---
title: "promise/param-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/param-names.md for this page in Markdown format

# promise/param-names Style [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#promise-param-names)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#what-it-does)

Enforce standard parameter names for Promise constructors.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#why-is-this-bad)

Ensures that new Promise() is instantiated with the parameter names resolve, reject to avoid confusion with order such as reject, resolve. The Promise constructor uses the RevealingConstructor pattern. Using the same parameter names as the language specification makes code more uniform and easier to understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
new Promise(function (reject, resolve) {
  /* ... */
}); // incorrect order
new Promise(function (ok, fail) {
  /* ... */
}); // non-standard parameter names
```

Examples of **correct** code for this rule:

javascript

```
new Promise(function (resolve, reject) {});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#configuration)

This rule accepts a configuration object with the following properties:

### rejectPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#rejectpattern)

type: `string | null`

Regex pattern used to validate the `reject` parameter name. If provided, this pattern is used instead of the default `^_?reject$` check.

### resolvePattern [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#resolvepattern)

type: `string | null`

Regex pattern used to validate the `resolve` parameter name. If provided, this pattern is used instead of the default `^_?resolve$` check.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/param-names": "error"
  }
}
```

bash

```
oxlint --deny promise/param-names --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/param_names.rs)
