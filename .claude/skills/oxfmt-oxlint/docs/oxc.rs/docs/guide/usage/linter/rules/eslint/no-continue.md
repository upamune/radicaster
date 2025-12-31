---
title: "eslint/no-continue | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-continue.md for this page in Markdown format

# eslint/no-continue Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#eslint-no-continue)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#what-it-does)

Disallow `continue` statements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#why-is-this-bad)

The continue statement terminates execution of the statements in the current iteration of the current or labeled loop, and continues execution of the loop with the next iteration. When used incorrectly it makes code less testable, less readable and less maintainable. Structured control flow statements such as if should be used instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var sum = 0,
  i;

for (i = 0; i < 10; i++) {
  if (i >= 5) {
    continue;
  }

  sum += i;
}
```

Examples of **correct** code for this rule:

javascript

```
var sum = 0,
  i;
for (i = 0; i < 10; i++) {
  if (i < 5) {
    sum += i;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-continue": "error"
  }
}
```

bash

```
oxlint --deny no-continue
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_continue.rs)
