---
title: "eslint/max-params | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-params.md for this page in Markdown format

# eslint/max-params Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#eslint-max-params)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#what-it-does)

Enforce a maximum number of parameters in function definitions which by default is three.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#why-is-this-bad)

Functions that take numerous parameters can be difficult to read and write because it requires the memorization of what each parameter is, its type, and the order they should appear in. As a result, many coders adhere to a convention that caps the number of parameters a function can take.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo(bar, baz, qux, qxx) {
  doSomething();
}
```

javascript

```
let foo = (bar, baz, qux, qxx) => {
  doSomething();
};
```

Examples of **correct** code for this rule:

javascript

```
function foo(bar, baz, qux) {
  doSomething();
}
```

javascript

```
let foo = (bar, baz, qux) => {
  doSomething();
};
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#configuration)

This rule accepts a configuration object with the following properties:

### countVoidThis [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#countvoidthis)

type: `boolean`

default: `false`

This option is for counting the `this` parameter if it is of type `void`.

For example `{ "countVoidThis": true }` would mean that having a function take a `this` parameter of type `void` is counted towards the maximum number of parameters.

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#max)

type: `integer`

default: `3`

Maximum number of parameters allowed in function definitions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-params": "error"
  }
}
```

bash

```
oxlint --deny max-params
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_params.rs)
