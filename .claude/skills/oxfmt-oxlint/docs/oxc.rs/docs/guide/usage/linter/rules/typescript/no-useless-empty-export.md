---
title: "typescript/no-useless-empty-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-useless-empty-export.md for this page in Markdown format

# typescript/no-useless-empty-export Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#typescript-no-useless-empty-export)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#what-it-does)

Disallow empty exports that don't change anything in a module file.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#why-is-this-bad)

An empty `export {}` statement is sometimes useful in TypeScript code to turn a file that would otherwise be a script file into a module file. Per the [TypeScript Handbook Modules page](https://www.typescriptlang.org/docs/handbook/modules/introduction.html):

In TypeScript, just as in ECMAScript 2015, any file containing a top-level import or export is considered a module. Conversely, a file without any top-level import or export declarations is treated as a script whose contents are available in the global scope (and therefore to modules as well).

However, an `export {}` statement does nothing if there are any other top-level import or export statements in a file.

This rule reports an `export {}` that doesn't do anything in a file already using ES modules.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#examples)

Examples of **incorrect** code for this rule:

ts

```
export const value = "Hello, world!";
export {};
```

Examples of **correct** code for this rule:

ts

```
export const value = "Hello, world!";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-useless-empty-export": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-useless-empty-export
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_useless_empty_export.rs)
