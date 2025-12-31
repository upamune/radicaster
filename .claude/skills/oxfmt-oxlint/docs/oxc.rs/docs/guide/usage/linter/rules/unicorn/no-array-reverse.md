---
title: "unicorn/no-array-reverse | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-reverse.md for this page in Markdown format

# unicorn/no-array-reverse Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#unicorn-no-array-reverse)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#what-it-does)

Prefer using `Array#toReversed()` over `Array#reverse()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#why-is-this-bad)

`Array#reverse()` modifies the original array in place, which can lead to unintended side effects—especially when the original array is used elsewhere in the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#examples)

Examples of **incorrect** code for this rule:

js

```
const reversed = [...array].reverse();
```

Examples of **correct** code for this rule:

js

```
const reversed = [...array].toReversed();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#configuration)

This rule accepts a configuration object with the following properties:

### allowExpressionStatement [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#allowexpressionstatement)

type: `boolean`

default: `true`

This rule allows `array.reverse()` as an expression statement by default. Set to `false` to forbid `Array#reverse()` even if it's an expression statement.

Examples of **incorrect** code for this rule with this option set to `false`:

js

```
array.reverse();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-reverse": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-reverse
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_reverse.rs)
