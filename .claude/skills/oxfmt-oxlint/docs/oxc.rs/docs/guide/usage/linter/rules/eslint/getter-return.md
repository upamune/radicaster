---
title: "eslint/getter-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/getter-return.md for this page in Markdown format

# eslint/getter-return Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#eslint-getter-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#what-it-does)

Requires all getters to have a `return` statement.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#why-is-this-bad)

Getters should always return a value. If they don't, it's probably a mistake.

This rule does not run on TypeScript files, since type checking will catch getters that do not return a value.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class Person {
  get name() {
    // no return
  }
}

const obj = {
  get foo() {
    // object getter are also checked
  },
};
```

Examples of **correct** code for this rule:

javascript

```
class Person {
  get name() {
    return this._name;
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#configuration)

This rule accepts a configuration object with the following properties:

### allowImplicit [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#allowimplicit)

type: `boolean`

default: `false`

When set to `true`, allows getters to implicitly return `undefined` with a `return` statement containing no expression.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "getter-return": "error"
  }
}
```

bash

```
oxlint --deny getter-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/getter_return.rs)
