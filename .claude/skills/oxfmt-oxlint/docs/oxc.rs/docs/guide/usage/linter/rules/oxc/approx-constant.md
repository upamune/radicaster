---
title: "oxc/approx-constant | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/approx-constant.md for this page in Markdown format

# oxc/approx-constant Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#oxc-approx-constant)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#what-it-does)

Disallows the use of approximate constants, instead preferring the use of the constants in the `Math` object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#why-is-this-bad)

Approximate constants are not as accurate as the constants in the `Math` object. Using the `Math` constants improves code readability and accuracy. See <https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math> for more information.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let log10e = 0.434294;
```

Examples of **correct** code for this rule:

javascript

```
let log10e = Math.LOG10E;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/approx-constant": "error"
  }
}
```

bash

```
oxlint --deny oxc/approx-constant
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/approx_constant.rs)
