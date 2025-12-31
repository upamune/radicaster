---
title: "typescript/consistent-indexed-object-style | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.md for this page in Markdown format

# typescript/consistent-indexed-object-style Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#typescript-consistent-indexed-object-style)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#what-it-does)

Choose between requiring either `Record` type or indexed signature types.

These two types are equivalent, this rule enforces consistency in picking one style over the other:

ts

```
type Foo = Record<string, unknown>;

type Foo = {
  [key: string]: unknown;
};
```

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#why-is-this-bad)

Inconsistent style for indexed object types can harm readability in a project.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#examples)

Examples of **incorrect** code for this rule with `consistent-indexed-object-style: ["error", "record"]` (default):

ts

```
interface Foo {
  [key: string]: unknown;
}
type Foo = {
  [key: string]: unknown;
};
```

Examples of **correct** code for this rule with `consistent-indexed-object-style: ["error", "record"]` (default):

ts

```
type Foo = Record<string, unknown>;
```

Examples of **incorrect** code for this rule with `consistent-indexed-object-style: ["error", "index-signature"]`:

ts

```
type Foo = Record<string, unknown>;
```

Examples of **correct** code for this rule with `consistent-indexed-object-style: ["error", "index-signature"]`:

ts

```
interface Foo {
  [key: string]: unknown;
}
type Foo = {
  [key: string]: unknown;
};
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#configuration)

This rule accepts one of the following string values:

### `"record"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#record)

When set to `record`, enforces the use of a `Record` for indexed object types, e.g. `Record<string, unknown>`.

### `"index-signature"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#index-signature)

When set to `index-signature`, enforces the use of indexed signature types, e.g. `{ [key: string]: unknown }`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/consistent-indexed-object-style": "error"
  }
}
```

bash

```
oxlint --deny typescript/consistent-indexed-object-style
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/consistent_indexed_object_style.rs)
