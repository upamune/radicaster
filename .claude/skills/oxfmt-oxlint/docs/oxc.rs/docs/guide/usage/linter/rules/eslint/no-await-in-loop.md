---
title: "eslint/no-await-in-loop | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-await-in-loop.md for this page in Markdown format

# eslint/no-await-in-loop Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#eslint-no-await-in-loop)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#what-it-does)

This rule disallows the use of `await` within loop bodies. (for, for-in, for-of, while, do-while).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#why-is-this-bad)

It potentially indicates that the async operations are not being effectively parallelized. Instead, they are being run in series, which can lead to poorer performance.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function bad() {
  for (const user of users) {
    const userRecord = await getUserRecord(user);
  }
}
```

Examples of **correct** code for this rule:

javascript

```
async function good() {
  await Promise.all(users.map((user) => getUserRecord(user)));
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-await-in-loop": "error"
  }
}
```

bash

```
oxlint --deny no-await-in-loop
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_await_in_loop.rs)
