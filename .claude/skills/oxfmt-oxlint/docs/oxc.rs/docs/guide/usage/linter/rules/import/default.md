---
title: "import/default | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/default"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/default.md for this page in Markdown format

# import/default Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#import-default)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#what-it-does)

If a default import is requested, this rule will report if there is no default export in the imported module.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#why-is-this-bad)

Using a default import when there is no default export can lead to confusion and runtime errors. It can make the code harder to understand and maintain, as it may suggest that a module has a default export when it does not, leading to unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// ./bar.js
export function bar() {
  return null;
}

// ./foo.js
import bar from "./bar"; // no default export found in ./bar
```

Examples of **correct** code for this rule:

javascript

```
// ./bar.js
export default function bar() {
  return null;
}

// ./foo.js
import { bar } from "./bar"; // correct usage of named import
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/default": "error"
  }
}
```

bash

```
oxlint --deny import/default --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/default.rs)
