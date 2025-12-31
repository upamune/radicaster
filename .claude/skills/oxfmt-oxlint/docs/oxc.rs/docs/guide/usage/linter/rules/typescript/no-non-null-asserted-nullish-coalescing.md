---
title: "typescript/no-non-null-asserted-nullish-coalescing | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.md for this page in Markdown format

# typescript/no-non-null-asserted-nullish-coalescing Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#typescript-no-non-null-asserted-nullish-coalescing)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#what-it-does)

Disallow non-null assertions in the left operand of a nullish coalescing operator.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#why-is-this-bad)

The ?? nullish coalescing runtime operator allows providing a default value when dealing with null or undefined. Using a ! non-null assertion type operator in the left operand of a nullish coalescing operator is redundant, and likely a sign of programmer error or confusion over the two operators.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#examples)

Examples of **incorrect** code for this rule:

ts

```
foo! ?? bar;
foo.bazz! ?? bar;
foo!.bazz! ?? bar;
foo()! ?? bar;

let x!: string;
x! ?? "";

let x: string;
x = foo();
x! ?? "";
```

Examples of **correct** code for this rule:

ts

```
foo ?? bar;
foo ?? bar!;
foo!.bazz ?? bar;
foo!.bazz ?? bar!;
foo() ?? bar;
```

ts

```
// This is considered correct code because there's no way for the user to satisfy it.
let x: string;
x! ?? "";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-non-null-asserted-nullish-coalescing": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-non-null-asserted-nullish-coalescing
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_non_null_asserted_nullish_coalescing.rs)
