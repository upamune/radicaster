---
title: "unicorn/require-module-specifiers | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/require-module-specifiers.md for this page in Markdown format

# unicorn/require-module-specifiers Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#unicorn-require-module-specifiers)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#what-it-does)

Enforce non-empty specifier list in `import` and `export` statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#why-is-this-bad)

Empty import/export specifiers add no value and can be confusing. If you want to import a module for side effects, use `import 'module'` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#examples)

Examples of **incorrect** code for this rule:

js

```
import {} from "foo";
import foo from "foo";
export {} from "foo";
export {};
```

Examples of **correct** code for this rule:

js

```
import "foo";
import foo from "foo";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/require-module-specifiers": "error"
  }
}
```

bash

```
oxlint --deny unicorn/require-module-specifiers
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/require_module_specifiers.rs)
