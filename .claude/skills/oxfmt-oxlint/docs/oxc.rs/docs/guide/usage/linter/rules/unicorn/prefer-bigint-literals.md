---
title: "unicorn/prefer-bigint-literals | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.md for this page in Markdown format

# unicorn/prefer-bigint-literals Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#unicorn-prefer-bigint-literals)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#what-it-does)

Requires using BigInt literals (e.g. `123n`) instead of calling the `BigInt()` constructor with literal arguments such as numbers or numeric strings

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#why-is-this-bad)

Using `BigInt(…)` with literal values is unnecessarily verbose and less idiomatic than using a BigInt literal.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#examples)

Examples of **incorrect** code for this rule:

js

```
BigInt(0);
BigInt(123);
BigInt(0xff);
BigInt(1e3);
BigInt("42");
BigInt("0x10");
```

Examples of **correct** code for this rule:

js

```
0n;
123n;
0xffn;
1000n;
// Non-integer, dynamic, or non-literal input:
BigInt(x);
BigInt("not-a-number");
BigInt("1.23");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-bigint-literals": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-bigint-literals
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_bigint_literals.rs)
