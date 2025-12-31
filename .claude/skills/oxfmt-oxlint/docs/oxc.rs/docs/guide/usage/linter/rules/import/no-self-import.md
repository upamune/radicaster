---
title: "import/no-self-import | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-self-import.md for this page in Markdown format

# import/no-self-import Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#import-no-self-import)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#what-it-does)

Forbids a module from importing itself. This can sometimes happen accidentally, especially during refactoring.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#why-is-this-bad)

Importing a module into itself creates a circular dependency, which can cause runtime issues, including infinite loops, unresolved imports, or `undefined` values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// foo.js
import foo from "./foo.js"; // Incorrect: module imports itself
const foo = require("./foo"); // Incorrect: module imports itself
```

Examples of **correct** code for this rule:

javascript

```
// foo.js
import bar from "./bar.js"; // Correct: module imports another module
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-self-import": "error"
  }
}
```

bash

```
oxlint --deny import/no-self-import --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_self_import.rs)
