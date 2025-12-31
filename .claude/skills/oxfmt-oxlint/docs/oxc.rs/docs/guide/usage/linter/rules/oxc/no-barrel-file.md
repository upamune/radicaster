---
title: "oxc/no-barrel-file | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-barrel-file.md for this page in Markdown format

# oxc/no-barrel-file Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#oxc-no-barrel-file)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#what-it-does)

Disallow the use of barrel files where the file contains `export *` statements, and the total number of modules exceed a threshold.

The default threshold is 100.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#why-is-this-bad)

Barrel files that re-export many modules can significantly slow down applications and bundlers. When a barrel file exports a large number of modules, importing from it forces the runtime or bundler to process all the exported modules, even if only a few are actually used. This leads to slower startup times and larger bundle sizes.

References:

* <https://github.com/thepassle/eslint-plugin-barrel-files>
* <https://marvinh.dev/blog/speeding-up-javascript-ecosystem-part-7>

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#example)

Invalid:

javascript

```
export * from "foo"; // where `foo` loads a subtree of 100 modules
import * as ns from "foo"; // where `foo` loads a subtree of 100 modules
```

Valid:

javascript

```
export { foo } from "foo";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#configuration)

This rule accepts a configuration object with the following properties:

### threshold [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#threshold)

type: `integer`

default: `100`

The maximum number of modules that can be re-exported via `export *` before the rule is triggered.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-barrel-file": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-barrel-file
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_barrel_file.rs)
