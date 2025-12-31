---
title: "oxc/bad-min-max-func | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-min-max-func.md for this page in Markdown format

# oxc/bad-min-max-func Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#oxc-bad-min-max-func)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#what-it-does)

Checks whether the clamp function `Math.min(Math.max(x, y), z)` always evaluate to a constant result because the arguments are in the wrong order.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#why-is-this-bad)

The `Math.min(Math.max(x, y), z)` function is used to clamp a value between two other values. If the arguments are in the wrong order, the function will always evaluate to a constant result.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Math.min(Math.max(100, x), 0);
Math.max(1000, Math.min(0, z));
```

Examples of **correct** code for this rule:

javascript

```
Math.max(0, Math.min(100, x));
Math.min(1000, Math.max(0, z));
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-min-max-func": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-min-max-func
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_min_max_func.rs)
