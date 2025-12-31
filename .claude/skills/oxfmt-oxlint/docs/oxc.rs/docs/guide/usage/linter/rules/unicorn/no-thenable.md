---
title: "unicorn/no-thenable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-thenable.md for this page in Markdown format

# unicorn/no-thenable Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#unicorn-no-thenable)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#what-it-does)

Disallow `then` property

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#why-is-this-bad)

If an object is defined as "thenable", once it's accidentally used in an await expression, it may cause problems:

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function example() {
  const foo = {
    unicorn: 1,
    then() {},
  };

  const { unicorn } = await foo;

  console.log("after"); //<- This will never execute
}
```

Examples of **correct** code for this rule:

javascript

```
async function example() {
  const foo = {
    unicorn: 1,
    bar() {},
  };

  const { unicorn } = await foo;

  console.log("after");
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-thenable": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-thenable
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_thenable.rs)
