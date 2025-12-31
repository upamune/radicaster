---
title: "eslint/radix | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/radix.md for this page in Markdown format

# eslint/radix Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#eslint-radix)

⚠️🛠️️ A dangerous auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#what-it-does)

Enforce the consistent use of the radix argument when using `parseInt()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#why-is-this-bad)

Using the `parseInt()` function without specifying the radix can lead to unexpected results.

See the [MDN documentation](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/parseInt#radix) for more information.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var num = parseInt("071"); // 57
```

Examples of **correct** code for this rule:

javascript

```
var num = parseInt("071", 10); // 71
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#always)

Always require the radix parameter when using `parseInt()`.

### `"as-needed"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#as-needed)

Only require the radix parameter when necessary.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "radix": "error"
  }
}
```

bash

```
oxlint --deny radix
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/radix.rs)
