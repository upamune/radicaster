---
title: "unicorn/prefer-includes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-includes.md for this page in Markdown format

# unicorn/prefer-includes Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#unicorn-prefer-includes)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#what-it-does)

Prefer `includes()` over `indexOf()` when checking for existence or non-existence. All built-ins have `.includes()` in addition to `.indexOf()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#why-is-this-bad)

The `.includes()` method is more readable and less error-prone than `.indexOf()`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (str.indexOf("foo") !== -1) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (str.includes("foo")) {
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-includes": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-includes
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_includes.rs)
