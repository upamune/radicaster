---
title: "eslint/no-useless-constructor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-constructor.md for this page in Markdown format

# eslint/no-useless-constructor Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#eslint-no-useless-constructor)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#what-it-does)

Disallow constructors that can be safely removed without changing how the class works.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#why-is-this-bad)

ES2015 provides a default class constructor if one is not specified. As such, it is unnecessary to provide an empty constructor or one that simply delegates into its parent class.

WARNING

Caveat: This lint rule will report on constructors whose sole purpose is to change visibility of a parent constructor. This is because the rule does not have type information to determine if the parent constructor is public, protected, or private.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A {
  constructor() {}
}

class B extends A {
  constructor(...args) {
    super(...args);
  }
}
```

Examples of **correct** code for this rule:

javascript

```
class A {}

class B {
  constructor() {
    doSomething();
  }
}

class C extends A {
  constructor() {
    super("foo");
  }
}

class D extends A {
  constructor() {
    super();
    doSomething();
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-constructor": "error"
  }
}
```

bash

```
oxlint --deny no-useless-constructor
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_constructor.rs)
