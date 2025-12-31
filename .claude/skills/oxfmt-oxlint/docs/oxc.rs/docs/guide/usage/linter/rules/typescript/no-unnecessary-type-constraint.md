---
title: "typescript/no-unnecessary-type-constraint | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.md for this page in Markdown format

# typescript/no-unnecessary-type-constraint Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#typescript-no-unnecessary-type-constraint)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#what-it-does)

Disallow unnecessary constraints on generic types.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#why-is-this-bad)

Generic type parameters (`<T>`) in TypeScript may be "constrained" with an `extends` keyword. When no `extends` is provided, type parameters default a constraint to `unknown`. It is therefore redundant to `extend` from `any` or `unknown`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
interface FooAny<T extends any> {}
interface FooUnknown<T extends unknown> {}

type BarAny<T extends any> = {};
type BarUnknown<T extends unknown> = {};

const QuuxAny = <T extends any>() => {};

function QuuzAny<T extends any>() {}
```

typescript

```
class BazAny<T extends any> {
  quxAny<U extends any>() {}
}
```

Examples of **correct** code for this rule:

typescript

```
interface Foo<T> {}

type Bar<T> = {};

const Quux = <T>() => {};

function Quuz<T>() {}
```

typescript

```
class Baz<T> {
  qux<U>() {}
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-type-constraint": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-unnecessary-type-constraint
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_type_constraint.rs)
