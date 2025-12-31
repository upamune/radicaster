---
title: "eslint/no-import-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-import-assign.md for this page in Markdown format

# eslint/no-import-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#eslint-no-import-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#what-it-does)

Disallow assigning to imported bindings.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#why-is-this-bad)

The updates of imported bindings by ES Modules cause runtime errors.

The TypeScript compiler generally enforces this check already. Although it should be noted that there are some cases TypeScript does not catch, such as assignments via `Object.assign`. So this rule is still useful for TypeScript code in those cases.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
import mod, { named } from "./mod.mjs";
import * as mod_ns from "./mod.mjs";

mod = 1; // ERROR: 'mod' is readonly.
named = 2; // ERROR: 'named' is readonly.
mod_ns.named = 3; // ERROR: The members of 'mod_ns' are readonly.
mod_ns = {}; // ERROR: 'mod_ns' is readonly.
// Can't extend 'mod_ns'
Object.assign(mod_ns, { foo: "foo" }); // ERROR: The members of 'mod_ns' are readonly.
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-import-assign": "error"
  }
}
```

bash

```
oxlint --deny no-import-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_import_assign.rs)
