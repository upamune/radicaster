---
title: "eslint/no-obj-calls | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-obj-calls.md for this page in Markdown format

# eslint/no-obj-calls Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#eslint-no-obj-calls)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#what-it-does)

Disallow calling some global objects as functions.

This rule can be disabled for TypeScript code, as the TypeScript compiler enforces this check.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#why-is-this-bad)

Some global objects are not intended to be called as functions. Calling them as functions will usually result in a TypeError being thrown.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let math = Math();
let newMath = new Math();

let json = JSON();
let newJson = new JSON();

let atomics = Atomics();
let newAtomics = new Atomics();

let intl = Intl();
let newIntl = new Intl();

let reflect = Reflect();
let newReflect = new Reflect();
```

Examples of **correct** code for this rule:

javascript

```
let area = (r) => 2 * Math.PI * r * r;
let object = JSON.parse("{}");
let first = Atomics.load(sharedArray, 0);
let segmenterFrom = Intl.Segmenter("fr", { granularity: "word" });
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-obj-calls": "error"
  }
}
```

bash

```
oxlint --deny no-obj-calls
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_obj_calls.rs)
