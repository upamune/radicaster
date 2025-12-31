---
title: "unicorn/consistent-existence-index-check | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.md for this page in Markdown format

# unicorn/consistent-existence-index-check Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#unicorn-consistent-existence-index-check)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#what-it-does)

Enforce consistent style for element existence checks with `indexOf()`, `lastIndexOf()`, `findIndex()`, and `findLastIndex()`. This ensures that comparisons are performed in a standard and clear way.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#why-is-this-bad)

This rule is meant to enforce a specific style and improve code clarity. Using inconsistent comparison styles (e.g., `index < 0`, `index >= 0`) can make the intention behind the code unclear, especially in large codebases.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const index = foo.indexOf("bar");
if (index < 0) {
}

const index = foo.indexOf("bar");
if (index >= 0) {
}
```

Examples of **correct** code for this rule:

javascript

```
const index = foo.indexOf("bar");
if (index === -1) {
}

const index = foo.indexOf("bar");
if (index !== -1) {
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/consistent-existence-index-check": "error"
  }
}
```

bash

```
oxlint --deny unicorn/consistent-existence-index-check
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/consistent_existence_index_check.rs)
