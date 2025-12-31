---
title: "eslint/no-script-url | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-script-url.md for this page in Markdown format

# eslint/no-script-url Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#eslint-no-script-url)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#what-it-does)

Disallow javascript: urls

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#why-is-this-bad)

Using `javascript:` URLs is considered by some as a form of `eval`. Code passed in `javascript:` URLs must be parsed and evaluated by the browser in the same way that `eval` is processed. This can lead to security and performance issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#examples)

Examples of **incorrect** code for this rule

javascript

```
/*eslint no-script-url: "error"*/

location.href = "javascript:void(0)";

location.href = `javascript:void(0)`;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-script-url": "error"
  }
}
```

bash

```
oxlint --deny no-script-url
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_script_url.rs)
