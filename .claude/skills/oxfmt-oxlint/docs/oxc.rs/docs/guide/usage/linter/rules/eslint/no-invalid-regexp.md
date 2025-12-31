---
title: "eslint/no-invalid-regexp | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-invalid-regexp.md for this page in Markdown format

# eslint/no-invalid-regexp Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#eslint-no-invalid-regexp)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#what-it-does)

Disallow invalid regular expression strings in RegExp constructors.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#why-is-this-bad)

An invalid pattern in a regular expression literal is a SyntaxError when the code is parsed, but an invalid string in RegExp constructors throws a SyntaxError only when the code is executed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#examples)

Examples of **incorrect** code for this rule:

js

```
RegExp("[");
RegExp(".", "z");
new RegExp("\\");
```

Examples of **correct** code for this rule:

js

```
RegExp(".");
new RegExp();
this.RegExp("[");
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#configuration)

This rule accepts a configuration object with the following properties:

### allowConstructorFlags [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#allowconstructorflags)

type: `string[]`

default: `[]`

Case-sensitive array of flags that will be allowed.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-invalid-regexp": "error"
  }
}
```

bash

```
oxlint --deny no-invalid-regexp
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_invalid_regexp.rs)
