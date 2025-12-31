---
title: "unicorn/no-unnecessary-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.md for this page in Markdown format

# unicorn/no-unnecessary-await Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#unicorn-no-unnecessary-await)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#what-it-does)

Disallow awaiting on non-promise values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#why-is-this-bad)

The `await` operator should only be used on `Promise` values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function bad() {
  await await promise;
}
```

Examples of **correct** code for this rule:

javascript

```
async function bad() {
  await promise;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unnecessary-await": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unnecessary-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unnecessary_await.rs)
