---
title: "eslint/prefer-object-has-own | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-object-has-own.md for this page in Markdown format

# eslint/prefer-object-has-own Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#eslint-prefer-object-has-own)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#what-it-does)

Disallow use of `Object.prototype.hasOwnProperty.call()` and prefer use of `Object.hasOwn()`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#why-is-this-bad)

It is very common to write code like:

javascript

```
if (Object.prototype.hasOwnProperty.call(object, "foo")) {
  console.log("has property foo");
}
```

This is a common practice because methods on Object.prototype can sometimes be unavailable or redefined (see the no-prototype-builtins rule). Introduced in ES2022, Object.hasOwn() is a shorter alternative to Object.prototype.hasOwnProperty.call():

javascript

```
if (Object.hasOwn(object, "foo")) {
  console.log("has property foo");
}
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#examples)

Examples of **incorrect** code for this rule:

js

```
Object.prototype.hasOwnProperty.call(obj, "a");
Object.hasOwnProperty.call(obj, "a");
({}).hasOwnProperty.call(obj, "a");
const hasProperty = Object.prototype.hasOwnProperty.call(object, property);
```

Examples of **correct** code for this rule:

js

```
Object.hasOwn(obj, "a");
const hasProperty = Object.hasOwn(object, property);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-object-has-own": "error"
  }
}
```

bash

```
oxlint --deny prefer-object-has-own
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_object_has_own.rs)
