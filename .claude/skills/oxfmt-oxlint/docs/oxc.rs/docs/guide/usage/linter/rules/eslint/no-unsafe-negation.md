---
title: "eslint/no-unsafe-negation | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unsafe-negation.md for this page in Markdown format

# eslint/no-unsafe-negation Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#eslint-no-unsafe-negation)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#what-it-does)

Disallows negating the left operand of relational operators to prevent logical errors caused by misunderstanding operator precedence or accidental use of negation.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#why-is-this-bad)

Negating the left operand of relational operators can result in unexpected behavior due to operator precedence, leading to logical errors. For instance, `!a in b` may be interpreted as `(!a) in b` instead of `!(a in b)`, which is not the intended logic.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if ((!key) in object) {
}

if ((!obj) instanceof Ctor) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (!(key in object)) {
}

if (!(obj instanceof Ctor)) {
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForOrderingRelations [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#enforcefororderingrelations)

type: `boolean`

default: `false`

The `enforceForOrderingRelations` option determines whether negation is allowed on the left-hand side of ordering relational operators (<, >, <=, >=).

The purpose is to avoid expressions such as `!a < b` (which is equivalent to `(a ? 0 : 1) < b`) when what is really intended is `!(a < b)`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unsafe-negation": "error"
  }
}
```

bash

```
oxlint --deny no-unsafe-negation
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unsafe_negation.rs)
