---
title: "eslint/no-constant-condition | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-constant-condition.md for this page in Markdown format

# eslint/no-constant-condition Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#eslint-no-constant-condition)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#what-it-does)

Disallow constant expressions in conditions

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#why-is-this-bad)

A constant expression (for example, a literal) as a test condition might be a typo or development trigger for a specific behavior.

This rule disallows constant expressions in the test condition of:

* `if`, `for`, `while`, or `do...while` statement
* `?`: ternary expression

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#examples)

Examples of **incorrect** code for this rule:

js

```
if (false) {
  doSomethingUnfinished();
}

if (new Boolean(x)) {
  doSomethingAlways();
}
if ((x ||= true)) {
  doSomethingAlways();
}

do {
  doSomethingForever();
} while ((x = -1));
```

Examples of **correct** code for this rule:

js

```
if (x === 0) {
  doSomething();
}

while (typeof x === "undefined") {
  doSomething();
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#configuration)

This rule accepts a configuration object with the following properties:

### checkLoops [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#checkloops)

type: `"all" | "allExceptWhileTrue" | "none"`

default: `"allExceptWhileTrue"`

Configuration option to specify whether to check for constant conditions in loops.

* `"all"` or `true` disallows constant expressions in loops
* `"allExceptWhileTrue"` disallows constant expressions in loops except while loops with expression `true`
* `"none"` or `false` allows constant expressions in loops

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-constant-condition": "error"
  }
}
```

bash

```
oxlint --deny no-constant-condition
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_constant_condition.rs)
