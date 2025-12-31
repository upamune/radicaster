---
title: "unicorn/no-array-method-this-argument | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.md for this page in Markdown format

# unicorn/no-array-method-this-argument Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#unicorn-no-array-method-this-argument)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#what-it-does)

Disallows the use of the `thisArg` parameter in array iteration methods such as `map`, `filter`, `some`, `every`, and similar.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#why-is-this-bad)

The `thisArg` parameter makes code harder to understand and reason about. Instead, prefer arrow functions or bind explicitly in a clearer way. Arrow functions inherit `this` from the lexical scope, which is more intuitive and less error-prone.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#examples)

Examples of **incorrect** code for this rule:

js

```
array.map(function (x) {
  return x + this.y;
}, this);
array.filter(function (x) {
  return x !== this.value;
}, this);
```

Examples of **correct** code for this rule:

js

```
array.map((x) => x + this.y);
array.filter((x) => x !== this.value);
const self = this;
array.map(function (x) {
  return x + self.y;
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-array-method-this-argument": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-array-method-this-argument
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_array_method_this_argument.rs)
