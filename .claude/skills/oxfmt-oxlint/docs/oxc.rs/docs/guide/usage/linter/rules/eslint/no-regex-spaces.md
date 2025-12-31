---
title: "eslint/no-regex-spaces | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-regex-spaces.md for this page in Markdown format

# eslint/no-regex-spaces Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#eslint-no-regex-spaces)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#what-it-does)

Disallow 2+ consecutive spaces in regular expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#why-is-this-bad)

In a regular expression, it is hard to tell how many spaces are intended to be matched. It is better to use only one space and then specify how many spaces are expected using a quantifier.

javascript

```
var re = /foo {3}bar/;
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var re = /foo   bar/;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-regex-spaces": "error"
  }
}
```

bash

```
oxlint --deny no-regex-spaces
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_regex_spaces.rs)
