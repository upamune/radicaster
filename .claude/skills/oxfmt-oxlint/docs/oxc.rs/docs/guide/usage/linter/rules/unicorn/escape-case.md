---
title: "unicorn/escape-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/escape-case.md for this page in Markdown format

# unicorn/escape-case Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#unicorn-escape-case)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#what-it-does)

Enforces defining escape sequence values with uppercase characters rather than lowercase ones. This promotes readability by making the escaped value more distinguishable from the identifier.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#why-is-this-bad)

Using lowercase characters in escape sequences makes them less readable and harder to distinguish from surrounding code. Most style guides recommend uppercase for consistency and clarity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = "\xa9";
const foo = "\ud834";
const foo = "\u{1d306}";
const foo = "\ca";
```

Examples of **correct** code for this rule:

javascript

```
const foo = "\xA9";
const foo = "\uD834";
const foo = "\u{1D306}";
const foo = "\cA";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/escape-case": "error"
  }
}
```

bash

```
oxlint --deny unicorn/escape-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/escape_case.rs)
