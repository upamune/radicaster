---
title: "unicorn/prefer-set-size | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-set-size.md for this page in Markdown format

# unicorn/prefer-set-size Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#unicorn-prefer-set-size)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#what-it-does)

Prefer `Set#size` over `Set#length` when the `Set` is converted to an array.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#why-is-this-bad)

Using `Set#size` is more readable and performant.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const length = [...new Set([1, 2, 3])].length;
```

Examples of **correct** code for this rule:

javascript

```
const size = new Set([1, 2, 3]).size;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-set-size": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-set-size
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_set_size.rs)
