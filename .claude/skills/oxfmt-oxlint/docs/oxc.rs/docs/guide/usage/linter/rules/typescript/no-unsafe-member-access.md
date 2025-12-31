---
title: "typescript/no-unsafe-member-access | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.md for this page in Markdown format

# typescript/no-unsafe-member-access Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#typescript-no-unsafe-member-access)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#what-it-does)

This rule disallows member access on a value with type `any`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#why-is-this-bad)

The `any` type in TypeScript disables type checking. When you access a member (property or method) on a value typed as `any`, TypeScript cannot verify that the member exists or what type it has. This can lead to runtime errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const anyValue: any;

anyValue.foo; // unsafe member access

anyValue.bar.baz; // unsafe nested member access

anyValue["key"]; // unsafe computed member access

const result = anyValue.method(); // unsafe method access
```

Examples of **correct** code for this rule:

ts

```
declare const obj: { foo: string; bar: { baz: number } };
declare const unknownValue: unknown;

obj.foo; // safe

obj.bar.baz; // safe

obj["foo"]; // safe

// Type guard for unknown
if (typeof unknownValue === "object" && unknownValue !== null && "foo" in unknownValue) {
  console.log(unknownValue.foo); // safe after type guard
}

// Explicit type assertion if needed
(anyValue as { foo: string }).foo; // explicitly unsafe but intentional
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#configuration)

This rule accepts a configuration object with the following properties:

### allowOptionalChaining [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#allowoptionalchaining)

type: `boolean`

default: `false`

Whether to allow `?.` optional chains on `any` values. When `true`, optional chaining on `any` values will not be flagged. Default is `false`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unsafe-member-access": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unsafe-member-access
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unsafe_member_access.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unsafe_member_access/no_unsafe_member_access.go)
