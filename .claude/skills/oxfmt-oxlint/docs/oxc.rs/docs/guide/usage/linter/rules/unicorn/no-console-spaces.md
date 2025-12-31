---
title: "unicorn/no-console-spaces | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-console-spaces.md for this page in Markdown format

# unicorn/no-console-spaces Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#unicorn-no-console-spaces)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#what-it-does)

Disallows leading/trailing space inside `console.log()` and similar methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#why-is-this-bad)

The `console.log()` method and similar methods join the parameters with a space so adding a leading/trailing space to a parameter, results in two spaces being added.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
console.log("abc ", "def");
```

Examples of **correct** code for this rule:

javascript

```
console.log("abc", "def");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-console-spaces": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-console-spaces
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_console_spaces.rs)
