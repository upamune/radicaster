---
title: "import/no-cycle | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-cycle.md for this page in Markdown format

# import/no-cycle Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#import-no-cycle)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#what-it-does)

Ensures that there is no resolvable path back to this module via its dependencies.

This includes cycles of depth 1 (imported module imports me) to "∞" (or Infinity), if the `maxDepth` option is not set.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#why-is-this-bad)

Dependency cycles lead to confusing architectures where bugs become hard to find. It is common to import an `undefined` value that is caused by a cyclic dependency.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// dep-b.js
import "./dep-a.js";
export function b() {
  /* ... */
}
```

javascript

```
// dep-a.js
import { b } from "./dep-b.js"; // reported: Dependency cycle detected.
export function a() {
  /* ... */
}
```

In this example, `dep-a.js` and `dep-b.js` import each other, creating a circular dependency, which is problematic.

Examples of **correct** code for this rule:

javascript

```
// dep-b.js
export function b() {
  /* ... */
}
```

javascript

```
// dep-a.js
import { b } from "./dep-b.js"; // no circular dependency
export function a() {
  /* ... */
}
```

In this corrected version, `dep-b.js` no longer imports `dep-a.js`, breaking the cycle.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#configuration)

This rule accepts a configuration object with the following properties:

### allowUnsafeDynamicCyclicDependency [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#allowunsafedynamiccyclicdependency)

type: `boolean`

default: `false`

Allow cyclic dependency if there is at least one dynamic import in the chain

### ignoreExternal [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#ignoreexternal)

type: `boolean`

default: `false`

Ignore external modules

### ignoreTypes [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#ignoretypes)

type: `boolean`

default: `true`

Ignore type-only imports

### maxDepth [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#maxdepth)

type: `integer`

default: `4294967295`

Maximum dependency depth to traverse

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-cycle": "error"
  }
}
```

bash

```
oxlint --deny import/no-cycle --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_cycle.rs)
