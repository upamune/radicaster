---
title: "unicorn/no-unnecessary-array-flat-depth | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.md for this page in Markdown format

# unicorn/no-unnecessary-array-flat-depth Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#unicorn-no-unnecessary-array-flat-depth)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#what-it-does)

Disallows passing `1` to `Array.prototype.flat`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#why-is-this-bad)

Passing `1` is unnecessary.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#examples)

Examples of **incorrect** code for this rule:

js

```
foo.flat(1);
```

Examples of **correct** code for this rule:

js

```
foo.flat();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unnecessary-array-flat-depth": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unnecessary-array-flat-depth
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unnecessary_array_flat_depth.rs)
