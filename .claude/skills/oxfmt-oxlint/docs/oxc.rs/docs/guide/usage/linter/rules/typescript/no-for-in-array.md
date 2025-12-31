---
title: "typescript/no-for-in-array | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-for-in-array.md for this page in Markdown format

# typescript/no-for-in-array Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#typescript-no-for-in-array)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#what-it-does)

This rule disallows iterating over an array with a for-in loop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#why-is-this-bad)

A for-in loop iterates over the enumerable properties of an object, which includes the array indices, but also includes any enumerable properties added to the array prototype or the array instance. This is almost never what you want when iterating over an array.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const arr = [1, 2, 3];

for (const i in arr) {
  console.log(arr[i]);
}

for (const i in arr) {
  console.log(i, arr[i]);
}
```

Examples of **correct** code for this rule:

ts

```
const arr = [1, 2, 3];

// Use for-of to iterate over array values
for (const value of arr) {
  console.log(value);
}

// Use regular for loop with index
for (let i = 0; i < arr.length; i++) {
  console.log(i, arr[i]);
}

// Use forEach
arr.forEach((value, index) => {
  console.log(index, value);
});

// for-in is fine for objects
const obj = { a: 1, b: 2 };
for (const key in obj) {
  console.log(key, obj[key]);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-for-in-array": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-for-in-array
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_for_in_array.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_for_in_array/no_for_in_array.go)
