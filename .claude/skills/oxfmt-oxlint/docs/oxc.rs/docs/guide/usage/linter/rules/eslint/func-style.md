---
title: "eslint/func-style | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/func-style.md for this page in Markdown format

# eslint/func-style Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#eslint-func-style)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#what-it-does)

Enforce the consistent use of either function declarations or expressions assigned to variables

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#why-is-this-bad)

This rule enforces a particular type of function style, either function declarations or expressions assigned to variables. You can specify which you prefer in the configuration.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#examples)

js

```
// function declaration
function doSomething() {
  // ...
}

// arrow function expression assigned to a variable
const doSomethingElse = () => {
  // ...
};

// function expression assigned to a variable
const doSomethingAgain = function () {
  // ...
};
```

Examples of incorrect code for this rule with the default "expression" option:

js

```
/*eslint func-style: ["error", "expression"]*/

function foo() {
  // ...
}
```

Examples of incorrect code for this rule with the "declaration" option:

js

```
/*eslint func-style: ["error", "declaration"]*/
var foo = function () {
  // ...
};

var foo = () => {};
```

Examples of incorrect code for this rule with the "declaration" and {"overrides": { "namedExports": "expression" }} option:

js

```
/*eslint func-style: ["error", "declaration", { "overrides": { "namedExports": "expression" } }]*/
export function foo() {
  // ...
}
```

Examples of incorrect code for this rule with the "expression" and {"overrides": { "namedExports": "declaration" }} option:

js

```
/*eslint func-style: ["error", "expression", { "overrides": { "namedExports": "declaration" } }]*/
export var foo = function () {
  // ...
};

export var bar = () => {};
```

Examples of correct code for this rule with the default "expression" option:

js

```
/*eslint func-style: ["error", "expression"]*/
var foo = function () {
  // ...
};
```

Examples of correct code for this rule with the "declaration" option:

js

```
/*eslint func-style: ["error", "declaration"]*/
function foo() {
  // ...
}
// Methods (functions assigned to objects) are not checked by this rule
SomeObject.foo = function () {
  // ...
};
```

Examples of additional correct code for this rule with the "declaration", { "allowArrowFunctions": true } options:

js

```
/*eslint func-style: ["error", "declaration", { "allowArrowFunctions": true }]*/
var foo = () => {};
```

Examples of correct code for this rule with the "declaration" and {"overrides": { "namedExports": "expression" }} option:

js

```
/*eslint func-style: ["error", "declaration", { "overrides": { "namedExports": "expression" } }]*/
export var foo = function () {
  // ...
};
export var bar = () => {};
```

Examples of correct code for this rule with the "expression" and {"overrides": { "namedExports": "declaration" }} option:

js

```
/*eslint func-style: ["error", "expression", { "overrides": { "namedExports": "declaration" } }]*/
export function foo() {
  // ...
}
```

Examples of correct code for this rule with the {"overrides": { "namedExports": "ignore" }} option:

js

```
/*eslint func-style: ["error", "expression", { "overrides": { "namedExports": "ignore" } }]*/
export var foo = function () {
  // ...
};

export var bar = () => {};
export function baz() {
  // ...
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#configuration)

This rule accepts a configuration object with the following properties:

### allowArrowFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#allowarrowfunctions)

type: `boolean`

default: `false`

When true, arrow functions are allowed regardless of the style setting.

### allowTypeAnnotation [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#allowtypeannotation)

type: `boolean`

default: `false`

When true, functions with type annotations are allowed regardless of the style setting.

### namedExports [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#namedexports)

type: `string | null`

default: `null`

Override the style specifically for named exports. Can be "expression", "declaration", or "ignore" (default).

### style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#style)

type: `"expression" | "declaration"`

default: `"expression"`

The style to enforce. Either "expression" (default) or "declaration".

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "func-style": "error"
  }
}
```

bash

```
oxlint --deny func-style
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/func_style.rs)
