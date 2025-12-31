---
title: "unicorn/no-static-only-class | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-static-only-class.md for this page in Markdown format

# unicorn/no-static-only-class Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#unicorn-no-static-only-class)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#what-it-does)

Disallow classes that only have static members.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#why-is-this-bad)

A class with only static members could just be an object instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A {
  static a() {}
}
```

Examples of **correct** code for this rule:

javascript

```
class A {
  static a() {}

  constructor() {}
}
```

javascript

```
const X = {
  foo: false,
  bar() {},
};
```

javascript

```
class X {
  static #foo = false; // private field
  static bar() {}
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-static-only-class": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-static-only-class
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_static_only_class.rs)
