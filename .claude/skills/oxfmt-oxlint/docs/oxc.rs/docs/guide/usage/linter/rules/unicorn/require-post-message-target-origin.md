---
title: "unicorn/require-post-message-target-origin | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.md for this page in Markdown format

# unicorn/require-post-message-target-origin Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#unicorn-require-post-message-target-origin)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#what-it-does)

Enforce using the targetOrigin argument with window.postMessage()

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#why-is-this-bad)

When calling window.postMessage() without the targetOrigin argument, the message cannot be received by any window.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#examples)

Examples of **incorrect** code for this rule:

js

```
window.postMessage(message);
```

Examples of **correct** code for this rule:

js

```
window.postMessage(message, "https://example.com");

window.postMessage(message, "*");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/require-post-message-target-origin": "error"
  }
}
```

bash

```
oxlint --deny unicorn/require-post-message-target-origin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/require_post_message_target_origin.rs)
