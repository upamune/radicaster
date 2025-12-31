---
title: "typescript/no-unnecessary-parameter-property-assignment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.md for this page in Markdown format

# typescript/no-unnecessary-parameter-property-assignment Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#typescript-no-unnecessary-parameter-property-assignment)

✅ This rule is turned on by default.

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#what-it-does)

Prevents unnecessary assignment of parameter properties.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#why-is-this-bad)

Constructor parameters marked with one of the visibility modifiers public, private, protected, or readonly are automatically initialized. Providing an explicit assignment is unnecessary and can be removed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#examples)

Examples of **incorrect** code for this rule:

js

```
class Foo {
  constructor(public name: unknown) {
    this.name = name;
  }
}
```

Examples of **correct** code for this rule:

js

```
class Foo {
  constructor(public name: unknown) {}
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-unnecessary-parameter-property-assignment": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-unnecessary-parameter-property-assignment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_unnecessary_parameter_property_assignment.rs)
