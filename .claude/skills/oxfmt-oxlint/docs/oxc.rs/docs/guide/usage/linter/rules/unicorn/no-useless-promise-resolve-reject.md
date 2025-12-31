---
title: "unicorn/no-useless-promise-resolve-reject | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.md for this page in Markdown format

# unicorn/no-useless-promise-resolve-reject Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#unicorn-no-useless-promise-resolve-reject)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#what-it-does)

Disallows returning values wrapped in `Promise.resolve` or `Promise.reject` in an async function or a `Promise#then`/`catch`/`finally` callback.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#why-is-this-bad)

Wrapping a return value in `Promise.resolve` in an async function or a `Promise#then`/`catch`/`finally` callback is unnecessary as all return values in async functions and promise callback functions are already wrapped in a `Promise`. Similarly, returning an error wrapped in `Promise.reject` is equivalent to simply `throw`ing the error. This is the same for `yield`ing in async generators as well.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async () => Promise.resolve(bar);
```

Examples of **correct** code for this rule:

javascript

```
async () => bar;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#configuration)

This rule accepts a configuration object with the following properties:

### allowReject [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#allowreject)

type: `boolean`

default: `false`

If set to `true`, allows the use of `Promise.reject` in async functions and promise callbacks.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-promise-resolve-reject": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-promise-resolve-reject
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_promise_resolve_reject.rs)
