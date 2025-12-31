---
title: "eslint/no-new-native-nonconstructor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.md for this page in Markdown format

# eslint/no-new-native-nonconstructor Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#eslint-no-new-native-nonconstructor)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#what-it-does)

Disallow `new` operators with global non-constructor functions (`Symbol`, `BigInt`).

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#why-is-this-bad)

Both `new Symbol` and `new BigInt` throw a type error because they are functions and not classes. It is easy to make this mistake by assuming the uppercase letters indicate classes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#examples)

Examples of **incorrect** code for this rule:

js

```
// throws a TypeError
let foo = new Symbol("foo");

// throws a TypeError
let result = new BigInt(9007199254740991);
```

Examples of **correct** code for this rule:

js

```
let foo = Symbol("foo");

let result = BigInt(9007199254740991);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-new-native-nonconstructor": "error"
  }
}
```

bash

```
oxlint --deny no-new-native-nonconstructor
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_new_native_nonconstructor.rs)
