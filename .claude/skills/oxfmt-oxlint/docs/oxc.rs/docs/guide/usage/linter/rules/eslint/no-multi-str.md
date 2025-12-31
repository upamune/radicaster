---
title: "eslint/no-multi-str | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-multi-str.md for this page in Markdown format

# eslint/no-multi-str Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#eslint-no-multi-str)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#what-it-does)

Disallow multiline strings.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#why-is-this-bad)

Some consider this to be a bad practice as it was an undocumented feature of JavaScript that was only formalized later.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x =
  "Line 1 \
 Line 2";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-multi-str": "error"
  }
}
```

bash

```
oxlint --deny no-multi-str
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_multi_str.rs)
