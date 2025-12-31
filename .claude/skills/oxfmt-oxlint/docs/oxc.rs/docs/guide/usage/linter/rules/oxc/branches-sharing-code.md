---
title: "oxc/branches-sharing-code | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/branches-sharing-code.md for this page in Markdown format

# oxc/branches-sharing-code Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#oxc-branches-sharing-code)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#what-it-does)

Checks if the `if` and `else` blocks contain shared code that can be moved out of the blocks.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#why-is-this-bad)

Duplicate code is less maintainable. Extracting common code from branches makes the code more DRY (Don't Repeat Yourself) and easier to maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (condition) {
  console.log("Hello");
  return 13;
} else {
  console.log("Hello");
  return 42;
}

if (condition) {
  doSomething();
  cleanup();
} else {
  doSomethingElse();
  cleanup();
}
```

Examples of **correct** code for this rule:

javascript

```
console.log("Hello");
if (condition) {
  return 13;
} else {
  return 42;
}

if (condition) {
  doSomething();
} else {
  doSomethingElse();
}
cleanup();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/branches-sharing-code": "error"
  }
}
```

bash

```
oxlint --deny oxc/branches-sharing-code
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/branches_sharing_code.rs)
