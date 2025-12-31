---
title: "unicorn/no-await-in-promise-methods | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.md for this page in Markdown format

# unicorn/no-await-in-promise-methods Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#unicorn-no-await-in-promise-methods)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#what-it-does)

Disallow using `await` in `Promise` method parameters

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#why-is-this-bad)

Using `await` on promises passed as arguments to `Promise.all()`, `Promise.allSettled()`, `Promise.any()`, or `Promise.race()` is likely a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function foo() {
  Promise.all([await promise, anotherPromise]);
  Promise.allSettled([await promise, anotherPromise]);
  Promise.any([await promise, anotherPromise]);
  Promise.race([await promise, anotherPromise]);
}
```

Examples of **correct** code for this rule:

javascript

```
async function foo() {
  Promise.all([promise, anotherPromise]);
  Promise.allSettled([promise, anotherPromise]);
  Promise.any([promise, anotherPromise]);
  Promise.race([promise, anotherPromise]);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-await-in-promise-methods": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-await-in-promise-methods
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_await_in_promise_methods.rs)
