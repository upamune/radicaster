---
title: "import/no-default-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-default-export.md for this page in Markdown format

# import/no-default-export Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#import-no-default-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#what-it-does)

Forbids a module from having default exports. This helps your editor provide better auto-import functionality, as named exports offer more explicit and predictable imports compared to default exports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#why-is-this-bad)

Default exports can lead to confusion, as the name of the imported value can vary based on how it's imported. This can make refactoring and auto-imports less reliable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export default 'bar';

const foo = 'foo';
export { foo as default }
```

Examples of **correct** code for this rule:

javascript

```
export const foo = "foo";
export const bar = "bar";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-default-export": "error"
  }
}
```

bash

```
oxlint --deny import/no-default-export --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_default_export.rs)
