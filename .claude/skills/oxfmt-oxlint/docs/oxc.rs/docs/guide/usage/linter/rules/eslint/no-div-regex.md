---
title: "eslint/no-div-regex | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-div-regex.md for this page in Markdown format

# eslint/no-div-regex Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#eslint-no-div-regex)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#what-it-does)

Disallow equal signs explicitly at the beginning of regular expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#why-is-this-bad)

Characters /= at the beginning of a regular expression literal can be confused with a division assignment operator.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function bar() {
  return /=foo/;
}
```

Examples of **correct** code for this rule:

javascript

```
function bar() {
  return /[=]foo/;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-div-regex": "error"
  }
}
```

bash

```
oxlint --deny no-div-regex
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_div_regex.rs)
