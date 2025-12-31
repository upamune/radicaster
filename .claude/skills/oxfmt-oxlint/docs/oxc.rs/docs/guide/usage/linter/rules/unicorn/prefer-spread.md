---
title: "unicorn/prefer-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-spread.md for this page in Markdown format

# unicorn/prefer-spread Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#unicorn-prefer-spread)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#what-it-does)

Enforces the use of [the spread operator (`...`)](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Spread_syntax) over outdated patterns.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#why-is-this-bad)

Using the spread operator is more concise and readable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = Array.from(set);
const foo = Array.from(new Set([1, 2]));
```

Examples of **correct** code for this rule:

javascript

```
[...set].map(() => {});
Array.from(...argumentsArray);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-spread": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_spread.rs)
