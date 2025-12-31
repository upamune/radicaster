---
title: "unicorn/no-instanceof-array | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-instanceof-array.md for this page in Markdown format

# unicorn/no-instanceof-array Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#unicorn-no-instanceof-array)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#what-it-does)

Require `Array.isArray()` instead of `instanceof Array`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#why-is-this-bad)

The instanceof Array check doesn't work across realms/contexts, for example, frames/windows in browsers or the vm module in Node.js.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
array instanceof Array;
[1, 2, 3] instanceof Array;
```

Examples of **correct** code for this rule:

javascript

```
Array.isArray(array);
Array.isArray([1, 2, 3]);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-instanceof-array": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-instanceof-array
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_instanceof_array.rs)
