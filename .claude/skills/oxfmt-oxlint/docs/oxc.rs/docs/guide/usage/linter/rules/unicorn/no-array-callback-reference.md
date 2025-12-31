---
title: "unicorn/no-array-callback-reference | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.md for this page in Markdown format

# unicorn/no-array-callback-reference Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#unicorn-no-array-callback-reference)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#what-it-does)

Prevents passing a function reference directly to iterator methods

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#why-is-this-bad)

Passing functions to iterator methods can cause issues when the function is changed without realizing that the iterator passes 2 more parameters to it (index and array). This can lead to unexpected behavior when the function signature changes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#examples)

Examples of **incorrect** code for this rule:

js

```
const foo = array.map(callback);
array.forEach(callback);
const result = array.filter(lib.method);
```

Examples of **correct** code for this rule:

js

```
const foo = array.map((element) => callback(element));
array.forEach((element) => {
  callback(element);
});
const result = array.filter((element) => lib.method(element));

// Built-in functions are allowed
const foo = array.map(String);
const bar = array.filter(Boolean);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-callback-reference": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-callback-reference
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_callback_reference.rs)
