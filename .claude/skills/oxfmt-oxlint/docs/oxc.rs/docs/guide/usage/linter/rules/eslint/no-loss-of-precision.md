---
title: "eslint/no-loss-of-precision | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-loss-of-precision.md for this page in Markdown format

# eslint/no-loss-of-precision Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#eslint-no-loss-of-precision)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#what-it-does)

Disallow precision loss of number literal.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#why-is-this-bad)

It can lead to unexpected results in certain situations. For example, when performing mathematical operations.

In JavaScript, Numbers are stored as double-precision floating-point numbers according to the IEEE 754 standard. Because of this, numbers can only retain accuracy up to a certain amount of digits. If the programmer enters additional digits, those digits will be lost in the conversion to the Number type and will result in unexpected/incorrect behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var x = 2e999;
```

javascript

```
var x = 9007199254740993;
```

javascript

```
var x = 5123000000000000000000000000001;
```

javascript

```
var x = 1230000000000000000000000.0;
```

javascript

```
var x = 0x200000_0000000_1;
```

Examples of **correct** code for this rule:

javascript

```
var x = 12345;
```

javascript

```
var x = 123.456;
```

javascript

```
var x = 123.0;
```

javascript

```
var x = 123e34;
```

javascript

```
var x = 0x1fff_ffff_fff_fff;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-loss-of-precision": "error"
  }
}
```

bash

```
oxlint --deny no-loss-of-precision
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_loss_of_precision.rs)
