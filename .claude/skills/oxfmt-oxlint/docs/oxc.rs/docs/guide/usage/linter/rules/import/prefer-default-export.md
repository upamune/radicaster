---
title: "import/prefer-default-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/prefer-default-export.md for this page in Markdown format

# import/prefer-default-export Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#import-prefer-default-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#what-it-does)

In exporting files, this rule checks if there is default export or not.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#why-is-this-bad)

This rule exists to standardize module exports by preferring default exports when a module only has one export, enhancing readability, maintainability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#examples)

Examples of **incorrect** code for the `{ target: "single" }` option:

js

```
export const foo = "foo";
```

Examples of **correct** code for the `{ target: "single" }` option:

js

```
export const foo = "foo";
const bar = "bar";
export default bar;
```

Examples of **incorrect** code for the `{ target: "any" }` option:

js

```
export const foo = "foo";
export const baz = "baz";
```

Examples of **correct** code for the `{ target: "any" }` option:

js

```
export default function bar() {}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#configuration)

This rule accepts a configuration object with the following properties:

### target [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#target)

type: `"single" | "any"`

default: `"single"`

Configuration option to specify the target type for preferring default exports.

* `"single"`: Prefer default export when there is only one export in the module.
* `"any"`: Prefer default export in any module that has exports.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/prefer-default-export": "error"
  }
}
```

bash

```
oxlint --deny import/prefer-default-export --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/prefer_default_export.rs)
