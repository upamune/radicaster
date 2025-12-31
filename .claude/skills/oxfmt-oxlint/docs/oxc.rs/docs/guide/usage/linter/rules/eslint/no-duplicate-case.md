---
title: "eslint/no-duplicate-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-duplicate-case.md for this page in Markdown format

# eslint/no-duplicate-case Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#eslint-no-duplicate-case)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#what-it-does)

Disallow duplicate case labels

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#why-is-this-bad)

If a switch statement has duplicate test expressions in case clauses, it is likely that a programmer copied a case clause but forgot to change the test expression.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#examples)

Examples of **incorrect** code for this rule:

js

```
var a = 1,
  one = 1;
switch (a) {
  case 1:
    break;
  case 2:
    break;
  case 1: // duplicate test expression
    break;
  default:
    break;
}

switch (a) {
  case one:
    break;
  case 2:
    break;
  case one: // duplicate test expression
    break;
  default:
    break;
}
```

Examples of **correct** code for this rule:

js

```
var a = 1,
  one = 1;
switch (a) {
  case 1:
    break;
  case 2:
    break;
  default:
    break;
}

switch (a) {
  case "1":
    break;
  case "2":
    break;
  default:
    break;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-duplicate-case": "error"
  }
}
```

bash

```
oxlint --deny no-duplicate-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_duplicate_case.rs)
