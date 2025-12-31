---
title: "unicorn/prefer-array-some | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-array-some.md for this page in Markdown format

# unicorn/prefer-array-some Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#unicorn-prefer-array-some)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#what-it-does)

Prefers using [`Array#some()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/some) over [`Array#find()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/find), [`Array#findLast()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findLast) with comparing to undefined, or [`Array#findIndex()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findIndex), [`Array#findLastIndex()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/findLastIndex) and a non-zero length check on the result of [`Array#filter()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array/filter)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#why-is-this-bad)

Using `.some()` is more idiomatic and easier to read.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = array.find(fn) ? bar : baz;
const foo = array.findLast((elem) => hasRole(elem)) !== null;
foo.findIndex(bar) < 0;
foo.findIndex((element) => element.bar === 1) !== -1;
foo.findLastIndex((element) => element.bar === 1) !== -1;
array.filter(fn).length === 0;
```

Examples of **correct** code for this rule:

javascript

```
const foo = array.some(fn) ? bar : baz;
foo.some((element) => element.bar === 1);
!array.some(fn);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-array-some": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-array-some
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_array_some.rs)
