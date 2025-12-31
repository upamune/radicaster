---
title: "oxc/no-async-await | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-async-await.md for this page in Markdown format

# oxc/no-async-await Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#oxc-no-async-await)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#what-it-does)

Disallows the use of `async`/`await`.

This rule should generally not be used in modern JavaScript/TypeScript codebases without good reason.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#why-is-this-bad)

This rule is useful for environments that don't support `async`/`await` syntax, or when you want to enforce the use of promises or other asynchronous patterns instead. It can also be used to maintain consistency in codebases that use alternative async patterns.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
async function foo() {
  await bar();
  return baz();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-async-await": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-async-await
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_async_await.rs)
