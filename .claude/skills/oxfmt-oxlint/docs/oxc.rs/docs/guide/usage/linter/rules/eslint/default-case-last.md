---
title: "eslint/default-case-last | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/default-case-last.md for this page in Markdown format

# eslint/default-case-last Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#eslint-default-case-last)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#what-it-does)

Requires the `default` clause in `switch` statements to be the last one.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#why-is-this-bad)

By convention and for readability, the `default` clause should be the last one in a `switch`. While it is legal to place it before or between `case` clauses, doing so is confusing and may lead to unexpected "fall-through" behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#examples)

Examples of **incorrect** code for this rule:

js

```
/* default-case-last: "error" */

switch (foo) {
  default:
    bar();
    break;
  case "a":
    baz();
    break;
}

switch (foo) {
  case 1:
    bar();
    break;
  default:
    baz();
    break;
  case 2:
    qux();
    break;
}
```

Examples of **correct** code for this rule:

js

```
/* default-case-last: "error" */

switch (foo) {
  case 1:
    bar();
    break;
  case 2:
    qux();
    break;
  default:
    baz();
    break;
}

switch (foo) {
  case "x":
    bar();
    break;
  case "y":
  default:
    baz();
    break;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "default-case-last": "error"
  }
}
```

bash

```
oxlint --deny default-case-last
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/default_case_last.rs)
