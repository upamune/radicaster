---
title: "eslint/prefer-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-spread.md for this page in Markdown format

# eslint/prefer-spread Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#eslint-prefer-spread)

This rule is combined 2 rules from `eslint:prefer-spread` and `unicorn:prefer-spread`.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#what-it-does)

Require spread operators instead of .apply()

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#why-is-this-bad)

Before ES2015, one must use Function.prototype.apply() to call variadic functions.

javascript

```
var args = [1, 2, 3, 4];
Math.max.apply(Math, args);
```

In ES2015, one can use spread syntax to call variadic functions.

javascript

```
var args = [1, 2, 3, 4];
Math.max(...args);
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.apply(undefined, args);
foo.apply(null, args);
obj.foo.apply(obj, args);
```

Examples of **correct** code for this rule:

javascript

```
// Using spread syntax
foo(...args);
obj.foo(...args);

// The `this` binding is different.
foo.apply(obj, args);
obj.foo.apply(null, args);
obj.foo.apply(otherObj, args);

// The argument list is not variadic.
// Those are warned by the `no-useless-call` rule.
foo.apply(undefined, [1, 2, 3]);
foo.apply(null, [1, 2, 3]);
obj.foo.apply(obj, [1, 2, 3]);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-spread": "error"
  }
}
```

bash

```
oxlint --deny prefer-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_spread.rs)
