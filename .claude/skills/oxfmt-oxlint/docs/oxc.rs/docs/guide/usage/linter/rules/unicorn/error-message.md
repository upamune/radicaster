---
title: "unicorn/error-message | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/error-message.md for this page in Markdown format

# unicorn/error-message Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#unicorn-error-message)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#what-it-does)

Enforces providing a `message` when creating built-in `Error` objects to improve readability and debugging.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#why-is-this-bad)

Throwing an `Error` without a message, like `throw new Error()`, provides no context on what went wrong, making debugging harder. A clear error message improves code clarity and helps developers quickly identify issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
throw Error();
throw new TypeError();
```

Examples of **correct** code for this rule:

javascript

```
throw new Error("Unexpected token");
throw new TypeError("Number expected");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/error-message": "error"
  }
}
```

bash

```
oxlint --deny unicorn/error-message
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/error_message.rs)
