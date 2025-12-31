---
title: "oxc/bad-char-at-comparison | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.md for this page in Markdown format

# oxc/bad-char-at-comparison Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#oxc-bad-char-at-comparison)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#what-it-does)

This rule warns when the return value of the `charAt` method is used to compare a string of length greater than 1.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#why-is-this-bad)

The `charAt` method returns a string of length 1. If the return value is compared with a string of length greater than 1, the comparison will always be false.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
a.charAt(4) === "a2";
a.charAt(4) === "/n";
```

Examples of **correct** code for this rule:

javascript

```
a.charAt(4) === "a";
a.charAt(4) === "\n";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-char-at-comparison": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-char-at-comparison
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_char_at_comparison.rs)
