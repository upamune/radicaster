---
title: "unicorn/no-array-sort | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-sort.md for this page in Markdown format

# unicorn/no-array-sort Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#unicorn-no-array-sort)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#what-it-does)

Prefer using `Array#toSorted()` over `Array#sort()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#why-is-this-bad)

`Array#sort()` modifies the original array in place, which can lead to unintended side effects—especially when the original array is used elsewhere in the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#examples)

Examples of **incorrect** code for this rule:

js

```
const sorted = [...array].sort();
```

Examples of **correct** code for this rule:

js

```
const sorted = [...array].toSorted();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#configuration)

This rule accepts a configuration object with the following properties:

### allowExpressionStatement [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#allowexpressionstatement)

type: `boolean`

default: `true`

When set to `true` (default), allows `array.sort()` as an expression statement. Set to `false` to forbid `Array#sort()` even if it's an expression statement.

Example of **incorrect** code for this rule with `allowExpressionStatement` set to `false`:

js

```
array.sort();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-sort": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-sort
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_sort.rs)
