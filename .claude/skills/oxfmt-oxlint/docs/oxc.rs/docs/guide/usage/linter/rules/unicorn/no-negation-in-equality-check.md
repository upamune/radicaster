---
title: "unicorn/no-negation-in-equality-check | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.md for this page in Markdown format

# unicorn/no-negation-in-equality-check Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#unicorn-no-negation-in-equality-check)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#what-it-does)

Disallow negated expressions on the left of (in)equality checks.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#why-is-this-bad)

A negated expression on the left of an (in)equality check is likely a mistake from trying to negate the whole condition.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (!foo === bar) {
}

if (!foo !== bar) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (foo !== bar) {
}

if (!(foo === bar)) {
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-negation-in-equality-check": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-negation-in-equality-check
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_negation_in_equality_check.rs)
