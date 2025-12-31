---
title: "unicorn/prefer-class-fields | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-class-fields.md for this page in Markdown format

# unicorn/prefer-class-fields Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#unicorn-prefer-class-fields)

🛠️💡 An auto-fix and a suggestion are available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#what-it-does)

Prefers class field declarations over `this` assignments in constructors for static values.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#why-is-this-bad)

Class field declarations are more readable and less error-prone than assigning static values to `this` in the constructor. Using class fields keeps the constructor cleaner and makes the intent clearer.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#examples)

Examples of **incorrect** code for this rule:

js

```
class Foo {
  constructor() {
    this.bar = 1;
  }
}
```

Examples of **correct** code for this rule:

js

```
class Foo {
  bar = 1;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-class-fields": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-class-fields
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_class_fields.rs)
