---
title: "unicorn/no-hex-escape | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-hex-escape.md for this page in Markdown format

# unicorn/no-hex-escape Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#unicorn-no-hex-escape)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#what-it-does)

Enforces a convention of using [Unicode escapes](https://mathiasbynens.be/notes/javascript-escapes#unicode) instead of [hexadecimal escapes](https://mathiasbynens.be/notes/javascript-escapes#hexadecimal) for consistency and clarity.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#why-is-this-bad)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = "\x1B";
const foo = `\x1B${bar}`;
```

Examples of **correct** code for this rule:

javascript

```
const foo = "\u001B";
const foo = `\u001B${bar}`;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-hex-escape": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-hex-escape
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_hex_escape.rs)
