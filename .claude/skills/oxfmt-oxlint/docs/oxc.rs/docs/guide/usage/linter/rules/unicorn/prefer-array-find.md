---
title: "unicorn/prefer-array-find | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-array-find.md for this page in Markdown format

# unicorn/prefer-array-find Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#unicorn-prefer-array-find)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#what-it-does)

Encourages using `Array.prototype.find` instead of `filter(...)[0]` or similar patterns when only the first matching element is needed.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#why-is-this-bad)

Using `filter(...)[0]` to get the first match is less efficient and more verbose than using `find(...)`. `find` short-circuits when a match is found, whereas `filter` evaluates the entire array.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#examples)

Examples of **incorrect** code for this rule:

js

```
const match = users.filter((u) => u.id === id)[0];
const match = users.filter(fn).shift();
```

Examples of **correct** code for this rule:

js

```
const match = users.find((u) => u.id === id);
const match = users.find(fn);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-array-find": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-array-find
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_array_find.rs)
