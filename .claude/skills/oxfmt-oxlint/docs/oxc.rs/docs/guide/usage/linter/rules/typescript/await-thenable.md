---
title: "typescript/await-thenable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/await-thenable.md for this page in Markdown format

# typescript/await-thenable Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#typescript-await-thenable)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#what-it-does)

This rule disallows awaiting a value that is not a Thenable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#why-is-this-bad)

While it is valid JavaScript to await a non-Promise-like value (it will resolve immediately), this practice can be confusing for readers who are not aware of this behavior. It can also be a sign of a programmer error, such as forgetting to add parentheses to call a function that returns a Promise.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#examples)

Examples of **incorrect** code for this rule:

```
await 12;
await (() => {});

// non-Promise values
await Math.random;
await { then() {} };

// this is not a Promise - it's a function that returns a Promise
declare const getPromise: () => Promise<string>;
await getPromise;
await getPromise;
```

Examples of **correct** code for this rule:

```
await Promise.resolve("value");
await Promise.reject(new Error());

// Promise-like values
await {
  then(onfulfilled, onrejected) {
    onfulfilled("value");
  },
};

// this is a Promise - produced by calling a function
declare const getPromise: () => Promise<string>;
await getPromise();
await getPromise();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/await-thenable": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/await-thenable
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/await_thenable.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/await_thenable/await_thenable.go)
