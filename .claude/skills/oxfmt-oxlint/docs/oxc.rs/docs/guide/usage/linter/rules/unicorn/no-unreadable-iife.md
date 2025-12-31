---
title: "unicorn/no-unreadable-iife | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.md for this page in Markdown format

# unicorn/no-unreadable-iife Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#unicorn-no-unreadable-iife)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#what-it-does)

This rule disallows IIFEs with a parenthesized arrow function body.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#why-is-this-bad)

IIFEs with a parenthesized arrow function body are unreadable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = ((bar) => (bar ? bar.baz : baz))(getBar());

const foo = ((bar, baz) => ({ bar, baz }))(bar, baz);
```

Examples of **correct** code for this rule:

javascript

```
const bar = getBar();
const foo = bar ? bar.baz : baz;

const getBaz = (bar) => (bar ? bar.baz : baz);
const foo = getBaz(getBar());

const foo = ((bar) => {
  return bar ? bar.baz : baz;
})(getBar());
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-unreadable-iife": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-unreadable-iife
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_unreadable_iife.rs)
