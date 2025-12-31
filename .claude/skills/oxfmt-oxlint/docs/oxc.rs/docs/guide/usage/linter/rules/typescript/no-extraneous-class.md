---
title: "typescript/no-extraneous-class | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-extraneous-class.md for this page in Markdown format

# typescript/no-extraneous-class Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#typescript-no-extraneous-class)

⚠️💡 A dangerous suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#what-it-does)

This rule reports when a class has no non-static members, such as for a class used exclusively as a static namespace. This rule also reports classes that have only a constructor and no fields. Those classes can generally be replaced with a standalone function.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#why-is-this-bad)

Users who come from a OOP paradigm may wrap their utility functions in an extra class, instead of putting them at the top level of an ECMAScript module. Doing so is generally unnecessary in JavaScript and TypeScript projects.

* Wrapper classes add extra cognitive complexity to code without adding any structural improvements
  + Whatever would be put on them, such as utility functions, are already organized by virtue of being in a module.
  + As an alternative, you can `import * as ...` the module to get all of them in a single object.
* IDEs can't provide as good suggestions for static class or namespace imported properties when you start typing property names
* It's more difficult to statically analyze code for unused variables, etc. when they're all on the class (see: [Finding dead code (and dead types) in TypeScript](https://effectivetypescript.com/2020/10/20/tsprune/)).

This rule also reports classes that have only a constructor and no fields. Those classes can generally be replaced with a standalone function.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#examples)

Examples of **incorrect** code for this rule:

ts

```
class StaticConstants {
  static readonly version = 42;

  static isProduction() {
    return process.env.NODE_ENV === "production";
  }
}

class HelloWorldLogger {
  constructor() {
    console.log("Hello, world!");
  }
}

abstract class Foo {}
```

Examples of **correct** code for this rule:

ts

```
const version = 42;
const isProduction = () => process.env.NODE_ENV === "production";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#configuration)

This rule accepts a configuration object with the following properties:

### allowConstructorOnly [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#allowconstructoronly)

type: `boolean`

default: `false`

Allow classes that only have a constructor.

### allowEmpty [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#allowempty)

type: `boolean`

default: `false`

Allow empty classes.

### allowStaticOnly [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#allowstaticonly)

type: `boolean`

default: `false`

Allow classes with only static members.

### allowWithDecorator [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#allowwithdecorator)

type: `boolean`

default: `false`

Allow classes with decorators.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-extraneous-class": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-extraneous-class
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_extraneous_class.rs)
