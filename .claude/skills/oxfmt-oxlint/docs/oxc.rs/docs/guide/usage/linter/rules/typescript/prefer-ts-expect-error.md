---
title: "typescript/prefer-ts-expect-error | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.md for this page in Markdown format

# typescript/prefer-ts-expect-error Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#typescript-prefer-ts-expect-error)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#what-it-does)

Enforce using @ts-expect-error over @ts-ignore.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#why-is-this-bad)

TypeScript allows you to suppress all errors on a line by placing a comment starting with @ts-ignore or @ts-expect-error immediately before the erroring line. The two directives work the same, except @ts-expect-error causes a type error if placed before a line that's not erroring in the first place.

This means it's easy for @ts-ignores to be forgotten about, and remain in code even after the error they were suppressing is fixed. This is dangerous, as if a new error arises on that line it'll be suppressed by the forgotten about @ts-ignore, and so be missed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// @ts-ignore
const str: string = 1;

/**
 * Explaining comment
 *
 * @ts-ignore */
const multiLine: number = "value";
```

Examples of **incorrect** code for this rule:

ts

```
/**
 * Explaining comment
 *
 * @ts-expect-error */
const multiLine: number = "value";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-ts-expect-error": "error"
  }
}
```

bash

```
oxlint --deny typescript/prefer-ts-expect-error
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_ts_expect_error.rs)
