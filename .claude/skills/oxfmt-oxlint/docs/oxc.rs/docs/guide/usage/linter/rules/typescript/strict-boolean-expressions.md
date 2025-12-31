---
title: "typescript/strict-boolean-expressions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.md for this page in Markdown format

# typescript/strict-boolean-expressions Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#typescript-strict-boolean-expressions)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#what-it-does)

Disallow certain types in boolean expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#why-is-this-bad)

Forbids usage of non-boolean types in expressions where a boolean is expected. `boolean` and `never` types are always allowed. Additional types which are considered safe in a boolean context can be configured via options.

The following nodes are checked:

* Arguments to the `!`, `&&`, and `||` operators
* The condition in a conditional expression (`cond ? x : y`)
* Conditions for `if`, `for`, `while`, and `do-while` statements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#examples)

Examples of **incorrect** code for this rule:

ts

```
const str = "hello";
if (str) {
  console.log("string");
}

const num = 42;
if (num) {
  console.log("number");
}

const obj = { foo: "bar" };
if (obj) {
  console.log("object");
}

declare const maybeString: string | undefined;
if (maybeString) {
  console.log(maybeString);
}

const result = str && num;
const ternary = str ? "yes" : "no";
```

Examples of **correct** code for this rule:

ts

```
const str = "hello";
if (str !== "") {
  console.log("string");
}

const num = 42;
if (num !== 0) {
  console.log("number");
}

const obj = { foo: "bar" };
if (obj !== null) {
  console.log("object");
}

declare const maybeString: string | undefined;
if (maybeString !== undefined) {
  console.log(maybeString);
}

const bool = true;
if (bool) {
  console.log("boolean");
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAny [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allowany)

type: `boolean`

default: `false`

Whether to allow `any` type in boolean contexts.

### allowNullableBoolean [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownullableboolean)

type: `boolean`

default: `false`

Whether to allow nullable boolean types (e.g., `boolean | null`) in boolean contexts.

### allowNullableEnum [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownullableenum)

type: `boolean`

default: `false`

Whether to allow nullable enum types in boolean contexts.

### allowNullableNumber [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownullablenumber)

type: `boolean`

default: `false`

Whether to allow nullable number types (e.g., `number | null`) in boolean contexts.

### allowNullableObject [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownullableobject)

type: `boolean`

default: `true`

Whether to allow nullable object types in boolean contexts.

### allowNullableString [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownullablestring)

type: `boolean`

default: `false`

Whether to allow nullable string types (e.g., `string | null`) in boolean contexts.

### allowNumber [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allownumber)

type: `boolean`

default: `true`

Whether to allow number types in boolean contexts (checks for non-zero numbers).

### allowRuleToRunWithoutStrictNullChecksIKnowWhatIAmDoing [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allowruletorunwithoutstrictnullchecksiknowwhatiamdoing)

type: `boolean`

default: `false`

Whether to allow this rule to run without `strictNullChecks` enabled. This is not recommended as the rule may produce incorrect results.

### allowString [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#allowstring)

type: `boolean`

default: `true`

Whether to allow string types in boolean contexts (checks for non-empty strings).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/strict-boolean-expressions": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/strict-boolean-expressions
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/strict_boolean_expressions.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/strict_boolean_expressions/strict_boolean_expressions.go)
