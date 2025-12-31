---
title: "import/max-dependencies | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/max-dependencies.md for this page in Markdown format

# import/max-dependencies Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#import-max-dependencies)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#what-it-does)

Forbid modules to have too many dependencies (import or require statements).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#why-is-this-bad)

This is a useful rule because a module with too many dependencies is a code smell, and usually indicates the module is doing too much and/or should be broken up into smaller modules.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#examples)

Given `{"max": 2}`

Examples of **incorrect** code for this rule:

javascript

```
import a from "./a";
import b from "./b";
import c from "./c"; // Too many dependencies: 3 (max: 2)
```

Examples of **correct** code for this rule:

javascript

```
import a from "./a";
import b from "./b"; // Allowed: 2 dependencies (max: 2)
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreTypeImports [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#ignoretypeimports)

type: `boolean`

default: `false`

Whether to ignore type imports when counting dependencies.

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#max)

type: `integer`

default: `10`

Maximum number of dependencies allowed in a module.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/max-dependencies": "error"
  }
}
```

bash

```
oxlint --deny import/max-dependencies --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/max_dependencies.rs)
