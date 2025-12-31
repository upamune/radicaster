---
title: "typescript/consistent-generic-constructors | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.md for this page in Markdown format

# typescript/consistent-generic-constructors Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#typescript-consistent-generic-constructors)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#what-it-does)

When constructing a generic class, you can specify the type arguments on either the left-hand side (as a type annotation) or the right-hand side (as part of the constructor call).

This rule enforces consistency in the way generic constructors are used.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#why-is-this-bad)

Inconsistent usage of generic constructors can make the code harder to read and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const a: Foo<string> = new Foo();
const a = new Foo<string>(); // prefer type annotation
```

Examples of **correct** code for this rule:

ts

```
const a = new Foo<string>();
const a: Foo<string> = new Foo(); // prefer type annotation
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#configuration)

This rule accepts a configuration object with the following properties:

### option [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#option)

type: `"constructor" | "type-annotation"`

default: `"constructor"`

Specifies where the generic type should be specified.

Possible values:

* `"constructor"` (default): Type arguments that only appear on the type annotation are disallowed.
* `"type-annotation"`: Type arguments that only appear on the constructor are disallowed.

#### `"constructor"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#constructor)

Type arguments that only appear on the type annotation are disallowed.

#### `"type-annotation"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#type-annotation)

Type arguments that only appear on the constructor are disallowed.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/consistent-generic-constructors": "error"
  }
}
```

bash

```
oxlint --deny typescript/consistent-generic-constructors
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/consistent_generic_constructors.rs)
