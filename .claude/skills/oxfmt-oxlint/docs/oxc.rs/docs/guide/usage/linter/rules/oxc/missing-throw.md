---
title: "oxc/missing-throw | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/missing-throw.md for this page in Markdown format

# oxc/missing-throw Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#oxc-missing-throw)

✅ This rule is turned on by default.

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#what-it-does)

Checks whether the `throw` keyword is missing in front of a `new` expression.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#why-is-this-bad)

The `throw` keyword is required in front of a `new` expression to throw an error. Omitting it is usually a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo() {
  throw Error();
}
const foo = () => {
  new Error();
};
```

Examples of **correct** code for this rule:

javascript

```
function foo() {
  throw new Error();
}
const foo = () => {
  throw new Error();
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/missing-throw": "error"
  }
}
```

bash

```
oxlint --deny oxc/missing-throw
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/missing_throw.rs)
