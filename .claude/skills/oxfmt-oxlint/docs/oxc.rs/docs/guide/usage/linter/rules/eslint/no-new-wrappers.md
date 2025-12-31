---
title: "eslint/no-new-wrappers | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-new-wrappers.md for this page in Markdown format

# eslint/no-new-wrappers Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#eslint-no-new-wrappers)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#what-it-does)

Disallow `new` operators with the `String`, `Number`, and `Boolean` objects

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#why-is-this-bad)

The first problem is that primitive wrapper objects are, in fact, objects. That means typeof will return `"object"` instead of `"string"`, `"number"`, or `"boolean"`. The second problem comes with boolean objects. Every object is truthy, that means an instance of `Boolean` always resolves to `true` even when its actual value is `false`.

<https://eslint.org/docs/latest/rules/no-new-wrappers>

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#examples)

Examples of **incorrect** code for this rule:

js

```
var stringObject = new String("Hello world");
var numberObject = new Number(33);
var booleanObject = new Boolean(false);
var symbolObject = new Symbol("foo"); // symbol is not a constructor
```

Examples of **correct** code for this rule:

js

```
var stringObject = "Hello world";
var stringObject2 = String(value);
var numberObject = Number(value);
var booleanObject = Boolean(value);
var symbolObject = Symbol("foo");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-new-wrappers": "error"
  }
}
```

bash

```
oxlint --deny no-new-wrappers
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_new_wrappers.rs)
