---
title: "oxc/bad-array-method-on-arguments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.md for this page in Markdown format

# oxc/bad-array-method-on-arguments Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#oxc-bad-array-method-on-arguments)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#what-it-does)

This rule applies when an array method is called on the arguments object itself.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#why-is-this-bad)

The [arguments object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Functions/arguments) is not an array, but an array-like object. It should be converted to a real array before calling an array method. Otherwise, a TypeError exception will be thrown because of the non-existent method.

Note that you probably don't need this rule if you are using exclusively TypeScript, as it will catch these errors when typechecking.

`arguments` usage is usually discouraged in modern JavaScript, and you should prefer using rest parameters instead, e.g. `function sum(...args)`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function add(x, y) {
  return x + y;
}
function sum() {
  return arguments.reduce(add, 0);
}
```

Examples of **correct** code for this rule:

javascript

```
function add(x, y) {
  return x + y;
}

function sum(...args) {
  return args.reduce(add, 0);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-array-method-on-arguments": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-array-method-on-arguments
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_array_method_on_arguments.rs)
