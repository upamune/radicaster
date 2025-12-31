---
title: "oxc/no-accumulating-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-accumulating-spread.md for this page in Markdown format

# oxc/no-accumulating-spread Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#oxc-no-accumulating-spread)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#what-it-does)

Prevents using object or array spreads on accumulators in `Array.prototype.reduce()` and in loops.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#why-is-this-bad)

Object and array spreads create a new object or array on each iteration. In the worst case, they also cause O(n) copies (both memory and time complexity). When used on an accumulator, this can lead to `O(n^2)` memory complexity and `O(n^2)` time complexity.

For a more in-depth explanation, see this [blog post](https://prateeksurana.me/blog/why-using-object-spread-with-reduce-bad-idea/) by Prateek Surana.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
arr.reduce((acc, x) => ({ ...acc, [x]: fn(x) }), {});
Object.keys(obj).reduce((acc, el) => ({ ...acc, [el]: fn(el) }), {});

let foo = [];
for (let i = 0; i < 10; i++) {
  foo = [...foo, i];
}
```

Examples of **correct** code for this rule:

javascript

```
function fn(x) {
  // ...
}

arr.reduce((acc, x) => acc.push(fn(x)), []);
Object.keys(obj).reduce((acc, el) => {
  acc[el] = fn(el);
}, {});
// spreading non-accumulators should be avoided if possible, but is not
// banned by this rule
Object.keys(obj).reduce((acc, el) => {
  acc[el] = { ...obj[el] };
  return acc;
}, {});

let foo = [];
for (let i = 0; i < 10; i++) {
  foo.push(i);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-accumulating-spread": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-accumulating-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_accumulating_spread.rs)
