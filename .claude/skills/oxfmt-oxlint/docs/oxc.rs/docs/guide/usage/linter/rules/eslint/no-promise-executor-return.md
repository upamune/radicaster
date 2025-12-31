---
title: "eslint/no-promise-executor-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-promise-executor-return.md for this page in Markdown format

# eslint/no-promise-executor-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#eslint-no-promise-executor-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#what-it-does)

Disallow returning values from Promise executor functions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#why-is-this-bad)

The `new Promise` constructor accepts an executor function as an argument, which has `resolve` and `reject` parameters that can be used to control the state of the created Promise.

The return value of the executor is ignored. Returning a value from an executor function is a possible error because the returned value cannot be used and it doesn't affect the promise in any way.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
new Promise((resolve, reject) => {
  if (someCondition) {
    return defaultResult;
  }
  getSomething((err, result) => {
    if (err) {
      reject(err);
    } else {
      resolve(result);
    }
  });
});

new Promise((resolve, reject) =>
  getSomething((err, data) => {
    if (err) {
      reject(err);
    } else {
      resolve(data);
    }
  }),
);

new Promise(() => {
  return 1;
});
```

Examples of **correct** code for this rule:

javascript

```
new Promise((resolve, reject) => {
  if (someCondition) {
    resolve(defaultResult);
    return;
  }
  getSomething((err, result) => {
    if (err) {
      reject(err);
    } else {
      resolve(result);
    }
  });
});

new Promise((resolve, reject) => {
  getSomething((err, data) => {
    if (err) {
      reject(err);
    } else {
      resolve(data);
    }
  });
});

new Promise((r) => {
  r(1);
});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#configuration)

This rule accepts a configuration object with the following properties:

### allowVoid [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#allowvoid)

type: `boolean`

default: `false`

If `true`, allows returning `void` expressions (e.g., `return void resolve()`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-promise-executor-return": "error"
  }
}
```

bash

```
oxlint --deny no-promise-executor-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_promise_executor_return.rs)
