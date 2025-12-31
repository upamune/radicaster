---
title: "import/no-amd | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-amd.md for this page in Markdown format

# import/no-amd Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#import-no-amd)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#what-it-does)

Forbids the use of AMD `require` and `define` calls.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#why-is-this-bad)

AMD (Asynchronous Module Definition) is an older module format that is less common in modern JavaScript development, especially with the widespread use of ES modules and CommonJS in Node.js. AMD introduces unnecessary complexity and is often considered outdated. This rule enforces the use of more modern module systems to improve maintainability and consistency across the codebase.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
require([a, b], function () {});
```

Examples of **correct** code for this rule:

javascript

```
require("../name");
require(`../name`);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-amd": "error"
  }
}
```

bash

```
oxlint --deny import/no-amd --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_amd.rs)
