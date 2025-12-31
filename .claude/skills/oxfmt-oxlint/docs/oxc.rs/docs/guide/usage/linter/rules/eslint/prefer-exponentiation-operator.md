---
title: "eslint/prefer-exponentiation-operator | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.md for this page in Markdown format

# eslint/prefer-exponentiation-operator Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#eslint-prefer-exponentiation-operator)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#what-it-does)

Disallow the use of Math.pow in favor of the \*\* operator

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#why-is-this-bad)

Introduced in ES2016, the infix exponentiation operator \*\* is an alternative for the standard Math.pow function. Infix notation is considered to be more readable and thus more preferable than the function notation.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Math.pow(a, b);
```

Examples of **correct** code for this rule:

javascript

```
a ** b;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-exponentiation-operator": "error"
  }
}
```

bash

```
oxlint --deny prefer-exponentiation-operator
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_exponentiation_operator.rs)
