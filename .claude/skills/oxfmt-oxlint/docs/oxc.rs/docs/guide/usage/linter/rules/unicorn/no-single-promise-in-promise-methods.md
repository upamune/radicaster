---
title: "unicorn/no-single-promise-in-promise-methods | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.md for this page in Markdown format

# unicorn/no-single-promise-in-promise-methods Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#unicorn-no-single-promise-in-promise-methods)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#what-it-does)

Disallow passing single-element arrays to Promise methods

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#why-is-this-bad)

Passing a single-element array to `Promise.all()`, `Promise.any()`, or `Promise.race()` is likely a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function bad() {
  const foo = await Promise.all([promise]);
  const foo = await Promise.any([promise]);
  const foo = await Promise.race([promise]);
  const promise = Promise.all([nonPromise]);
}
```

Examples of **correct** code for this rule:

javascript

```
async function good() {
  const foo = await promise;
  const promise = Promise.resolve(nonPromise);
  const foo = await Promise.all(promises);
  const foo = await Promise.any([promise, anotherPromise]);
  const [{ value: foo, reason: error }] = await Promise.allSettled([promise]);
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-single-promise-in-promise-methods": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-single-promise-in-promise-methods
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_single_promise_in_promise_methods.rs)
