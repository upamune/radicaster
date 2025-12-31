---
title: "typescript/prefer-as-const | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-as-const.md for this page in Markdown format

# typescript/prefer-as-const Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#typescript-prefer-as-const)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#what-it-does)

Enforce the use of `as const` over literal type.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#why-is-this-bad)

There are two common ways to tell TypeScript that a literal value should be interpreted as its literal type (e.g. `2`) rather than general primitive type (e.g. `number`);

`as const`: telling TypeScript to infer the literal type automatically `as` with the literal type: explicitly telling the literal type to TypeScript

`as const` is generally preferred, as it doesn't require re-typing the literal value. This rule reports when an `as` with an explicit literal type can be replaced with an `as const`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#examples)

Examples of **incorrect** code for this rule:

ts

```
let bar: 2 = 2;
let foo = { bar: "baz" as "baz" };
```

Examples of **correct** code for this rule:

ts

```
let foo = "bar";
let foo = "bar" as const;
let foo: "bar" = "bar" as const;
let bar = "bar" as string;
let foo = { bar: "baz" };
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-as-const": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-as-const
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_as_const.rs)
