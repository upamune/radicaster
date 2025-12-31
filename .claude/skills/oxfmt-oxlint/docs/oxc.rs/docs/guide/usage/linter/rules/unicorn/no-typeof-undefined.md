---
title: "unicorn/no-typeof-undefined | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.md for this page in Markdown format

# unicorn/no-typeof-undefined Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#unicorn-no-typeof-undefined)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#what-it-does)

Disallow `typeof` comparisons with `undefined`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#why-is-this-bad)

Checking if a value is `undefined` by using `typeof value === 'undefined'` is needlessly verbose. It's generally better to compare against `undefined` directly. The only time `typeof` is needed is when a global variable potentially does not exists, in which case, using `globalThis.value === undefined` may be better.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
typeof foo === "undefined";
```

Examples of **correct** code for this rule:

javascript

```
foo === undefined;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#configuration)

This rule accepts a configuration object with the following properties:

### checkGlobalVariables [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#checkglobalvariables)

type: `boolean`

default: `false`

If set to `true`, also report `typeof x === "undefined"` when `x` may be a global variable that is not declared (commonly checked via `typeof foo === "undefined"`).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-typeof-undefined": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-typeof-undefined
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_typeof_undefined.rs)
