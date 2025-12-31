---
title: "typescript/no-unsafe-declaration-merging | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.md for this page in Markdown format

# typescript/no-unsafe-declaration-merging Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#typescript-no-unsafe-declaration-merging)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#what-it-does)

Disallow unsafe declaration merging.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#why-is-this-bad)

Declaration merging between classes and interfaces is unsafe. The TypeScript compiler doesn't check whether properties are initialized, which can lead to TypeScript not detecting code that will cause runtime errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#examples)

Examples of **incorrect** code for this rule:

ts

```
interface Foo {}
class Foo {}
```

Examples of **correct** code for this rule:

ts

```
interface Foo {}
class Bar {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-declaration-merging": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-unsafe-declaration-merging
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_declaration_merging.rs)
