---
title: "unicorn/require-number-to-fixed-digits-argument | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.md for this page in Markdown format

# unicorn/require-number-to-fixed-digits-argument Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#unicorn-require-number-to-fixed-digits-argument)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#what-it-does)

Enforce using the digits argument with Number.toFixed()

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#why-is-this-bad)

It's better to make it clear what the value of the digits argument is when calling Number.toFixed(), instead of relying on the default value of 0.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
number.toFixed();
```

Examples of **correct** code for this rule:

javascript

```
number.toFixed(0);
number.toFixed(2);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/require-number-to-fixed-digits-argument": "error"
  }
}
```

bash

```
oxlint --deny unicorn/require-number-to-fixed-digits-argument
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/require_number_to_fixed_digits_argument.rs)
