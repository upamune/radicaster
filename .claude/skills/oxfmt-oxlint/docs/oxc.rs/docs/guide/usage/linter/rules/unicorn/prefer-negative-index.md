---
title: "unicorn/prefer-negative-index | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-negative-index.md for this page in Markdown format

# unicorn/prefer-negative-index Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#unicorn-prefer-negative-index)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#what-it-does)

Prefer negative index over `.length` - index when possible

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#why-is-this-bad)

Conciseness and readability

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#examples)

Examples of **incorrect** code for this rule:

js

```
foo.slice(foo.length - 2, foo.length - 1);
foo.at(foo.length - 1);
```

Examples of **correct** code for this rule:

js

```
foo.slice(-2, -1);
foo.at(-1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-negative-index": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-negative-index
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_negative_index.rs)
