---
title: "eslint/no-unused-expressions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unused-expressions.md for this page in Markdown format

# eslint/no-unused-expressions Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#eslint-no-unused-expressions)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#what-it-does)

This rule disallows unused expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#why-is-this-bad)

Unused expressions are usually a mistake. They can be a symptom of a bug or a misunderstanding of the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#examples)

Examples of **incorrect** code for this rule:

ts

```
Set<number>;
1 as number;
window!;
```

Examples of **correct** code for this rule:

ts

```
const foo = new Set<number>();
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#configuration)

This rule accepts a configuration object with the following properties:

### allowShortCircuit [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#allowshortcircuit)

type: `boolean`

default: `false`

When set to `true`, allows short circuit evaluations in expressions.

### allowTaggedTemplates [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#allowtaggedtemplates)

type: `boolean`

default: `false`

When set to `true`, allows tagged template literals in expressions.

### allowTernary [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#allowternary)

type: `boolean`

default: `false`

When set to `true`, allows ternary operators in expressions.

### enforceForJSX [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#enforceforjsx)

type: `boolean`

default: `false`

When set to `true`, enforces the rule for unused JSX expressions also.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unused-expressions": "error"
  }
}
```

bash

```
oxlint --deny no-unused-expressions
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unused_expressions.rs)
