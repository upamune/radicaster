---
title: "eslint/no-new-func | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-new-func.md for this page in Markdown format

# eslint/no-new-func Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#eslint-no-new-func)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#what-it-does)

The rule disallow `new` operators with the `Function` object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#why-is-this-bad)

Using `new Function` or `Function` can lead to code that is difficult to understand and maintain. It can introduce security risks similar to those associated with `eval` because it generates a new function from a string of code, which can be a vector for injection attacks. Additionally, it impacts performance negatively as these functions are not optimized by the JavaScript engine.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#examples)

Examples of **incorrect** code for this rule:

js

```
var x = new Function("a", "b", "return a + b");
var x = Function("a", "b", "return a + b");
var x = Function.call(null, "a", "b", "return a + b");
var x = Function.apply(null, ["a", "b", "return a + b"]);
var x = Function.bind(null, "a", "b", "return a + b")();
var f = Function.bind(null, "a", "b", "return a + b");
```

Examples of **correct** code for this rule:

js

```
let x = function (a, b) {
  return a + b;
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-new-func": "error"
  }
}
```

bash

```
oxlint --deny no-new-func
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_new_func.rs)
