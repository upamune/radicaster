---
title: "unicorn/no-lonely-if | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-lonely-if.md for this page in Markdown format

# unicorn/no-lonely-if Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#unicorn-no-lonely-if)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#what-it-does)

Disallow `if` statements as the only statement in `if` blocks without `else`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#why-is-this-bad)

It can be confusing to have an `if` statement without an `else` clause as the only statement in an `if` block.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (foo) {
  if (bar) {
  }
}
if (foo) if (bar) baz();
```

Examples of **correct** code for this rule:

javascript

```
if (foo && bar) {
}
if (foo && bar) baz();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-lonely-if": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-lonely-if
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_lonely_if.rs)
