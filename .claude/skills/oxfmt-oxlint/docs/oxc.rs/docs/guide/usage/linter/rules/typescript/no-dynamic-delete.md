---
title: "typescript/no-dynamic-delete | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-dynamic-delete.md for this page in Markdown format

# typescript/no-dynamic-delete Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#typescript-no-dynamic-delete)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#what-it-does)

Disallow using the delete operator on computed key expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#why-is-this-bad)

Deleting dynamically computed keys can be dangerous and in some cases not well optimized. Using the delete operator on keys that aren't runtime constants could be a sign that you're using the wrong data structures. Consider using a Map or Set if you’re using an object as a key-value collection.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const container: { [i: string]: 0 } = {};
delete container["aa" + "b"];
```

Examples of **correct** code for this rule:

ts

```
const container: { [i: string]: 0 } = {};
delete container.aab;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-dynamic-delete": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-dynamic-delete
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_dynamic_delete.rs)
