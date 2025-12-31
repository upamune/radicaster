---
title: "eslint/no-loop-func | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-loop-func.md for this page in Markdown format

# eslint/no-loop-func Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#eslint-no-loop-func)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#what-it-does)

Disallows function declarations and expressions inside loop statements when they reference variables declared in the outer scope that may change across iterations.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#why-is-this-bad)

Writing functions within loops tends to result in errors due to the way closures work in JavaScript. Functions capture variables by reference, not by value. When using `var`, which is function-scoped, all iterations share the same variable binding, leading to unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#examples)

Examples of **incorrect** code for this rule:

js

```
for (var i = 0; i < 10; i++) {
  funcs[i] = function () {
    return i;
  };
}
```

Examples of **correct** code for this rule:

js

```
for (let i = 0; i < 10; i++) {
  funcs[i] = function () {
    return i;
  };
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-loop-func": "error"
  }
}
```

bash

```
oxlint --deny no-loop-func
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_loop_func.rs)
