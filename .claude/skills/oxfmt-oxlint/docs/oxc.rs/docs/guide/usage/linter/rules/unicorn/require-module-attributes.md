---
title: "unicorn/require-module-attributes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/require-module-attributes.md for this page in Markdown format

# unicorn/require-module-attributes Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#unicorn-require-module-attributes)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#what-it-does)

This rule enforces non-empty attribute list in import/export statements and import() expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#why-is-this-bad)

Import attributes are meant to provide metadata about how a module should be loaded (e.g., `with { type: "json" }`). An empty attribute object provides no information and should be removed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#examples)

Examples of **incorrect** code for this rule:

js

```
import foo from "foo" with {};

export { foo } from "foo" with {};

const foo = await import("foo", {});

const foo = await import("foo", { with: {} });
```

Examples of **correct** code for this rule:

js

```
import foo from "foo";

export { foo } from "foo";

const foo = await import("foo");

const foo = await import("foo");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/require-module-attributes": "error"
  }
}
```

bash

```
oxlint --deny unicorn/require-module-attributes
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/require_module_attributes.rs)
