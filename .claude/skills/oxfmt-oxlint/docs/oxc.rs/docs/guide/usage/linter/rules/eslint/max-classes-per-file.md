---
title: "eslint/max-classes-per-file | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-classes-per-file.md for this page in Markdown format

# eslint/max-classes-per-file Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#eslint-max-classes-per-file)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#what-it-does)

Enforce a maximum number of classes per file

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#why-is-this-bad)

Files containing multiple classes can often result in a less navigable and poorly structured codebase. Best practice is to keep each file limited to a single responsibility.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class Foo {}
class Bar {}
```

Examples of **correct** code for this rule:

js

```
function foo() {
  var bar = 1;
  let baz = 2;
  const qux = 3;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#ignoreexpressions)

type: `boolean`

default: `false`

Whether to ignore class expressions when counting classes.

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#max)

type: `integer`

default: `1`

The maximum number of classes allowed per file.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-classes-per-file": "error"
  }
}
```

bash

```
oxlint --deny max-classes-per-file
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_classes_per_file.rs)
