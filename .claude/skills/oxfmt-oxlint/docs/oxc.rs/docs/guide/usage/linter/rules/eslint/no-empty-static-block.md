---
title: "eslint/no-empty-static-block | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-empty-static-block.md for this page in Markdown format

# eslint/no-empty-static-block Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#eslint-no-empty-static-block)

✅ This rule is turned on by default.

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#what-it-does)

Disallows the usages of empty static blocks

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#why-is-this-bad)

Empty block statements, while not technically errors, usually occur due to refactoring that wasn’t completed. They can cause confusion when reading code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#examples)

Examples of **incorrect** code for this rule:

js

```
class Foo {
  static {}
}
```

Examples of **correct** code for this rule:

js

```
class Foo {
  static {
    // blocks with comments are allowed
  }
}
class Bar {
  static {
    doSomething();
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-empty-static-block": "error"
  }
}
```

bash

```
oxlint --deny no-empty-static-block
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_empty_static_block.rs)
