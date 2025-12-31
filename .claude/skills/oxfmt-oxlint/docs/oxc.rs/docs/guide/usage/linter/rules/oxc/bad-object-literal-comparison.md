---
title: "oxc/bad-object-literal-comparison | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.md for this page in Markdown format

# oxc/bad-object-literal-comparison Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#oxc-bad-object-literal-comparison)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#what-it-does)

Checks for comparisons between object and array literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#why-is-this-bad)

Comparing a variable to an object or array literal will always return false as object and array literals are never equal to each other.

If you want to check if an object or array is empty, use `Object.entries()` or `Object.keys()` and their lengths.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (x === {}) {
}
if (arr !== []) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (typeof x === "object" && Object.keys(x).length === 0) {
}
if (Array.isArray(x) && x.length === 0) {
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-object-literal-comparison": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-object-literal-comparison
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_object_literal_comparison.rs)
