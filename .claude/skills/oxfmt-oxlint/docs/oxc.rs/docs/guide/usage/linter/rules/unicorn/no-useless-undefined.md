---
title: "unicorn/no-useless-undefined | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-useless-undefined.md for this page in Markdown format

# unicorn/no-useless-undefined Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#unicorn-no-useless-undefined)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#what-it-does)

Do not use useless `undefined`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#why-is-this-bad)

`undefined` is the default value for new variables, parameters, return statements, etc… so specifying it doesn't make any difference.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let foo = undefined;
```

Examples of **correct** code for this rule:

javascript

```
let foo;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#configuration)

This rule accepts a configuration object with the following properties:

### checkArguments [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#checkarguments)

type: `boolean`

default: `true`

Whether to check for useless `undefined` in function call arguments.

### checkArrowFunctionBody [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#checkarrowfunctionbody)

type: `boolean`

default: `true`

Whether to check for useless `undefined` in arrow function bodies.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-useless-undefined": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-useless-undefined
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_useless_undefined.rs)
