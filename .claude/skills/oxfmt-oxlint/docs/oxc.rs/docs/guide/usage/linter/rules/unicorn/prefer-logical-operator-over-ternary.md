---
title: "unicorn/prefer-logical-operator-over-ternary | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.md for this page in Markdown format

# unicorn/prefer-logical-operator-over-ternary Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#unicorn-prefer-logical-operator-over-ternary)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#what-it-does)

This rule finds ternary expressions that can be simplified to a logical operator.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#why-is-this-bad)

Using a logical operator is shorter and simpler than a ternary expression.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = bar ? bar : baz;
console.log(foo ? foo : bar);
```

Examples of **correct** code for this rule:

javascript

```
const foo = bar || baz;
console.log(foo ?? bar);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-logical-operator-over-ternary": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-logical-operator-over-ternary
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_logical_operator_over_ternary.rs)
