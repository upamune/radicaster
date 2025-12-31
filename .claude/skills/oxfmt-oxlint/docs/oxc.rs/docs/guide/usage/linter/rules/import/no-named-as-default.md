---
title: "import/no-named-as-default | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-named-as-default.md for this page in Markdown format

# import/no-named-as-default Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#import-no-named-as-default)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#what-it-does)

Reports use of an exported name as the locally imported name of a default export. This happens when an imported default export is assigned a name that conflicts with a named export from the same module.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#why-is-this-bad)

Using a named export's identifier for a default export can cause confusion and errors in understanding which value is being imported. It also reduces code clarity, making it harder for other developers to understand the intended imports.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#examples)

Given

javascript

```
// foo.js
export default "foo";
export const bar = "baz";
```

Examples of **incorrect** code for this rule:

javascript

```
// Invalid: using exported name 'bar' as the identifier for default export.
import bar from "./foo.js";
```

Examples of **correct** code for this rule:

javascript

```
// Valid: correctly importing default export with a non-conflicting name.
import foo from "./foo.js";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-named-as-default": "error"
  }
}
```

bash

```
oxlint --deny import/no-named-as-default --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_named_as_default.rs)
