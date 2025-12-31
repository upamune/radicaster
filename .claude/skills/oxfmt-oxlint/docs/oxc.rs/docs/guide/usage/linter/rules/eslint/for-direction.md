---
title: "eslint/for-direction | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/for-direction.md for this page in Markdown format

# eslint/for-direction Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#eslint-for-direction)

✅ This rule is turned on by default.

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#what-it-does)

Disallow `for` loops where the update clause moves the counter in the wrong direction, preventing the loop from reaching its stop condition.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#why-is-this-bad)

A `for` loop with a stop condition that can never be reached will run infinitely. While infinite loops can be intentional, they are usually written as `while` loops. More often, an infinite `for` loop is a bug.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#examples)

Examples of **incorrect** code for this rule:

js

```
/* for-direction: "error" */

for (var i = 0; i < 10; i--) {}

for (var i = 10; i >= 0; i++) {}

for (var i = 0; i > 10; i++) {}

for (var i = 0; 10 > i; i--) {}

const n = -2;
for (let i = 0; i < 10; i += n) {}
```

Examples of **correct** code for this rule:

js

```
/* for-direction: "error" */

for (var i = 0; i < 10; i++) {}

for (var i = 0; 10 > i; i++) {
  // with counter "i" on the right
}

for (let i = 10; i >= 0; i += this.step) {
  // direction unknown
}

for (let i = MIN; i <= MAX; i -= 0) {
  // not increasing or decreasing
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "for-direction": "error"
  }
}
```

bash

```
oxlint --deny for-direction
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/for_direction.rs)
