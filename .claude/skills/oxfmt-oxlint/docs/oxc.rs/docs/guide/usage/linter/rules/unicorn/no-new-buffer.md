---
title: "unicorn/no-new-buffer | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-new-buffer.md for this page in Markdown format

# unicorn/no-new-buffer Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#unicorn-no-new-buffer)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#what-it-does)

Disallows the deprecated `new Buffer()` constructor.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#why-is-this-bad)

Enforces the use of [Buffer.from](https://nodejs.org/api/buffer.html#static-method-bufferfromarray) and [Buffer.alloc()](https://nodejs.org/api/buffer.html#static-method-bufferallocsize-fill-encoding) instead of [new Buffer()](https://nodejs.org/api/buffer.html#new-bufferarray), which has been deprecated since Node.js 4.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const buffer = new Buffer(10);
```

Examples of **correct** code for this rule:

javascript

```
const buffer = Buffer.alloc(10);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-new-buffer": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-new-buffer
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_new_buffer.rs)
