---
title: "oxc/bad-bitwise-operator | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.md for this page in Markdown format

# oxc/bad-bitwise-operator Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#oxc-bad-bitwise-operator)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#what-it-does)

This rule applies when bitwise operators are used where logical operators are expected.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#why-is-this-bad)

Bitwise operators have different results from logical operators and a `TypeError` exception may be thrown because short-circuit evaluation is not applied. (In short-circuit evaluation, right operand evaluation is skipped according to left operand value, e.g. `x` is `false` in `x && y`.)

It is obvious that logical operators are expected in the following code patterns:

javascript

```
e && e.x;
e || {};
e || "";
```

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (obj & obj.prop) {
  console.log(obj.prop);
}
options = options | {};
input |= "";
```

Examples of **correct** code for this rule:

javascript

```
if (obj && obj.prop) {
  console.log(obj.prop);
}
options = options || {};
input ||= "";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-bitwise-operator": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-bitwise-operator
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_bitwise_operator.rs)
