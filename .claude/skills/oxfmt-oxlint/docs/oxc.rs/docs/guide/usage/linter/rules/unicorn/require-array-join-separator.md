---
title: "unicorn/require-array-join-separator | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/require-array-join-separator.md for this page in Markdown format

# unicorn/require-array-join-separator Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#unicorn-require-array-join-separator)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#what-it-does)

Enforce using the separator argument with Array#join()

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#why-is-this-bad)

It's better to make it clear what the separator is when calling Array#join(), instead of relying on the default comma (',') separator.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.join();
```

Examples of **correct** code for this rule:

javascript

```
foo.join(",");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/require-array-join-separator": "error"
  }
}
```

bash

```
oxlint --deny unicorn/require-array-join-separator
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/require_array_join_separator.rs)
