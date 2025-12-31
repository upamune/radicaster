---
title: "eslint/no-extend-native | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-extend-native.md for this page in Markdown format

# eslint/no-extend-native Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#eslint-no-extend-native)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#what-it-does)

Prevents extending native global objects such as `Object`, `String`, or `Array` with new properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#why-is-this-bad)

Extending native objects can cause unexpected behavior and conflicts with other code.

For example:

js

```
// Adding a new property, which might seem okay
Object.prototype.extra = 55;

// Defining a user object
const users = {
  1: "user1",
  2: "user2",
};

for (const id in users) {
  // This will print "extra" as well as "1" and "2":
  console.log(id);
}
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#examples)

Examples of **incorrect** code for this rule:

js

```
Object.prototype.p = 0;
Object.defineProperty(Array.prototype, "p", { value: 0 });
```

Examples of **correct** code for this rule:

js

```
x.prototype.p = 0;
Object.defineProperty(x.prototype, "p", { value: 0 });
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#configuration)

This rule accepts a configuration object with the following properties:

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#exceptions)

type: `string[]`

default: `[]`

A list of objects which are allowed to be exceptions to the rule.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-extend-native": "error"
  }
}
```

bash

```
oxlint --deny no-extend-native
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_extend_native.rs)
