---
title: "typescript/no-restricted-types | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-restricted-types.md for this page in Markdown format

# typescript/no-restricted-types Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#typescript-no-restricted-types)

🛠️💡 An auto-fix and a suggestion are available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#what-it-does)

Disallow certain types from being used.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#why-is-this-bad)

Some built-in types have aliases, while some types are considered dangerous or harmful. It's often a good idea to ban certain types to help with consistency and safety.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#examples)

Given `{ "types": { "Foo": { "message": "Use Bar instead", "fixWith": "Bar" } } }`:

Examples of **incorrect** code for this rule:

ts

```
let value: Foo;
```

Examples of **correct** code for this rule:

ts

```
let value: Bar;
```

Other examples of configuration option setups for this rule:

* Banning the `Foo` type with just a message, no fixes or suggestions: `{ "types": { "Foo": "Use` OtherType `instead." } }`
* Banning `Bar` type with suggestion: `{ "types": { "Bar": { "message": "Avoid using` Bar`.", "suggest": "BazQux" } } }`
* Banning `Object` type with a generic message: `{ "types": { "Object": true } }`

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#configuration)

This rule accepts a configuration object with the following properties:

### types [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#types)

type: `object`

default: `{}`

A mapping of type names to ban configurations.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-restricted-types": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-restricted-types
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_restricted_types.rs)
