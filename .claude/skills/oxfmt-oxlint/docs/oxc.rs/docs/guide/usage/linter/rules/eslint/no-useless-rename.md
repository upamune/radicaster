---
title: "eslint/no-useless-rename | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-rename.md for this page in Markdown format

# eslint/no-useless-rename Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#eslint-no-useless-rename)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#what-it-does)

Disallow renaming import, export, and destructured assignments to the same name.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#why-is-this-bad)

It is unnecessary to rename a variable to the same name.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import { foo as foo } from "foo";
const { bar: bar } = obj;
export { baz as baz };
```

Examples of **correct** code for this rule:

javascript

```
import { foo } from "foo";
const { bar: renamed } = obj;
export { baz };
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreDestructuring [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#ignoredestructuring)

type: `boolean`

default: `false`

When set to `true`, allows using the same name in destructurings.

### ignoreExport [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#ignoreexport)

type: `boolean`

default: `false`

When set to `true`, allows renaming exports to the same name.

### ignoreImport [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#ignoreimport)

type: `boolean`

default: `false`

When set to `true`, allows renaming imports to the same name.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-rename": "error"
  }
}
```

bash

```
oxlint --deny no-useless-rename
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_rename.rs)
