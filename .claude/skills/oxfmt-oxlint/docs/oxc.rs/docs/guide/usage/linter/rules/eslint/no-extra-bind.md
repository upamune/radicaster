---
title: "eslint/no-extra-bind | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-extra-bind.md for this page in Markdown format

# eslint/no-extra-bind Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#eslint-no-extra-bind)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#what-it-does)

Disallow unnecessary calls to .bind()

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#why-is-this-bad)

This rule is aimed at avoiding the unnecessary use of bind() and as such will warn whenever an immediately-invoked function expression (IIFE) is using bind() and doesn’t have an appropriate this value. This rule won’t flag usage of bind() that includes function argument binding.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#examples)

Examples of **incorrect** code for this rule:

js

```
const x = function () {
  foo();
}.bind(bar);

const z = (() => {
  this.foo();
}).bind(this);
```

Examples of **correct** code for this rule:

js

```
const x = function () {
  this.foo();
}.bind(bar);
const y = function (a) {
  return a + 1;
}.bind(foo, bar);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-extra-bind": "error"
  }
}
```

bash

```
oxlint --deny no-extra-bind
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_extra_bind.rs)
