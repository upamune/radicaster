---
title: "eslint/no-object-constructor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-object-constructor.md for this page in Markdown format

# eslint/no-object-constructor Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#eslint-no-object-constructor)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#what-it-does)

Disallow calls to the Object constructor without an argument

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#why-is-this-bad)

Use of the Object constructor to construct a new empty object is generally discouraged in favor of object literal notation because of conciseness and because the Object global may be redefined. The exception is when the Object constructor is used to intentionally wrap a specified value which is passed as an argument.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#examples)

Examples of **incorrect** code for this rule:

js

```
Object();
new Object();
```

Examples of **correct** code for this rule:

js

```
Object("foo");
const obj = { a: 1, b: 2 };
const isObject = (value) => value === Object(value);
const createObject = (Object) => new Object();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-object-constructor": "error"
  }
}
```

bash

```
oxlint --deny no-object-constructor
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_object_constructor.rs)
