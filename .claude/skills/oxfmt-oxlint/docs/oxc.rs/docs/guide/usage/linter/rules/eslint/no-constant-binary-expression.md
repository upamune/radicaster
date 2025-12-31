---
title: "eslint/no-constant-binary-expression | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.md for this page in Markdown format

# eslint/no-constant-binary-expression Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#eslint-no-constant-binary-expression)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#what-it-does)

Disallow expressions where the operation doesn't affect the value

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#why-is-this-bad)

Comparisons which will always evaluate to true or false and logical expressions (`||`, `&&`, `??`) which either always short-circuit or never short-circuit are both likely indications of programmer error.

These errors are especially common in complex expressions where operator precedence is easy to misjudge.

Additionally, this rule detects comparisons to newly constructed objects/arrays/functions/etc. In JavaScript, where objects are compared by reference, a newly constructed object can never `===` any other value. This can be surprising for programmers coming from languages where objects are compared by value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// One might think this would evaluate as `a + (b ?? c)`:
const x = a + b ?? c;

// But it actually evaluates as `(a + b) ?? c`. Since `a + b` can never be null,
// the `?? c` has no effect.

// Programmers coming from a language where objects are compared by value might expect this to work:
const isEmpty = x === [];

// However, this will always result in `isEmpty` being `false`.
```

Examples of **correct** code for this rule:

javascript

```
const x = a + (b ?? c);

const isEmpty = x.length === 0;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-constant-binary-expression": "error"
  }
}
```

bash

```
oxlint --deny no-constant-binary-expression
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_constant_binary_expression.rs)
