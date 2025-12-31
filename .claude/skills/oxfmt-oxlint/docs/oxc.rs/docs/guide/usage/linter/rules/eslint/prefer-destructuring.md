---
title: "eslint/prefer-destructuring | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/prefer-destructuring.md for this page in Markdown format

# eslint/prefer-destructuring Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#eslint-prefer-destructuring)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#what-it-does)

Require destructuring from arrays and/or objects

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#why-is-this-bad)

With JavaScript ES2015, a new syntax was added for creating variables from an array index or object property, called destructuring. This rule enforces usage of destructuring instead of accessing a property through a member expression.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#examples)

Examples of **incorrect** code for this rule:

js

```
// With `array` enabled
const foo = array[0];
bar.baz = array[0];
// With `object` enabled
const qux = object.qux;
const quux = object["quux"];
```

Examples of **correct** code for this rule:

js

```
// With `array` enabled
const [foo] = array;
const arr = array[someIndex];
[bar.baz] = array;

// With `object` enabled
const { baz } = object;
const obj = object.bar;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#configuration)

This rule accepts a configuration object with the following properties:

### AssignmentExpression [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#assignmentexpression)

type: `object`

default: `{"array":true, "object":true}`

Configuration for destructuring in assignment expressions, configured for arrays and objects independently.

#### AssignmentExpression.array [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#assignmentexpression-array)

type: `boolean`

default: `true`

#### AssignmentExpression.object [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#assignmentexpression-object)

type: `boolean`

default: `true`

### VariableDeclarator [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#variabledeclarator)

type: `object`

default: `{"array":true, "object":true}`

Configuration for destructuring in variable declarations, configured for arrays and objects independently.

#### VariableDeclarator.array [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#variabledeclarator-array)

type: `boolean`

default: `true`

#### VariableDeclarator.object [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#variabledeclarator-object)

type: `boolean`

default: `true`

### enforceForRenamedProperties [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#enforceforrenamedproperties)

type: `boolean`

default: `false`

Determines whether the object destructuring rule applies to renamed variables.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "prefer-destructuring": "error"
  }
}
```

bash

```
oxlint --deny prefer-destructuring
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/prefer_destructuring.rs)
