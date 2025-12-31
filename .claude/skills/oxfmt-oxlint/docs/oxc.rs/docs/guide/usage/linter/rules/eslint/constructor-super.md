---
title: "eslint/constructor-super | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/constructor-super.md for this page in Markdown format

# eslint/constructor-super Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#eslint-constructor-super)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#what-it-does)

Requires `super()` calls in constructors of derived classes and disallows `super()` calls in constructors of non-derived classes.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#why-is-this-bad)

In JavaScript, calling `super()` in the constructor of a derived class (a class that extends another class) is required. Failing to do so will result in a ReferenceError at runtime. Conversely, calling `super()` in a non-derived class is a syntax error.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#examples)

Examples of **incorrect** code for this rule:

js

```
// Missing super() call
class A extends B {
    constructor() { }
}

// super() in non-derived class
class A {
    constructor() {
        super();
    }
}

// super() only in some code paths
class C extends D {
    constructor() {
        if (condition) {
            super();
        }
    }
}
```

Examples of **correct** code for this rule:

js

```
// Proper super() call in derived class
class A extends B {
  constructor() {
    super();
  }
}

// No super() in non-derived class
class A {
  constructor() {}
}

// super() in all code paths
class C extends D {
  constructor() {
    if (condition) {
      super();
    } else {
      super();
    }
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "constructor-super": "error"
  }
}
```

bash

```
oxlint --deny constructor-super
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/constructor_super.rs)
