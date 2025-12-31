---
title: "import/first | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/first"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/first.md for this page in Markdown format

# import/first Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#import-first)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#what-it-does)

Forbids any non-import statements before imports except directives.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#why-is-this-bad)

Notably, imports are hoisted, which means the imported modules will be evaluated before any of the statements interspersed between them. Keeping all imports together at the top of the file may prevent surprises resulting from this part of the spec

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#examples)

Examples of **incorrect** code for this rule:

js

```
import { x } from "./foo";
export { x };
import { y } from "./bar";
```

Examples of **correct** code for this rule:

js

```
import { x } from "./foo";
import { y } from "./bar";
export { x, y };
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#configuration)

This rule accepts one of the following string values:

### `"absolute-first"` [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#absolute-first)

Forces absolute imports to be listed before relative imports.

Examples of **incorrect** code for this rule with `"absolute-first"`:

js

```
import { x } from "./foo";
import { y } from "bar";
```

Examples of **correct** code for this rule with `"absolute-first"`:

js

```
import { y } from "bar";
import { x } from "./foo";
```

### `"disable-absolute-first"` [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#disable-absolute-first)

Disables the absolute-first behavior. This is the default behavior.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/first": "error"
  }
}
```

bash

```
oxlint --deny import/first --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/first.rs)
