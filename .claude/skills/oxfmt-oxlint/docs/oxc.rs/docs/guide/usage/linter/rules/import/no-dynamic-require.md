---
title: "import/no-dynamic-require | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-dynamic-require.md for this page in Markdown format

# import/no-dynamic-require Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#import-no-dynamic-require)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#what-it-does)

Forbids imports that use an expression for the module argument. This includes dynamically resolved paths in `require` or `import` statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#why-is-this-bad)

Using expressions that are resolved at runtime in import statements makes it difficult to determine where the module is being imported from. This can complicate code navigation and hinder static analysis tools, which rely on predictable module paths for linting, bundling, and other optimizations.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
require(name);
require(`../${name}`);
```

Examples of **correct** code for this rule:

javascript

```
require("../name");
require(`../name`);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#configuration)

This rule accepts a configuration object with the following properties:

### esmodule [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#esmodule)

type: `boolean`

default: `false`

When `true`, also check `import()` expressions for dynamic module specifiers.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-dynamic-require": "error"
  }
}
```

bash

```
oxlint --deny import/no-dynamic-require --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_dynamic_require.rs)
