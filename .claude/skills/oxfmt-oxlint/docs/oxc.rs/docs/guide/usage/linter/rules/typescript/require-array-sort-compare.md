---
title: "typescript/require-array-sort-compare | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/require-array-sort-compare.md for this page in Markdown format

# typescript/require-array-sort-compare Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#typescript-require-array-sort-compare)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#what-it-does)

This rule requires Array.sort() to be called with a comparison function.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#why-is-this-bad)

When Array.sort() is called without a comparison function, it converts elements to strings and sorts them lexicographically. This often leads to unexpected results, especially with numbers where `[1, 10, 2].sort()` returns `[1, 10, 2]` instead of `[1, 2, 10]`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const numbers = [3, 1, 4, 1, 5];
numbers.sort(); // Lexicographic sort, not numeric

const mixedArray = ["10", "2", "1"];
mixedArray.sort(); // Might be intended, but explicit compareFn is clearer

[3, 1, 4].sort(); // Will sort as strings: ['1', '3', '4']
```

Examples of **correct** code for this rule:

ts

```
const numbers = [3, 1, 4, 1, 5];

// Numeric sort
numbers.sort((a, b) => a - b);

// Reverse numeric sort
numbers.sort((a, b) => b - a);

// String sort (explicit)
const strings = ["banana", "apple", "cherry"];
strings.sort((a, b) => a.localeCompare(b));

// Custom object sorting
interface Person {
  name: string;
  age: number;
}

const people: Person[] = [
  { name: "Alice", age: 30 },
  { name: "Bob", age: 25 },
];

people.sort((a, b) => a.age - b.age);
people.sort((a, b) => a.name.localeCompare(b.name));
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreStringArrays [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#ignorestringarrays)

type: `boolean`

default: `true`

Whether to ignore arrays in which all elements are strings.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/require-array-sort-compare": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/require-array-sort-compare
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/require_array_sort_compare.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/require_array_sort_compare/require_array_sort_compare.go)
