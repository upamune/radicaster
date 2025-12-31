---
title: "oxc/number-arg-out-of-range | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.md for this page in Markdown format

# oxc/number-arg-out-of-range Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#oxc-number-arg-out-of-range)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#what-it-does)

Checks whether the radix or precision arguments of number-related functions exceeds the limit.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#why-is-this-bad)

The radix argument of `Number.prototype.toString` should be between 2 and 36. The precision argument of `Number.prototype.toFixed` and `Number.prototype.toExponential` should be between 0 and 20. The precision argument of `Number.prototype.toPrecision` should be between 1 and 21.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x = 42;
var s_radix_64 = x.toString(64);
var s = x.toString(1);
```

Examples of **correct** code for this rule:

javascript

```
var x = 42;
var s_radix_16 = x.toString(16);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/number-arg-out-of-range": "error"
  }
}
```

bash

```
oxlint --deny oxc/number-arg-out-of-range
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/number_arg_out_of_range.rs)
