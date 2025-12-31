---
title: "unicorn/no-process-exit | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-process-exit.md for this page in Markdown format

# unicorn/no-process-exit Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#unicorn-no-process-exit)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#what-it-does)

Disallow `process.exit()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#why-is-this-bad)

Only use `process.exit()` in CLI apps. Throw an error instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (problem) process.exit(1);
```

Examples of **correct** code for this rule:

javascript

```
if (problem) throw new Error("message");
```

```
#!/usr/bin/env node
if (problem) process.exit(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-process-exit": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-process-exit
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_process_exit.rs)
