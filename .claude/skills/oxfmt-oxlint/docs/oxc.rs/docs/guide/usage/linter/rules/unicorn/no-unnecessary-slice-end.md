---
title: "unicorn/no-unnecessary-slice-end | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.md for this page in Markdown format

# unicorn/no-unnecessary-slice-end Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#unicorn-no-unnecessary-slice-end)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#what-it-does)

Omitting the end argument defaults it to the object's .length. Passing it explicitly or using Infinity is unnecessary

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#why-is-this-bad)

In JavaScript, omitting the end index already causes .slice() to run to the end of the target, so explicitly passing its length or Infinity is redundant.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#examples)

Examples of **incorrect** code for this rule:

js

```
const foo = string.slice(1, string.length);
const foo = string.slice(1, Infinity);
const foo = string.slice(1, Number.POSITIVE_INFINITY);
```

Examples of **correct** code for this rule:

js

```
const foo = string.slice(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unnecessary-slice-end": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unnecessary-slice-end
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unnecessary_slice_end.rs)
