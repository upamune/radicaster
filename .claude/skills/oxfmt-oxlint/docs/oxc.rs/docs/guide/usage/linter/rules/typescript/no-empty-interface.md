---
title: "typescript/no-empty-interface | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-empty-interface.md for this page in Markdown format

# typescript/no-empty-interface Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#typescript-no-empty-interface)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#what-it-does)

Disallow the declaration of empty interfaces.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#why-is-this-bad)

An empty interface in TypeScript does very little: any non-nullable value is assignable to {}. Using an empty interface is often a sign of programmer error, such as misunderstanding the concept of {} or forgetting to fill in fields. This rule aims to ensure that only meaningful interfaces are declared in the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#examples)

Examples of **incorrect** code for this rule:

ts

```
interface Foo {}
interface Bar extends Foo {}
```

Examples of **correct** code for this rule:

ts

```
interface Foo {
  member: string;
}
interface Bar extends Foo {
  member: string;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#configuration)

This rule accepts a configuration object with the following properties:

### allowSingleExtends [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#allowsingleextends)

type: `boolean`

default: `false`

When set to `true`, allows empty interfaces that extend a single interface.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-empty-interface": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-empty-interface
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_empty_interface.rs)
