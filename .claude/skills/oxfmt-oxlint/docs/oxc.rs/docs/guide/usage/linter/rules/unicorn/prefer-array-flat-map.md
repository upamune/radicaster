---
title: "unicorn/prefer-array-flat-map | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.md for this page in Markdown format

# unicorn/prefer-array-flat-map Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#unicorn-prefer-array-flat-map)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#what-it-does)

Prefers the use of `.flatMap()` when `map().flat()` are used together.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#why-is-this-bad)

It is slightly more efficient to use `.flatMap(…)` instead of `.map(…).flat()`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const bar = [1, 2, 3].map((i) => [i]).flat();
```

Examples of **correct** code for this rule:

javascript

```
const bar = [1, 2, 3].flatMap((i) => [i]);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-array-flat-map": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-array-flat-map
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_array_flat_map.rs)
