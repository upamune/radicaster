---
title: "typescript/prefer-return-this-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/prefer-return-this-type.md for this page in Markdown format

# typescript/prefer-return-this-type Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#typescript-prefer-return-this-type)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#what-it-does)

This rule enforces using `this` types for return types when possible.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#why-is-this-bad)

Classes that have methods which return the instance itself should use `this` as the return type instead of the class name. This provides better type safety for inheritance, as the return type will be the actual subclass type rather than the base class type.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#examples)

Examples of **incorrect** code for this rule:

ts

```
class Builder {
  private value: string = "";

  setValue(value: string): Builder {
    // Should return 'this'
    this.value = value;
    return this;
  }

  build(): string {
    return this.value;
  }
}

class FluentAPI {
  method1(): FluentAPI {
    // Should return 'this'
    return this;
  }

  method2(): FluentAPI {
    // Should return 'this'
    return this;
  }
}
```

Examples of **correct** code for this rule:

ts

```
class Builder {
  private value: string = "";

  setValue(value: string): this {
    this.value = value;
    return this;
  }

  build(): string {
    return this.value;
  }
}

class FluentAPI {
  method1(): this {
    return this;
  }

  method2(): this {
    return this;
  }
}

// Now inheritance works correctly
class ExtendedBuilder extends Builder {
  setPrefix(prefix: string): this {
    // The return type is 'this' (ExtendedBuilder), not Builder
    return this.setValue(prefix + this.getValue());
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/prefer-return-this-type": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/prefer-return-this-type
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/prefer_return_this_type.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/prefer_return_this_type/prefer_return_this_type.go)
