---
title: "typescript/no-misused-new | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-misused-new.md for this page in Markdown format

# typescript/no-misused-new Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#typescript-no-misused-new)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#what-it-does)

Enforces valid definition of new and constructor. This rule prevents classes from defining a method named `new` and interfaces from defining a method named `constructor`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#why-is-this-bad)

JavaScript classes may define a constructor method that runs when a class instance is newly created.

TypeScript allows interfaces that describe a static class object to define a `new()` method (though this is rarely used in real world code). Developers new to JavaScript classes and/or TypeScript interfaces may sometimes confuse when to use constructor or new.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
declare class C {
  new(): C;
}
```

typescript

```
interface I {
  new (): I;
  constructor(): void;
}
```

Examples of **correct** code for this rule:

typescript

```
declare class C {
  constructor();
}
```

typescript

```
interface I {
  new (): C;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-misused-new": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-misused-new
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_misused_new.rs)
