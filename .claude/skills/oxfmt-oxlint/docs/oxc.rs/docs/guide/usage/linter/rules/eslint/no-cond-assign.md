---
title: "eslint/no-cond-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-cond-assign.md for this page in Markdown format

# eslint/no-cond-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#eslint-no-cond-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#what-it-does)

Disallow assignment operators in conditional expressions

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#why-is-this-bad)

In conditional statements, it is very easy to mistype a comparison operator (such as `==`) as an assignment operator (such as `=`).

There are valid reasons to use assignment operators in conditional statements. However, it can be difficult to tell whether a specific assignment was intentional.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#examples)

Examples of **incorrect** code for this rule:

js

```
// Check the user's job title
if ((user.jobTitle = "manager")) {
  // user.jobTitle is now incorrect
}
```

Examples of **correct** code for this rule:

js

```
// Check the user's job title
if (user.jobTitle === "manager") {
  // correctly compared `jobTitle`
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#configuration)

This rule accepts one of the following string values:

### `"except-parens"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#except-parens)

Allow assignments in conditional expressions only if they are enclosed in parentheses.

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#always)

Disallow all assignments in conditional expressions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-cond-assign": "error"
  }
}
```

bash

```
oxlint --deny no-cond-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_cond_assign.rs)
