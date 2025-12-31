---
title: "eslint/no-dupe-class-members | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-dupe-class-members.md for this page in Markdown format

# eslint/no-dupe-class-members Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#eslint-no-dupe-class-members)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#what-it-does)

Disallow duplicate class members.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#why-is-this-bad)

If there are declarations of the same name in class members, the last declaration overwrites other declarations silently. It can cause unexpected behaviors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A {
  foo() {
    console.log("foo");
  }
  foo = 123;
}
let a = new A();
a.foo(); // Uncaught TypeError: a.foo is not a function
```

Examples of **correct** code for this rule:

javascript

```
class A {
  foo() {
    console.log("foo");
  }
}
let a = new A();
a.foo();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-dupe-class-members": "error"
  }
}
```

bash

```
oxlint --deny no-dupe-class-members
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_dupe_class_members.rs)
