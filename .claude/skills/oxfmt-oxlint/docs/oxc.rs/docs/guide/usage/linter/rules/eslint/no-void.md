---
title: "eslint/no-void | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-void.md for this page in Markdown format

# eslint/no-void Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#eslint-no-void)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#what-it-does)

Disallows the use of the `void` operator.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#why-is-this-bad)

The `void` operator is often used to get `undefined`, but this is unnecessary because `undefined` can be used directly instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#examples)

Examples of **incorrect** code for this rule:

ts

```
void 0;
var foo = void 0;
```

Examples of **correct** code for this rule:

ts

```
"var foo = bar()";
"foo.void()";
"foo.void = bar";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAsStatement [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#allowasstatement)

type: `boolean`

default: `false`

If set to `true`, using `void` as a standalone statement is allowed.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-void": "error"
  }
}
```

bash

```
oxlint --deny no-void
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_void.rs)
