---
title: "eslint/no-unsafe-optional-chaining | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.md for this page in Markdown format

# eslint/no-unsafe-optional-chaining Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#eslint-no-unsafe-optional-chaining)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#what-it-does)

Disallow use of optional chaining in contexts where the undefined value is not allowed

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#why-is-this-bad)

The optional chaining (`?.`) expression can short-circuit with a return value of undefined. Therefore, treating an evaluated optional chaining expression as a function, object, number, etc., can cause TypeError or unexpected results. For example:

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var obj = undefined;
1 in obj?.foo; // TypeError
with (obj?.foo); // TypeError
for (bar of obj?.foo); // TypeError
bar instanceof obj?.foo; // TypeError
const { bar } = obj?.foo; // TypeError
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#configuration)

This rule accepts a configuration object with the following properties:

### disallowArithmeticOperators [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#disallowarithmeticoperators)

type: `boolean`

default: `false`

Disallow arithmetic operations on optional chaining expressions. If this is true, this rule warns arithmetic operations on optional chaining expressions, which possibly result in NaN.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unsafe-optional-chaining": "error"
  }
}
```

bash

```
oxlint --deny no-unsafe-optional-chaining
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unsafe_optional_chaining.rs)
