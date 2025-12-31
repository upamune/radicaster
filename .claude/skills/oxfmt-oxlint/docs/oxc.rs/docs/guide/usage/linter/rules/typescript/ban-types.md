---
title: "typescript/ban-types | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/ban-types.md for this page in Markdown format

# typescript/ban-types Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#typescript-ban-types)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#what-it-does)

This rule bans specific types and can suggest alternatives. Note that it does not ban the corresponding runtime objects from being used.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#why-is-this-bad)

Some built-in types have aliases, while some types are considered dangerous or harmful. It's often a good idea to ban certain types to help with consistency and safety.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
let foo: String = "foo";

let bar: Boolean = true;
```

Examples of **correct** code for this rule:

typescript

```
let foo: string = "foo";

let bar: boolean = true;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/ban-types": "error"
  }
}
```

bash

```
oxlint --deny typescript/ban-types
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/ban_types.rs)
