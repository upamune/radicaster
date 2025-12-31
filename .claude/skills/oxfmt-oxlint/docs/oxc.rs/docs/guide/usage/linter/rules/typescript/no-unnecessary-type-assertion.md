---
title: "typescript/no-unnecessary-type-assertion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.md for this page in Markdown format

# typescript/no-unnecessary-type-assertion Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#typescript-no-unnecessary-type-assertion)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#what-it-does)

This rule disallows type assertions that do not change the type of an expression.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#why-is-this-bad)

Type assertions that don't actually change the type of an expression are unnecessary and can be safely removed. They add visual noise without providing any benefit and may indicate confusion about TypeScript's type system.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const str: string = "hello";
const redundant = str as string; // unnecessary, str is already string

function getString(): string {
  return "hello";
}
const result = getString() as string; // unnecessary, getString() already returns string

const num = 42;
const alsoRedundant = num as 42; // unnecessary if TypeScript can infer literal type

// Unnecessary assertion to wider type
const literal = "hello" as string;
```

Examples of **correct** code for this rule:

ts

```
const unknown: unknown = "hello";
const str = unknown as string; // necessary to narrow type

const element = document.getElementById("myElement") as HTMLInputElement; // necessary for specific element type

const obj = { name: "John" };
const name = obj.name as const; // necessary for literal type

// No assertion needed
const str2: string = "hello";
const num: number = 42;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#configuration)

This rule accepts a configuration object with the following properties:

### checkLiteralConstAssertions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#checkliteralconstassertions)

type: `boolean`

default: `false`

Whether to check literal const assertions like `'foo' as const`. When `false` (default), const assertions on literal types are not flagged. When `true`, these will be reported as unnecessary since the type is already a literal.

### typesToIgnore [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#typestoignore)

type: `string[]`

default: `[]`

A list of type names to ignore when checking for unnecessary assertions. Type assertions to these types will not be flagged even if they appear unnecessary. Example: `["Foo", "Bar"]` to allow `x as Foo` or `x as Bar`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-type-assertion": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unnecessary-type-assertion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_type_assertion.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unnecessary_type_assertion/no_unnecessary_type_assertion.go)
