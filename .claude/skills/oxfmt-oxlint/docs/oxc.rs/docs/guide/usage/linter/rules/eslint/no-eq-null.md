---
title: "eslint/no-eq-null | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-eq-null.md for this page in Markdown format

# eslint/no-eq-null Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#eslint-no-eq-null)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#what-it-does)

Disallow `null` comparisons without type-checking operators.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#why-is-this-bad)

Comparing to `null` without a type-checking operator (`==` or `!=`), can have unintended results as the comparison will evaluate to `true` when comparing to not just a `null`, but also an `undefined` value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#examples)

Examples of **incorrect** code for this rule:

js

```
if (foo == null) {
  bar();
}
if (baz != null) {
  bar();
}
```

Examples of **correct** code for this rule:

js

```
if (foo === null) {
  bar();
}

if (baz !== null) {
  bar();
}

if (bang === undefined) {
  bar();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-eq-null": "error"
  }
}
```

bash

```
oxlint --deny no-eq-null
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_eq_null.rs)
