---
title: "unicorn/prefer-string-raw | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-string-raw.md for this page in Markdown format

# unicorn/prefer-string-raw Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#unicorn-prefer-string-raw)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#what-it-does)

Prefers use of String.raw to avoid escaping .

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#why-is-this-bad)

Excessive backslashes can make string values less readable which can be avoided by using `String.raw`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const file = "C:\\windows\\style\\path\\to\\file.js";
const regexp = new RegExp("foo\\.bar");
```

Examples of **correct** code for this rule:

javascript

```
const file = String.raw`C:\windows\style\path\to\file.js`;
const regexp = new RegExp(String.raw`foo\.bar`);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-string-raw": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-string-raw
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_string_raw.rs)
