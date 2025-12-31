---
title: "typescript/related-getter-setter-pairs | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.md for this page in Markdown format

# typescript/related-getter-setter-pairs Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#typescript-related-getter-setter-pairs)

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#what-it-does)

This rule enforces that getters and setters for the same property are defined together and have related types.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#why-is-this-bad)

When you define a getter and setter for the same property, they should typically be defined together and work with compatible types. Having mismatched types or defining them separately can lead to confusion and potential runtime errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#examples)

Examples of **incorrect** code for this rule:

ts

```
class Example {
  // Getter and setter with incompatible types
  get value(): string {
    return this._value.toString();
  }

  set value(val: number) {
    // Incompatible with getter
    this._value = val;
  }

  private _value: number = 0;
}

// Getter without corresponding setter or vice versa might be flagged
class IncompleteProperty {
  get readOnlyValue(): string {
    return "constant";
  }
  // Missing setter - might be intended, but should be consistent
}
```

Examples of **correct** code for this rule:

ts

```
class Example {
  // Getter and setter with compatible types
  get value(): string {
    return this._value;
  }

  set value(val: string) {
    this._value = val;
  }

  private _value: string = "";
}

// Read-only property with only getter
class ReadOnlyProperty {
  get constant(): string {
    return "constant value";
  }
}

// Write-only property with only setter (less common but valid)
class WriteOnlyProperty {
  set logger(message: string) {
    console.log(message);
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/related-getter-setter-pairs": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/related-getter-setter-pairs
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/related_getter_setter_pairs.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/related_getter_setter_pairs/related_getter_setter_pairs.go)
