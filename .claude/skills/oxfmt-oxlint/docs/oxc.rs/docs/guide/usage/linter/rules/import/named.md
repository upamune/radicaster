---
title: "import/named | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/named"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/named.md for this page in Markdown format

# import/named Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#import-named)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#what-it-does)

Verifies that all named imports are part of the set of named exports in the referenced module.

For `export`, verifies that all named exports exist in the referenced module.

Note: for packages, the plugin will find exported names from `jsnext:main` (deprecated) or `module`, if present in `package.json`. Redux's npm module includes this key, and thereby is lintable, for example.

A module path that is ignored or not unambiguously an ES module will not be reported when imported. Note that type imports and exports, as used by Flow, are always ignored.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#why-is-this-bad)

Importing or exporting names that do not exist in the referenced module can lead to runtime errors and confusion. It may suggest that certain functionality is available when it is not, making the code harder to maintain and understand. This rule helps ensure that your code accurately reflects the available exports, improving reliability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#examples)

Given

js

```
// ./foo.js
export const foo = "I'm so foo";
```

Examples of **incorrect** code for this rule:

js

```
// ./baz.js
import { notFoo } from "./foo";

// re-export
export { notFoo as defNotBar } from "./foo";

// will follow 'jsnext:main', if available
import { dontCreateStore } from "redux";
```

Examples of **correct** code for this rule:

js

```
// ./bar.js
import { foo } from "./foo";

// re-export
export { foo as bar } from "./foo";

// node_modules without jsnext:main are not analyzed by default
// (import/ignore setting)
import { SomeNonsenseThatDoesntExist } from "react";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/named": "error"
  }
}
```

bash

```
oxlint --deny import/named --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/named.rs)
