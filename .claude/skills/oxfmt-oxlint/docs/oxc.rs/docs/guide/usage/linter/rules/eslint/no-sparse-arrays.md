---
title: "eslint/no-sparse-arrays | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-sparse-arrays.md for this page in Markdown format

# eslint/no-sparse-arrays Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#eslint-no-sparse-arrays)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#what-it-does)

Disallow sparse arrays.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#why-is-this-bad)

Take the following example:

javascript

```
const items = [, ,];
```

While the items array in this example has a length of 2, there are actually no values in items[0] or items[1]. The fact that the array literal is valid with only commas inside, coupled with the length being set and actual item values not being set, make sparse arrays confusing for many developers.

The confusion around sparse arrays is enough that it’s recommended to avoid using them unless you are certain that they are useful in your code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var items = [, ,];
```

javascript

```
var colors = ["red", , "blue"];
```

Examples of **correct** code for this rule:

javascript

```
var items = [];
```

// trailing comma (after the last element) is not a problem

javascript

```
var colors = ["red", "blue"];
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-sparse-arrays": "error"
  }
}
```

bash

```
oxlint --deny no-sparse-arrays
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_sparse_arrays.rs)
