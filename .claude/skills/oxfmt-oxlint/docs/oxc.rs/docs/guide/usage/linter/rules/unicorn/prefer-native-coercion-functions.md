---
title: "unicorn/prefer-native-coercion-functions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.md for this page in Markdown format

# unicorn/prefer-native-coercion-functions Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#unicorn-prefer-native-coercion-functions)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#what-it-does)

Prefers built in functions, over custom ones with the same functionality.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#why-is-this-bad)

If a function is equivalent to [`String`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String), [`Number`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number), [`BigInt`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/BigInt), [`Boolean`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Boolean), or [`Symbol`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Symbol), you should use the built-in one directly. Wrapping the built-in in a function is moot.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = (v) => String(v);
foo(1);
const foo = (v) => Number(v);
array.some((v) => /* comment */ v);
```

Examples of **correct** code for this rule:

javascript

```
String(1);
Number(1);
array.some(Boolean);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-native-coercion-functions": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-native-coercion-functions
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_native_coercion_functions.rs)
