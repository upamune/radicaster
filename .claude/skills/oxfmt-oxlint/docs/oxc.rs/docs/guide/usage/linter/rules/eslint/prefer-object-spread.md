---
title: "eslint/prefer-object-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-object-spread.md for this page in Markdown format

# eslint/prefer-object-spread Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#eslint-prefer-object-spread)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#what-it-does)

Disallow using `Object.assign` with an object literal as the first argument and prefer the use of object spread instead

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#why-is-this-bad)

When `Object.assign` is called using an object literal as the first argument, this rule requires using the object spread syntax instead. This rule also warns on cases where an `Object.assign` call is made using a single argument that is an object literal, in this case, the `Object.assign` call is not needed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#examples)

Examples of **incorrect** code for this rule:

js

```
Object.assign({}, foo);

Object.assign({}, { foo: "bar" });

Object.assign({ foo: "bar" }, baz);

Object.assign({}, baz, { foo: "bar" });

Object.assign({}, { ...baz });

// Object.assign with a single argument that is an object literal
Object.assign({});

Object.assign({ foo: bar });
```

Examples of **correct** code for this rule:

js

```
({ ...foo });

({ ...baz, foo: "bar" });

// Any Object.assign call without an object literal as the first argument
Object.assign(foo, { bar: baz });

Object.assign(foo, bar);

Object.assign(foo, { bar, baz });

Object.assign(foo, { ...baz });
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-object-spread": "error"
  }
}
```

bash

```
oxlint --deny prefer-object-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_object_spread.rs)
