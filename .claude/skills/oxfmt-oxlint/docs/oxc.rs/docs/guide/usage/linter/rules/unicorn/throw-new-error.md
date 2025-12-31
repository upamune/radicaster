---
title: "unicorn/throw-new-error | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/throw-new-error.md for this page in Markdown format

# unicorn/throw-new-error Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#unicorn-throw-new-error)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#what-it-does)

This rule makes sure you always use `new` when throwing an error.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#why-is-this-bad)

In JavaScript, omitting `new` (e.g., `throw Error('message')`) is allowed, but it does not properly initialize the error object. This can lead to missing stack traces or incorrect prototype chains. Using `new` makes the intent clear, ensures consistent behavior, and helps avoid subtle bugs.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
throw Error("🦄");
throw TypeError("unicorn");
throw lib.TypeError("unicorn");
```

Examples of **correct** code for this rule:

javascript

```
throw new Error("🦄");
throw new TypeError("unicorn");
throw new lib.TypeError("unicorn");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/throw-new-error": "error"
  }
}
```

bash

```
oxlint --deny unicorn/throw-new-error
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/throw_new_error.rs)
