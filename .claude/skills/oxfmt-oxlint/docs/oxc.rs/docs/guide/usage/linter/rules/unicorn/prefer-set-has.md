---
title: "unicorn/prefer-set-has | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-set-has.md for this page in Markdown format

# unicorn/prefer-set-has Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#unicorn-prefer-set-has)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#what-it-does)

Prefer `Set#has()` over `Array#includes()` when checking for existence or non-existence.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#why-is-this-bad)

Set#has() is faster than Array#includes().

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#examples)

Examples of **incorrect** code for this rule:

js

```
const array = [1, 2, 3];
const hasValue = (value) => array.includes(value);
```

Examples of **correct** code for this rule:

js

```
const set = new Set([1, 2, 3]);
const hasValue = (value) => set.has(value);
```

js

```
const array = [1, 2, 3];
const hasOne = array.includes(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-set-has": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-set-has
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_set_has.rs)
