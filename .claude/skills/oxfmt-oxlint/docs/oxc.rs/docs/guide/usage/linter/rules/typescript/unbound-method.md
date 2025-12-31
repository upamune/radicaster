---
title: "typescript/unbound-method | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/unbound-method.md for this page in Markdown format

# typescript/unbound-method Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#typescript-unbound-method)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#what-it-does)

This rule enforces unbound methods are called with their expected scope.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#why-is-this-bad)

When you extract a method from an object and call it separately, the `this` context is lost. This can lead to runtime errors or unexpected behavior, especially with methods that rely on `this` to access instance properties or other methods.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#examples)

Examples of **incorrect** code for this rule:

ts

```
class MyClass {
  private value = 42;

  getValue() {
    return this.value;
  }

  processValue() {
    return this.value * 2;
  }
}

const instance = new MyClass();

// Unbound method - loses 'this' context
const getValue = instance.getValue;
getValue(); // Runtime error: cannot read property 'value' of undefined

// Passing unbound method as callback
[1, 2, 3].map(instance.processValue); // 'this' will be undefined

// Destructuring methods
const { getValue: unboundGetValue } = instance;
unboundGetValue(); // Runtime error
```

Examples of **correct** code for this rule:

ts

```
class MyClass {
  private value = 42;

  getValue() {
    return this.value;
  }

  processValue() {
    return this.value * 2;
  }
}

const instance = new MyClass();

// Call method on instance
const value = instance.getValue(); // Correct

// Bind method to preserve context
const boundGetValue = instance.getValue.bind(instance);
boundGetValue(); // Correct

// Use arrow function to preserve context
[1, 2, 3].map(() => instance.processValue()); // Correct

// Use arrow function in class for auto-binding
class MyClassWithArrow {
  private value = 42;

  getValue = () => {
    return this.value;
  };
}

const instance2 = new MyClassWithArrow();
const getValue = instance2.getValue; // Safe - arrow function preserves 'this'
getValue(); // Correct
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreStatic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#ignorestatic)

type: `boolean`

default: `false`

Whether to ignore unbound methods that are static. When true, static methods can be referenced without binding.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/unbound-method": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/unbound-method
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/unbound_method.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/unbound_method/unbound_method.go)
