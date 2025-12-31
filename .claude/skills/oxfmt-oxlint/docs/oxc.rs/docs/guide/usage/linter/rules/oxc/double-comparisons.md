---
title: "oxc/double-comparisons | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/double-comparisons.md for this page in Markdown format

# oxc/double-comparisons Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#oxc-double-comparisons)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#what-it-does)

This rule checks for double comparisons in logical expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#why-is-this-bad)

Redundant comparisons can be confusing and make code harder to understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
x === y || x < y;
x < y || x === y;
```

Examples of **correct** code for this rule:

javascript

```
x <= y;
x >= y;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/double-comparisons": "error"
  }
}
```

bash

```
oxlint --deny oxc/double-comparisons
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/double_comparisons.rs)
