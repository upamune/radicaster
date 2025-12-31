---
title: "import/no-mutable-exports | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-mutable-exports.md for this page in Markdown format

# import/no-mutable-exports Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#import-no-mutable-exports)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#what-it-does)

Forbids the use of mutable exports with var or let.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#why-is-this-bad)

In general, we should always export constants

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#examples)

Examples of **incorrect** code for this rule:

js

```
export let count = 2;
export var count = 3;

let count = 4;
export { count };
```

Examples of **correct** code for this rule:

js

```
export const count = 1;
export function getCount() {}
export class Counter {}
```

### Functions/Classes [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#functions-classes)

Note that exported function/class declaration identifiers may be reassigned, but are not flagged by this rule at this time. They may be in the future.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-mutable-exports": "error"
  }
}
```

bash

```
oxlint --deny import/no-mutable-exports --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_mutable_exports.rs)
