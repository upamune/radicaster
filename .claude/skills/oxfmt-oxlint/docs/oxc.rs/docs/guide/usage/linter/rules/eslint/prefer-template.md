---
title: "eslint/prefer-template | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-template.md for this page in Markdown format

# eslint/prefer-template Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#eslint-prefer-template)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#what-it-does)

Require template literals instead of string concatenation.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#why-is-this-bad)

In ES2015 (ES6), we can use template literals instead of string concatenation.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#examples)

Examples of **incorrect** code for this rule:

js

```
const str = "Hello, " + name + "!";
const str1 = "Time: " + 12 * 60 * 60 * 1000;
```

Examples of **correct** code for this rule:

js

```
const str = "Hello World!";
const str2 = `Time: ${12 * 60 * 60 * 1000}`;
const str4 = "Hello, " + "World!";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-template": "error"
  }
}
```

bash

```
oxlint --deny prefer-template
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_template.rs)
