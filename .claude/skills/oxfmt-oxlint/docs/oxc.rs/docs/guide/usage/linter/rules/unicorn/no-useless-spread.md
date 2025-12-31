---
title: "unicorn/no-useless-spread | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-spread.md for this page in Markdown format

# unicorn/no-useless-spread Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#unicorn-no-useless-spread)

✅ This rule is turned on by default.

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#what-it-does)

Disallows using spread syntax in following, unnecessary cases:

* Spread an array literal as elements of an array literal
* Spread an array literal as arguments of a call or a `new` call
* Spread an object literal as properties of an object literal
* Use spread syntax to clone an array created inline

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#why-is-this-bad)

The following builtins accept an iterable, so it's unnecessary to convert the iterable to an array:

* `Map` constructor
* `WeakMap` constructor
* `Set` constructor
* `WeakSet` constructor
* `TypedArray` constructor
* `Array.from(…)`
* `TypedArray.from(…)`
* `Promise.{all,allSettled,any,race}(…)`
* `Object.fromEntries(…)`

The `for…of` loop can iterate over any iterable object not just array, so it's unnecessary to convert the iterable to an array.

The `yield*` can delegate to another iterable, so it's unnecessary to convert the iterable to an array.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const array = [firstElement, ...[secondElement], thirdElement];

await Promise.all([...iterable]);

for (const foo of [...set]);

function* foo() {
  yield* [...anotherGenerator()];
}

function foo(bar) {
  return [...bar.map((x) => x * 2)];
}
```

Examples of **correct** code for this rule:

javascript

```
const array = [firstElement, secondElement, thirdElement];

await Promise.all(iterable);

for (const foo of set);

function* foo() {
  yield* anotherGenerator();
}

function foo(bar) {
  return bar.map((x) => x * 2);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-spread": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-spread
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_spread/mod.rs)
