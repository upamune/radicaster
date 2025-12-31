---
title: "eslint/no-func-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-func-assign.md for this page in Markdown format

# eslint/no-func-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#eslint-no-func-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#what-it-does)

Disallow reassigning `function` declarations.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#why-is-this-bad)

Overwriting/reassigning a function written as a FunctionDeclaration is often indicative of a mistake or issue.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo() {}
foo = bar;
```

javascript

```
function foo() {
  foo = bar;
}
```

javascript

```
let a = function hello() {
  hello = 123;
};
```

Examples of **correct** code for this rule:

javascript

```
let foo = function () {};
foo = bar;
```

javascript

```
function baz(baz) {
  // `baz` is shadowed.
  baz = bar;
}
```

```
function qux() {
  const qux = bar;  // `qux` is shadowed.
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-func-assign": "error"
  }
}
```

bash

```
oxlint --deny no-func-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_func_assign.rs)
