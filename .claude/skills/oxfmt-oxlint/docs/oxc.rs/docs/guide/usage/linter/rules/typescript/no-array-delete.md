---
title: "typescript/no-array-delete | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-array-delete.md for this page in Markdown format

# typescript/no-array-delete Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#typescript-no-array-delete)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#what-it-does)

This rule disallows using the delete operator on array values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#why-is-this-bad)

When using the delete operator on an array, the element is not actually removed, but instead the array slot is turned into undefined. This is usually not the intended behavior. Instead, you should use methods like Array.prototype.splice() to properly remove elements from an array.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const arr: number[];
delete arr[0];
```

Examples of **correct** code for this rule:

ts

```
declare const arr: number[];
arr.splice(0, 1);

// or with a filter
const filteredArr = arr.filter((_, index) => index !== 0);

// delete on object is allowed
declare const obj: { a?: number };
delete obj.a;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-array-delete": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-array-delete
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_array_delete.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_array_delete/no_array_delete.go)
