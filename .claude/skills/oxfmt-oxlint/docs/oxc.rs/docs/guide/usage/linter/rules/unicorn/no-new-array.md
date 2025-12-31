---
title: "unicorn/no-new-array | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-new-array.md for this page in Markdown format

# unicorn/no-new-array Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#unicorn-no-new-array)

✅ This rule is turned on by default.

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#what-it-does)

Disallow `new Array()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#why-is-this-bad)

When using the `Array` constructor with one argument, it's not clear whether the argument is meant to be the length of the array or the only element.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const array = new Array(1);
const array = new Array(42);
const array = new Array(foo);
```

Examples of **correct** code for this rule:

javascript

```
const array = Array.from({ length: 42 });
const array = [42];
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-new-array": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-new-array
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_new_array.rs)
