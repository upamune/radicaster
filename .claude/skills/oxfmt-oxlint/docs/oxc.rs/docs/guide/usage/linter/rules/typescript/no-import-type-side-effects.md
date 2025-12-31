---
title: "typescript/no-import-type-side-effects | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.md for this page in Markdown format

# typescript/no-import-type-side-effects Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#typescript-no-import-type-side-effects)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#what-it-does)

Enforce the use of top-level `import type` qualifier when an import only has specifiers with inline type qualifiers.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#why-is-this-bad)

The `--verbatimModuleSyntax` compiler option causes TypeScript to do simple and predictable transpilation on import declarations. Namely, it completely removes import declarations with a top-level type qualifier, and it removes any import specifiers with an inline type qualifier.

The latter behavior does have one potentially surprising effect in that in certain cases TS can leave behind a "side effect" import at runtime:

ts

```
import { type A, type B } from "mod";
```

is transpiled to

ts

```
import {} from "mod";
// which is the same as
import "mod";
```

For the rare case of needing to import for side effects, this may be desirable - but for most cases you will not want to leave behind an unnecessary side effect import.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#examples)

Examples of **incorrect** code for this rule:

ts

```
import { type A } from "mod";
import { type A as AA } from "mod";
import { type A, type B } from "mod";
import { type A as AA, type B as BB } from "mod";
```

Examples of **correct** code for this rule:

ts

```
import type { A } from "mod";
import type { A as AA } from "mod";
import type { A, B } from "mod";
import type { A as AA, B as BB } from "mod";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-import-type-side-effects": "error"
  }
}
```

bash

```
oxlint --deny typescript/no-import-type-side-effects
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_import_type_side_effects.rs)
