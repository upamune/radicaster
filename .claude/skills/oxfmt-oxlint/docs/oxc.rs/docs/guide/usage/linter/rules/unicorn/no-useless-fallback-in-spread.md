---
title: "unicorn/no-useless-fallback-in-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.md for this page in Markdown format

# unicorn/no-useless-fallback-in-spread Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#unicorn-no-useless-fallback-in-spread)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#what-it-does)

Disallow useless fallback when spreading in object literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#why-is-this-bad)

Spreading [falsy values](https://developer.mozilla.org/en-US/docs/Glossary/Falsy) in object literals won't add any unexpected properties, so it's unnecessary to add an empty object as fallback.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const object = { ...(foo || {}) };
```

Examples of **correct** code for this rule:

javascript

```
const object = { ...foo };
const object = { ...(foo || { not: "empty" }) };
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-fallback-in-spread": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-fallback-in-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_fallback_in_spread.rs)
