---
title: "eslint/sort-keys | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/sort-keys.md for this page in Markdown format

# eslint/sort-keys Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#eslint-sort-keys)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#what-it-does)

When declaring multiple properties, sorting property names alphabetically makes it easier to find and/or diff necessary properties at a later time.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#why-is-this-bad)

Unsorted property keys can make the code harder to read and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#examples)

Examples of **incorrect** code for this rule:

js

```
let myObj = {
  c: 1,
  a: 2,
};
```

Examples of **correct** code for this rule:

js

```
let myObj = {
  a: 2,
  c: 1,
};
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#configuration)

### The 1st option [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#the-1st-option)

type: `"desc" | "asc"`

Sorting order for keys. Accepts "asc" for ascending or "desc" for descending.

### The 2nd option [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#the-2nd-option)

This option is an object with the following properties:

#### allowLineSeparatedGroups [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#allowlineseparatedgroups)

type: `boolean`

default: `false`

When true, groups of properties separated by a blank line are sorted independently.

#### caseSensitive [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#casesensitive)

type: `boolean`

default: `true`

Whether the sort comparison is case-sensitive (A < a when true).

#### minKeys [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#minkeys)

type: `integer`

default: `2`

Minimum number of properties required in an object before sorting is enforced.

#### natural [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#natural)

type: `boolean`

default: `false`

Use natural sort order so that, for example, "a2" comes before "a10".

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "sort-keys": "error"
  }
}
```

bash

```
oxlint --deny sort-keys
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/sort_keys.rs)
