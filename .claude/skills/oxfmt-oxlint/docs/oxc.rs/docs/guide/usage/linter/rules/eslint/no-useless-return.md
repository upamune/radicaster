---
title: "eslint/no-useless-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-return.md for this page in Markdown format

# eslint/no-useless-return Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#eslint-no-useless-return)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#what-it-does)

Disallows redundant return statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#why-is-this-bad)

A `return;` statement with nothing after it is redundant, and has no effect on the runtime behavior of a function. This can be confusing, so it's better to disallow these redundant statements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#examples)

Examples of **incorrect** code for this rule:

js

```
function foo() {
  return;
}

function bar() {
  doSomething();
  return;
}

function baz() {
  if (condition) {
    doSomething();
    return;
  }
}
```

Examples of **correct** code for this rule:

js

```
function foo() {
  return 5;
}

function bar() {
  if (condition) {
    return;
  }
  doSomething();
}

function baz() {
  return doSomething();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-return": "error"
  }
}
```

bash

```
oxlint --deny no-useless-return
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_return.rs)
