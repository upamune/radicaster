---
title: "eslint/no-unneeded-ternary | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.md for this page in Markdown format

# eslint/no-unneeded-ternary Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#eslint-no-unneeded-ternary)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#what-it-does)

Disallow ternary operators when simpler alternatives exist

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#why-is-this-bad)

It’s a common mistake in JavaScript to use a conditional expression to select between two Boolean values instead of using ! to convert the test to a Boolean.

Another common mistake is using a single variable as both the conditional test and the consequent. In such cases, the logical OR can be used to provide the same functionality.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#examples)

Examples of **incorrect** code for this rule:

js

```
const isYes = answer === 1 ? true : false;
const isNo = answer === 1 ? false : true;

foo(bar ? bar : 1);
```

Examples of **correct** code for this rule:

js

```
const isYes = answer === 1;
const isNo = answer !== 1;

foo(bar || 1);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#configuration)

This rule accepts a configuration object with the following properties:

### defaultAssignment [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#defaultassignment)

type: `boolean`

default: `true`

Whether to allow the default assignment pattern `x ? x : y`.

When set to `false`, the rule also flags cases like `x ? x : y` and suggests using the logical OR form `x || y` instead. When `true` (default), such default assignments are allowed and not reported.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unneeded-ternary": "error"
  }
}
```

bash

```
oxlint --deny no-unneeded-ternary
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unneeded_ternary.rs)
