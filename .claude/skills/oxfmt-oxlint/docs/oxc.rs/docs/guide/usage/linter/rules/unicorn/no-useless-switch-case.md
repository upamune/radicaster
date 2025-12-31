---
title: "unicorn/no-useless-switch-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.md for this page in Markdown format

# unicorn/no-useless-switch-case Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#unicorn-no-useless-switch-case)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#what-it-does)

Disallows useless default cases in switch statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#why-is-this-bad)

An empty case before the last default case is useless.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
switch (foo) {
  case 1:
  default:
    handleDefaultCase();
    break;
}
```

Examples of **correct** code for this rule:

javascript

```
switch (foo) {
  case 1:
  case 2:
    handleCase1And2();
    break;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-switch-case": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-switch-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_switch_case.rs)
