---
title: "eslint/no-negated-condition | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-negated-condition.md for this page in Markdown format

# eslint/no-negated-condition Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#eslint-no-negated-condition)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#what-it-does)

Disallow negated conditions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#why-is-this-bad)

Negated conditions are more difficult to understand. Code can be made more readable by inverting the condition.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (!a) {
  doSomethingC();
} else {
  doSomethingB();
}

!a ? doSomethingC() : doSomethingB();
```

Examples of **correct** code for this rule:

javascript

```
if (a) {
  doSomethingB();
} else {
  doSomethingC();
}

a ? doSomethingB() : doSomethingC();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-negated-condition": "error"
  }
}
```

bash

```
oxlint --deny no-negated-condition
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_negated_condition.rs)
