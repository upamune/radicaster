---
title: "typescript/no-extra-non-null-assertion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.md for this page in Markdown format

# typescript/no-extra-non-null-assertion Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#typescript-no-extra-non-null-assertion)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#what-it-does)

Disallow extra non-null assertions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#why-is-this-bad)

The `!` non-null assertion operator in TypeScript is used to assert that a value's type does not include null or undefined. Using the operator any more than once on a single value does nothing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const foo: { bar: number } | null = null;
const bar = foo!!!.bar;
```

ts

```
function foo(bar: number | undefined) {
  const bar: number = bar!!!;
}
```

ts

```
function foo(bar?: { n: number }) {
  return bar!?.n;
}
```

Examples of **correct** code for this rule:

ts

```
const foo: { bar: number } | null = null;
const bar = foo!.bar;
```

ts

```
function foo(bar: number | undefined) {
  const bar: number = bar!;
}
```

ts

```
function foo(bar?: { n: number }) {
  return bar?.n;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-extra-non-null-assertion": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-extra-non-null-assertion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_extra_non_null_assertion.rs)
