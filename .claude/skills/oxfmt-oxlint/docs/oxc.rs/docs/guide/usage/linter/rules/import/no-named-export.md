---
title: "import/no-named-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-named-export.md for this page in Markdown format

# import/no-named-export Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#import-no-named-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#what-it-does)

Prohibit named exports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#why-is-this-bad)

Named exports require strict identifier matching and can lead to fragile imports, while default exports enforce a single, consistent module entry point.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#examples)

Examples of **incorrect** code for this rule:

js

```
export const foo = "foo";

const bar = "bar";
export { bar };
```

Examples of **correct** code for this rule:

js

```
export default 'bar';

const foo = 'foo';
export { foo as default }
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-named-export": "error"
  }
}
```

bash

```
oxlint --deny import/no-named-export --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_named_export.rs)
