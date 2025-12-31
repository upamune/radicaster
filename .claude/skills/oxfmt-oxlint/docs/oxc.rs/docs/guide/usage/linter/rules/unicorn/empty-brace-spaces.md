---
title: "unicorn/empty-brace-spaces | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.md for this page in Markdown format

# unicorn/empty-brace-spaces Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#unicorn-empty-brace-spaces)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#what-it-does)

Removes the extra spaces or new line characters inside a pair of braces that does not contain additional code. This ensures that braces are clean and do not contain unnecessary spaces or newlines.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#why-is-this-bad)

Extra spaces inside braces can negatively impact the readability of the code. Keeping braces clean and free of unnecessary characters improves consistency and makes the code easier to understand and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const a = {};
class A {}
```

Examples of **correct** code for this rule:

javascript

```
const a = {};
class A {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/empty-brace-spaces": "error"
  }
}
```

bash

```
oxlint --deny unicorn/empty-brace-spaces
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/empty_brace_spaces.rs)
