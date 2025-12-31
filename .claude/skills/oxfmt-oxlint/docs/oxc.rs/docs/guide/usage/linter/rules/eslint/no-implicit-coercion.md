---
title: "eslint/no-implicit-coercion | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-implicit-coercion.md for this page in Markdown format

# eslint/no-implicit-coercion Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#eslint-no-implicit-coercion)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#what-it-does)

Disallows shorthand type conversions using operators like `!!`, `+`, `""+` , etc.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#why-is-this-bad)

Implicit type coercions using operators can be less clear than using explicit type conversion functions like `Boolean()`, `Number()`, and `String()`. Using explicit conversions makes the intent clearer and the code more readable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var b = !!foo;
var n = +foo;
var s = "" + foo;
```

Examples of **correct** code for this rule:

javascript

```
var b = Boolean(foo);
var n = Number(foo);
var s = String(foo);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#allow)

type: `string[]`

List of operators to allow. Valid values: `"!!"`, `"~"`, `"+"`, `"-"`, `"- -"`, `"*"`

### boolean [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#boolean)

type: `boolean`

default: `true`

When `true`, warns on implicit boolean coercion (e.g., `!!foo`).

### disallowTemplateShorthand [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#disallowtemplateshorthand)

type: `boolean`

default: `false`

When `true`, disallows using template literals for string coercion (e.g., `` `${foo}` ``).

### number [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#number)

type: `boolean`

default: `true`

When `true`, warns on implicit number coercion (e.g., `+foo`).

### string [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#string)

type: `boolean`

default: `true`

When `true`, warns on implicit string coercion (e.g., `"" + foo`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-implicit-coercion": "error"
  }
}
```

bash

```
oxlint --deny no-implicit-coercion
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_implicit_coercion.rs)
