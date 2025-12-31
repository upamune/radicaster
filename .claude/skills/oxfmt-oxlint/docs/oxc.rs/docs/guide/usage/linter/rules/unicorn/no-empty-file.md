---
title: "unicorn/no-empty-file | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-empty-file.md for this page in Markdown format

# unicorn/no-empty-file Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html#unicorn-no-empty-file)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html#what-it-does)

Disallows files that do not contain any meaningful code.

This includes files that consist only of:

* Whitespace
* Comments
* Directives (e.g., `"use strict"`)
* Empty statements (`;`)
* Empty blocks (`{}`)
* Hashbangs (`#!/usr/bin/env node`)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html#why-is-this-bad)

Files with no executable or exportable content are typically unintentional or left over from refactoring. They clutter the codebase and may confuse tooling or developers by appearing to serve a purpose when they do not.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-empty-file": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-empty-file
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_empty_file.rs)
