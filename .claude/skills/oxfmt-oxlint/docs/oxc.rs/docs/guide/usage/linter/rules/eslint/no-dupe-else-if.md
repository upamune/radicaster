---
title: "eslint/no-dupe-else-if | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-dupe-else-if.md for this page in Markdown format

# eslint/no-dupe-else-if Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#eslint-no-dupe-else-if)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#what-it-does)

Disallow duplicate conditions in if-else-if chains

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#why-is-this-bad)

if-else-if chains are commonly used when there is a need to execute only one branch (or at most one branch) out of several possible branches, based on certain conditions. Two identical test conditions in the same chain are almost always a mistake in the code. Unless there are side effects in the expressions, a duplicate will evaluate to the same true or false value as the identical expression earlier in the chain, meaning that its branch can never execute.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (a) {
  foo();
} else if (b) {
  bar();
} else if (b) {
  baz();
}
```

javascript

```
if (a || b) {
  foo();
} else if (a) {
  bar();
}
```

javascript

```
if (n === 1) {
  foo();
} else if (n === 2) {
  bar();
} else if (n === 3) {
  baz();
} else if (n === 2) {
  quux();
} else if (n === 5) {
  quuux();
}
```

Examples of **correct** code for this rule:

javascript

```
if (a) {
  foo();
} else if (b) {
  bar();
} else if (c) {
  baz();
}
```

javascript

```
if (a || b) {
  foo();
} else if (c) {
  bar();
}
```

javascript

```
if (n === 1) {
  foo();
} else if (n === 2) {
  bar();
} else if (n === 3) {
  baz();
} else if (n === 4) {
  quux();
} else if (n === 5) {
  quuux();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-dupe-else-if": "error"
  }
}
```

bash

```
oxlint --deny no-dupe-else-if
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_dupe_else_if.rs)
