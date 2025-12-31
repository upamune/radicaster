---
title: "unicorn/prefer-prototype-methods | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.md for this page in Markdown format

# unicorn/prefer-prototype-methods Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#unicorn-prefer-prototype-methods)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#what-it-does)

This rule prefers borrowing methods from the prototype instead of the instance.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#why-is-this-bad)

“Borrowing” a method from an instance of `Array` or `Object` is less clear than getting it from the corresponding prototype.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const array = [].slice.apply(bar);
const type = {}.toString.call(foo);
Reflect.apply([].forEach, arrayLike, [callback]);
```

Examples of **correct** code for this rule:

javascript

```
const array = Array.prototype.slice.apply(bar);
const type = Object.prototype.toString.call(foo);
Reflect.apply(Array.prototype.forEach, arrayLike, [callback]);
const maxValue = Math.max.apply(Math, numbers);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-prototype-methods": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-prototype-methods
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_prototype_methods.rs)
