---
title: "unicorn/no-zero-fractions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-zero-fractions.md for this page in Markdown format

# unicorn/no-zero-fractions Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#unicorn-no-zero-fractions)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#what-it-does)

Prevents the use of zero fractions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#why-is-this-bad)

There is no difference in JavaScript between, for example, `1`, `1.0` and `1.`, so prefer the former for consistency and brevity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = 1.0;
const foo = -1.0;
const foo = 123_456.000_000;
```

Examples of **correct** code for this rule:

javascript

```
const foo = 1;
const foo = -1;
const foo = 123456;
const foo = 1.1;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-zero-fractions": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-zero-fractions
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_zero_fractions.rs)
