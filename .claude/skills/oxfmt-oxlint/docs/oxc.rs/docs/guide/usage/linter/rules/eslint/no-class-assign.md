---
title: "eslint/no-class-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-class-assign.md for this page in Markdown format

# eslint/no-class-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#eslint-no-class-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#what-it-does)

Disallow reassigning class variables.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#why-is-this-bad)

`ClassDeclaration` creates a variable that can be re-assigned, but the re-assignment is a mistake in most cases.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A {}
A = 0;
```

javascript

```
A = 0;
class A {}
```

javascript

```
class A {
  b() {
    A = 0;
  }
}
```

javascript

```
let A = class A {
  b() {
    A = 0;
    // `let A` is shadowed by the class name.
  }
};
```

Examples of **correct** code for this rule:

javascript

```
let A = class A {};
A = 0; // A is a variable.
```

javascript

```
let A = class {
  b() {
    A = 0; // A is a variable.
  }
};
```

javascript

```
class A {
  b(A) {
    A = 0; // A is a parameter.
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-class-assign": "error"
  }
}
```

bash

```
oxlint --deny no-class-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_class_assign.rs)
