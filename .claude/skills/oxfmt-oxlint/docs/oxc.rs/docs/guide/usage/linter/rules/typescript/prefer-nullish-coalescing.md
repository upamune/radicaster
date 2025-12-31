---
title: "typescript/prefer-nullish-coalescing | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.md for this page in Markdown format

# typescript/prefer-nullish-coalescing Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#typescript-prefer-nullish-coalescing)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#what-it-does)

Enforce using the nullish coalescing operator (`??`) instead of logical OR (`||`) or conditional expressions when the left operand might be `null` or `undefined`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#why-is-this-bad)

The `||` operator returns the right-hand side when the left-hand side is any falsy value (`false`, `0`, `''`, `null`, `undefined`, `NaN`). This can lead to unexpected behavior when you only want to provide a default for `null` or `undefined`.

The nullish coalescing operator (`??`) only returns the right-hand side when the left-hand side is `null` or `undefined`, making the intent clearer and avoiding bugs with other falsy values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const x: string | null;

// Using || when ?? would be more appropriate
const foo = x || "default";

// Ternary that could use ??
const bar = x !== null && x !== undefined ? x : "default";
const baz = x != null ? x : "default";

// If statement that could use ??
let value = "default";
if (x !== null && x !== undefined) {
  value = x;
}
```

Examples of **correct** code for this rule:

ts

```
declare const x: string | null;

// Using nullish coalescing
const foo = x ?? "default";

// || is fine when you want falsy behavior
declare const y: string;
const bar = y || "default";

// Boolean coercion (can be ignored with ignoreBooleanCoercion)
const bool = Boolean(x || y);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#configuration)

This rule accepts a configuration object with the following properties:

### allowRuleToRunWithoutStrictNullChecksIKnowWhatIAmDoing [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#allowruletorunwithoutstrictnullchecksiknowwhatiamdoing)

type: `boolean`

default: `false`

Unless this is set to `true`, the rule will error on every file whose `tsconfig.json` does *not* have the `strictNullChecks` compiler option (or `strict`) set to `true`.

It is *not* recommended to enable this config option.

### ignoreBooleanCoercion [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignorebooleancoercion)

type: `boolean`

default: `false`

Whether to ignore arguments to the `Boolean` constructor.

### ignoreConditionalTests [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignoreconditionaltests)

type: `boolean`

default: `true`

Whether to ignore cases that are located within a conditional test.

### ignoreIfStatements [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignoreifstatements)

type: `boolean`

default: `false`

Whether to ignore any if statements that could be simplified by using the nullish coalescing operator.

### ignoreMixedLogicalExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignoremixedlogicalexpressions)

type: `boolean`

default: `false`

Whether to ignore any logical or expressions that are part of a mixed logical expression (with `&&`).

### ignorePrimitives [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignoreprimitives)

type: `boolean`

Represents the different ways `ignorePrimitives` can be specified in JSON. Can be:

* `true` - ignore all primitive types
* An object specifying which primitives to ignore

### ignoreTernaryTests [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#ignoreternarytests)

type: `boolean`

default: `false`

Whether to ignore any ternary expressions that could be simplified by using the nullish coalescing operator.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-nullish-coalescing": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/prefer-nullish-coalescing
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_nullish_coalescing.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/prefer_nullish_coalescing/prefer_nullish_coalescing.go)
