---
title: "eslint/class-methods-use-this | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/class-methods-use-this.md for this page in Markdown format

# eslint/class-methods-use-this Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#eslint-class-methods-use-this)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#what-it-does)

Enforce that class methods utilize this.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#examples)

Examples of **incorrect** code for this rule:

js

```
class A {
  foo() {
    console.log("Hello World");
  }
}
```

Examples of **correct** code for this rule:

js

```
class A {
  foo() {
    this.bar = "Hello World"; // OK, this is used
  }
}

class B {
  constructor() {
    // OK. constructor is exempt
  }
}

class C {
  static foo() {
    // OK. static methods aren't expected to use this.
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForClassFields [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#enforceforclassfields)

type: `boolean`

default: `true`

Enforce this rule for class fields that are functions.

### exceptMethods [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#exceptmethods)

type: `array`

default: `[]`

List of method names to exempt from this rule.

#### exceptMethods[n] [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#exceptmethods-n)

type: `object`

##### exceptMethods[n].name [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#exceptmethods-n-name)

type: `string`

##### exceptMethods[n].private [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#exceptmethods-n-private)

type: `boolean`

### ignoreClassesWithImplements [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#ignoreclasseswithimplements)

type: `"all" | "public-fields"`

default: `null`

Whether to ignore classes that implement interfaces.

### ignoreOverrideMethods [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#ignoreoverridemethods)

type: `boolean`

default: `false`

Whether to ignore methods that are overridden.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "class-methods-use-this": "error"
  }
}
```

bash

```
oxlint --deny class-methods-use-this
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/class_methods_use_this.rs)
