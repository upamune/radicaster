---
title: "promise/prefer-await-to-then | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/promise/prefer-await-to-then.md for this page in Markdown format

# promise/prefer-await-to-then Style [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#promise-prefer-await-to-then)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#what-it-does)

Prefer `await` to `then()`/`catch()`/`finally()` for reading Promise values

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#why-is-this-bad)

Async/await syntax can be seen as more readable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function foo() {
  hey.then((x) => {});
}
```

Examples of **correct** code for this rule:

javascript

```
async function hi() {
  await thing();
}
```

### Example with strict mode [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#example-with-strict-mode)

Examples of **incorrect** code with `{ strict: true }`:

javascript

```
async function hi() {
  await thing().then((x) => {});
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#configuration)

This rule accepts a configuration object with the following properties:

### strict [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#strict)

type: `boolean`

default: `false`

If true, enforces the rule even after an `await` or `yield` expression.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["promise"],
  "rules": {
    "promise/prefer-await-to-then": "error"
  }
}
```

bash

```
oxlint --deny promise/prefer-await-to-then --promise-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/promise/prefer_await_to_then.rs)
