---
title: "eslint/no-unused-private-class-members | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.md for this page in Markdown format

# eslint/no-unused-private-class-members Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#eslint-no-unused-private-class-members)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#what-it-does)

Disallow unused private class members

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#why-is-this-bad)

Private class members that are declared and not used anywhere in the code are most likely an error due to incomplete refactoring. Such class members take up space in the code and can lead to confusion by readers.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class A {
  #unusedMember = 5;
}

class B {
  #usedOnlyInWrite = 5;
  method() {
    this.#usedOnlyInWrite = 42;
  }
}

class C {
  #usedOnlyToUpdateItself = 5;
  method() {
    this.#usedOnlyToUpdateItself++;
  }
}

class D {
  #unusedMethod() {}
}

class E {
  get #unusedAccessor() {}
  set #unusedAccessor(value) {}
}
```

Examples of **correct** code for this rule:

javascript

```
class A {
  #usedMember = 42;
  method() {
    return this.#usedMember;
  }
}

class B {
  #usedMethod() {
    return 42;
  }
  anotherMethod() {
    return this.#usedMethod();
  }
}

class C {
  get #usedAccessor() {}
  set #usedAccessor(value) {}

  method() {
    this.#usedAccessor = 42;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-unused-private-class-members": "error"
  }
}
```

bash

```
oxlint --deny no-unused-private-class-members
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_unused_private_class_members.rs)
