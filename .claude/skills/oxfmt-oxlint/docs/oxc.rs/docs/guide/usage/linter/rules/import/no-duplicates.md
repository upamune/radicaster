---
title: "import/no-duplicates | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/no-duplicates.md for this page in Markdown format

# import/no-duplicates Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#import-no-duplicates)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#what-it-does)

Reports if a resolved path is imported more than once in the same module. This helps avoid unnecessary duplicate imports and keeps the code clean.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#why-is-this-bad)

Importing the same module multiple times can lead to redundancy and unnecessary complexity. It also affects maintainability, as it might confuse developers and result in inconsistent usage of imports across the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { foo } from "./module";
import { bar } from "./module";

import a from "./module";
import { b } from "./module";
```

Examples of **correct** code for this rule:

typescript

```
import { foo, bar } from "./module";

import * as a from "foo"; // separate statements for namespace imports
import { b } from "foo";

import { c } from "foo"; // separate type imports, unless
import type { d } from "foo"; // `prefer-inline` is true
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#configuration)

This rule accepts a configuration object with the following properties:

### preferInline [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#preferinline)

type: `boolean`

default: `false`

When set to `true`, prefer inline type imports instead of separate type import statements for TypeScript code.

Examples of **correct** code with this option set to `true`:

typescript

```
import { Foo, type Bar } from "./module";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/no-duplicates": "error"
  }
}
```

bash

```
oxlint --deny import/no-duplicates --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/no_duplicates.rs)
