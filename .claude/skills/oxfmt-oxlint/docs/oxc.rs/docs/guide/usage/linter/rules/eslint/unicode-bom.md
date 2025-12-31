---
title: "eslint/unicode-bom | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/unicode-bom.md for this page in Markdown format

# eslint/unicode-bom Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#eslint-unicode-bom)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#what-it-does)

Require or disallow Unicode byte order mark (BOM)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#why-is-this-bad)

The Unicode Byte Order Mark (BOM) is used to specify whether code units are big endian or little endian. That is, whether the most significant or least significant bytes come first. UTF-8 does not require a BOM because byte ordering does not matter when characters are a single byte. Since UTF-8 is the dominant encoding of the web, we make "never" the default option.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var a = 123;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#always)

Always require a Unicode BOM (Byte Order Mark) at the beginning of the file.

### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#never)

Never allow a Unicode BOM (Byte Order Mark) at the beginning of the file. This is the default option.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicode-bom": "error"
  }
}
```

bash

```
oxlint --deny unicode-bom
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/unicode_bom.rs)
