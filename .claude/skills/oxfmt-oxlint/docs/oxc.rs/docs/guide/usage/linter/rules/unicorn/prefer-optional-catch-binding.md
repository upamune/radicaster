---
title: "unicorn/prefer-optional-catch-binding | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.md for this page in Markdown format

# unicorn/prefer-optional-catch-binding Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#unicorn-prefer-optional-catch-binding)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#what-it-does)

Prefers omitting the catch binding parameter if it is unused

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#why-is-this-bad)

It is unnecessary to bind the error to a variable if it is not used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
try {
  // ...
} catch (e) {}
```

Examples of **correct** code for this rule:

javascript

```
try {
  // ...
} catch {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-optional-catch-binding": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-optional-catch-binding
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_optional_catch_binding.rs)
