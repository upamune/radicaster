---
title: "oxc/bad-comparison-sequence | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.md for this page in Markdown format

# oxc/bad-comparison-sequence Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#oxc-bad-comparison-sequence)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#what-it-does)

This rule applies when the comparison operator is applied two or more times in a row.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#why-is-this-bad)

Because comparison operator is a binary operator, it is impossible to compare three or more operands at once. If comparison operator is used to compare three or more operands, only the first two operands are compared and the rest is compared with its result of boolean type.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if ((a == b) == c) {
  console.log("a, b, and c are the same");
}
```

Examples of **correct** code for this rule:

javascript

```
if (a == b && b == c) {
  console.log("a, b, and c are the same");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-comparison-sequence": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-comparison-sequence
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_comparison_sequence.rs)
