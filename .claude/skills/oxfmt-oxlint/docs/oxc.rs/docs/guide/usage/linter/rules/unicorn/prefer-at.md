---
title: "unicorn/prefer-at | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-at.md for this page in Markdown format

# unicorn/prefer-at Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#unicorn-prefer-at)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#what-it-does)

Prefer `.at()` method for index access and `String#charAt()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#why-is-this-bad)

The `.at()` method is more readable and consistent for accessing elements by index, especially for negative indices which access elements from the end.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#examples)

Examples of **incorrect** code for this rule:

js

```
const foo = array[array.length - 1];
const foo = array.slice(-1)[0];
const foo = string.charAt(string.length - 1);
```

Examples of **correct** code for this rule:

js

```
const foo = array.at(-1);
const foo = array.at(-5);
const foo = string.at(-1);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#configuration)

This rule accepts a configuration object with the following properties:

### checkAllIndexAccess [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#checkallindexaccess)

type: `boolean`

default: `false`

Check all index access, not just special patterns like `array.length - 1`. When enabled, `array[0]`, `array[1]`, etc. will also be flagged.

### getLastElementFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#getlastelementfunctions)

type: `string[]`

default: `[]`

List of function names to treat as "get last element" functions. These functions will be checked for `.at(-1)` usage.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-at": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-at
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_at.rs)
