---
title: "unicorn/new-for-builtins | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/new-for-builtins.md for this page in Markdown format

# unicorn/new-for-builtins Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#unicorn-new-for-builtins)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#what-it-does)

Enforces the use of `new` for the following builtins: `Object`, `Array`, `ArrayBuffer`, `BigInt64Array`, `BigUint64Array`, `DataView`, `Date`, `Error`, `Float32Array`, `Float64Array`, `Function`, `Int8Array`, `Int16Array`, `Int32Array`, `Map`, `WeakMap`, `Set`, `WeakSet`, `Promise`, `RegExp`, `Uint8Array`, `Uint16Array`, `Uint32Array`, `Uint8ClampedArray`, `SharedArrayBuffer`, `Proxy`, `WeakRef`, `FinalizationRegistry`.

Disallows the use of `new` for the following builtins: `String`, `Number`, `Boolean`, `Symbol`, `BigInt`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#why-is-this-bad)

Using `new` inconsistently can cause confusion. Constructors like `Array` and `RegExp` should always use `new` to ensure the expected instance type. Meanwhile, `String`, `Number`, `Boolean`, `Symbol`, and `BigInt` should not use `new`, as they create object wrappers instead of primitive values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = new String("hello world");
const bar = Array(1, 2, 3);
```

Examples of **correct** code for this rule:

javascript

```
const foo = String("hello world");
const bar = new Array(1, 2, 3);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/new-for-builtins": "error"
  }
}
```

bash

```
oxlint --deny unicorn/new-for-builtins
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/new_for_builtins.rs)
