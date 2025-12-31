---
title: "unicorn/no-useless-error-capture-stack-trace | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.md for this page in Markdown format

# unicorn/no-useless-error-capture-stack-trace Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#unicorn-no-useless-error-capture-stack-trace)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#what-it-does)

Disallows unnecessary `Error.captureStackTrace(…)` in error constructors.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#why-is-this-bad)

Calling `Error.captureStackTrace(…)` inside the constructor of a built-in `Error` subclass is unnecessary, since the `Error` constructor calls it automatically.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#examples)

Examples of **incorrect** code for this rule:

js

```
class MyError extends Error {
  constructor() {
    Error.captureStackTrace(this, MyError);
  }
}
```

Examples of **correct** code for this rule:

js

```
class MyError extends Error {
  constructor() {
    // No need to call Error.captureStackTrace
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-error-capture-stack-trace": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-error-capture-stack-trace
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_error_capture_stack_trace.rs)
