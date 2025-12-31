---
title: "unicorn/no-null | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-null.md for this page in Markdown format

# unicorn/no-null Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#unicorn-no-null)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#what-it-does)

Disallow the use of the `null` literal, to encourage using `undefined` instead.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#why-is-this-bad)

There are some reasons for using `undefined` instead of `null`.

* From experience, most developers use `null` and `undefined` inconsistently and interchangeably, and few know when to use which.
* Supporting both `null` and `undefined` complicates input validation.
* Using `null` makes TypeScript types more verbose: `type A = {foo?: string | null}` vs `type A = {foo?: string}`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let foo = null;
```

Examples of **correct** code for this rule:

javascript

```
let foo;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#configuration)

This rule accepts a configuration object with the following properties:

### checkStrictEquality [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#checkstrictequality)

type: `boolean`

default: `false`

When set to `true`, the rule will also check strict equality/inequality comparisons (`===` and `!==`) against `null`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-null": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-null
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_null.rs)
