---
title: "import/consistent-type-specifier-style | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/import/consistent-type-specifier-style.md for this page in Markdown format

# import/consistent-type-specifier-style Style [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#import-consistent-type-specifier-style)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#what-it-does)

This rule either enforces or bans the use of inline type-only markers for named imports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#why-is-this-bad)

Mixing top-level `import type { Foo } from 'foo'` with inline `{ type Bar }` forces readers to mentally switch contexts when scanning your imports. Enforcing one style makes it immediately obvious which imports are types and which are value imports.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#examples)

Examples of incorrect code for the default `prefer-top-level` option:

typescript

```
import { type Foo } from "Foo";
import Foo, { type Bar } from "Foo";
```

Examples of correct code for the default option:

typescript

```
import type { Foo } from "Foo";
import type Foo, { Bar } from "Foo";
```

Examples of incorrect code for the `prefer-inline` option:

typescript

```
import type { Foo } from "Foo";
import type Foo, { Bar } from "Foo";
```

Examples of correct code for the `prefer-inline` option:

typescript

```
import { type Foo } from "Foo";
import Foo, { type Bar } from "Foo";
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#configuration)

This rule accepts one of the following string values:

### `"prefer-top-level"` [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#prefer-top-level)

Prefer `import type { Foo } from 'foo'` for type imports.

### `"prefer-inline"` [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#prefer-inline)

Prefer `import { type Foo } from 'foo'` for type imports.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["import"],
  "rules": {
    "import/consistent-type-specifier-style": "error"
  }
}
```

bash

```
oxlint --deny import/consistent-type-specifier-style --import-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/import/consistent_type_specifier_style.rs)
