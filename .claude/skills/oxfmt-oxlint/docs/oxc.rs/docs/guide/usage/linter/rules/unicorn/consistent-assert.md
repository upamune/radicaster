---
title: "unicorn/consistent-assert | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/consistent-assert.md for this page in Markdown format

# unicorn/consistent-assert Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#unicorn-consistent-assert)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#what-it-does)

Enforces consistent usage of the `assert` module.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#why-is-this-bad)

Inconsistent usage of the `assert` module can lead to confusion and errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#examples)

Examples of **incorrect** code for this rule:

js

```
import assert from "node:assert";

assert(divide(10, 2) === 5);
```

Examples of **correct** code for this rule:

js

```
import assert from "node:assert";

assert.ok(divide(10, 2) === 5);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/consistent-assert": "error"
  }
}
```

bash

```
oxlint --deny unicorn/consistent-assert
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/consistent_assert.rs)
