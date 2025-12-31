---
title: "oxc/misrefactored-assign-op | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.md for this page in Markdown format

# oxc/misrefactored-assign-op Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#oxc-misrefactored-assign-op)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#what-it-does)

<https://rust-lang.github.io/rust-clippy/master/#/misrefactored_assign_op>

Checks for `a op= a op b` or `a op= b op a` patterns.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#why-is-this-bad)

Most likely these are bugs where one meant to write `a op= b`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
a += a + b;
a -= a - b;
```

Examples of **correct** code for this rule:

javascript

```
a += b;
a -= b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/misrefactored-assign-op": "error"
  }
}
```

bash

```
oxlint --deny oxc/misrefactored-assign-op
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/misrefactored_assign_op.rs)
