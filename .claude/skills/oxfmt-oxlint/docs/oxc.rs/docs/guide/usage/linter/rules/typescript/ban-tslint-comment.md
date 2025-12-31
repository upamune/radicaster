---
title: "typescript/ban-tslint-comment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/ban-tslint-comment.md for this page in Markdown format

# typescript/ban-tslint-comment Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#typescript-ban-tslint-comment)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#what-it-does)

This rule disallows `tslint:<rule-flag>` comments

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#why-is-this-bad)

Useful when migrating from TSLint to ESLint. Once TSLint has been removed, this rule helps locate TSLint annotations

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#examples)

Examples of **incorrect** code for this rule:

ts

```
// tslint:disable-next-line
someCode();
```

Examples of **correct** code for this rule:

ts

```
someCode();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/ban-tslint-comment": "error"
  }
}
```

bash

```
oxlint --deny typescript/ban-tslint-comment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/ban_tslint_comment.rs)
