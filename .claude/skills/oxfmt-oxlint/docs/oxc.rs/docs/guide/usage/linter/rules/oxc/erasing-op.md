---
title: "oxc/erasing-op | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/erasing-op.md for this page in Markdown format

# oxc/erasing-op Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#oxc-erasing-op)

✅ This rule is turned on by default.

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#what-it-does)

Checks for erasing operations, e.g., `x \* 0``.

Based on <https://rust-lang.github.io/rust-clippy/master/#/erasing_op>

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#why-is-this-bad)

The whole expression can be replaced by zero. This is most likely not the intended outcome and should probably be corrected.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let x = 1;
let y = x * 0;
```

Examples of **correct** code for this rule:

javascript

```
let x = 1;
let y = 0;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/erasing-op": "error"
  }
}
```

bash

```
oxlint --deny oxc/erasing-op
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/erasing_op.rs)
