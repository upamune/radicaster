---
title: "oxc/const-comparisons | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/const-comparisons.md for this page in Markdown format

# oxc/const-comparisons Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#oxc-const-comparisons)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#what-it-does)

Checks for redundant or logically impossible comparisons. This includes:

* Ineffective double comparisons against constants.
* Impossible comparisons involving constants.
* Redundant comparisons where both operands are the same (e.g., a < a).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#why-is-this-bad)

Such comparisons can lead to confusing or incorrect logic in the program. In many cases:

* Only one of the comparisons has any effect on the result, suggesting that the programmer might have made a mistake, such as flipping one of the comparison operators or using the wrong variable.
* Comparisons like a < a or a >= a are always false or true respectively, making the logic redundant and potentially misleading.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
status_code <= 400 && status_code > 500;
status_code < 200 && status_code <= 299;
status_code > 500 && status_code >= 500;
a < a; // Always false
a >= a; // Always true
```

Examples of **correct** code for this rule:

javascript

```
status_code >= 400 && status_code < 500;
500 <= status_code && 600 > status_code;
500 <= status_code && status_code <= 600;
a < b;
a <= b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/const-comparisons": "error"
  }
}
```

bash

```
oxlint --deny oxc/const-comparisons
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/const_comparisons.rs)
