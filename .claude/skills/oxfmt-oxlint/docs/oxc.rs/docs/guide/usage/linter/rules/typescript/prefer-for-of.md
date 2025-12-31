---
title: "typescript/prefer-for-of | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-for-of.md for this page in Markdown format

# typescript/prefer-for-of Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#typescript-prefer-for-of)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#what-it-does)

Enforces the use of for-of loop instead of a for loop with a simple iteration.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#why-is-this-bad)

Using a for loop with a simple iteration over an array can be replaced with a more concise and readable for-of loop. For-of loops are easier to read and less error-prone, as they eliminate the need for an index variable and manual array access.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
for (let i = 0; i < arr.length; i++) {
  console.log(arr[i]);
}
```

Examples of **correct** code for this rule:

typescript

```
for (const item of arr) {
  console.log(item);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-for-of": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-for-of
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_for_of.rs)
