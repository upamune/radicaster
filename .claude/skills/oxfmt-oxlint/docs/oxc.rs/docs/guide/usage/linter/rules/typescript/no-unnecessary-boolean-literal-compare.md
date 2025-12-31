---
title: "typescript/no-unnecessary-boolean-literal-compare | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.md for this page in Markdown format

# typescript/no-unnecessary-boolean-literal-compare Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#typescript-no-unnecessary-boolean-literal-compare)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#what-it-does)

This rule disallows unnecessary equality comparisons with boolean literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#why-is-this-bad)

Comparing boolean values to boolean literals is unnecessary when the comparison can be eliminated. These comparisons make code more verbose without adding value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#examples)

Examples of **incorrect** code for this rule:

ts

```
declare const someCondition: boolean;

if (someCondition === true) {
  // ...
}

if (someCondition === false) {
  // ...
}

if (someCondition !== true) {
  // ...
}

if (someCondition !== false) {
  // ...
}

const result = someCondition == true;
```

Examples of **correct** code for this rule:

ts

```
declare const someCondition: boolean;

if (someCondition) {
  // ...
}

if (!someCondition) {
  // ...
}

// Comparisons with non-boolean types are allowed
declare const someValue: unknown;
if (someValue === true) {
  // ...
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#configuration)

This rule accepts a configuration object with the following properties:

### allowComparingNullableBooleansToFalse [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#allowcomparingnullablebooleanstofalse)

type: `boolean`

default: `true`

Whether to allow comparing nullable boolean expressions to `false`. When false, `x === false` where x is `boolean | null` will be flagged.

### allowComparingNullableBooleansToTrue [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#allowcomparingnullablebooleanstotrue)

type: `boolean`

default: `true`

Whether to allow comparing nullable boolean expressions to `true`. When false, `x === true` where x is `boolean | null` will be flagged.

### allowRuleToRunWithoutStrictNullChecksIKnowWhatIAmDoing [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#allowruletorunwithoutstrictnullchecksiknowwhatiamdoing)

type: `boolean`

default: `false`

Whether to allow this rule to run without `strictNullChecks` enabled. This is not recommended as the rule may produce incorrect results.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-boolean-literal-compare": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-unnecessary-boolean-literal-compare
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_boolean_literal_compare.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_unnecessary_boolean_literal_compare/no_unnecessary_boolean_literal_compare.go)
