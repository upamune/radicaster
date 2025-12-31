---
title: "eslint/no-const-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-const-assign.md for this page in Markdown format

# eslint/no-const-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#eslint-no-const-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#what-it-does)

Disallow reassigning `const` variables.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#why-is-this-bad)

We cannot modify variables that are declared using the `const` keyword, as it will raise a runtime error.

Note that this rule is not necessary for TypeScript code, as TypeScript will already catch this as an error.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#examples)

Examples of **incorrect** code for this rule:

js

```
const a = 0;
a = 1;

const b = 0;
b += 1;
```

Examples of **correct** code for this rule:

js

```
const a = 0;
console.log(a);

var b = 0;
b += 1;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-const-assign": "error"
  }
}
```

bash

```
oxlint --deny no-const-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_const_assign.rs)
