---
title: "eslint/no-global-assign | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-global-assign.md for this page in Markdown format

# eslint/no-global-assign Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#eslint-no-global-assign)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#what-it-does)

Disallow modifications to read-only global variables.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#why-is-this-bad)

In almost all cases, you don't want to assign a value to these global variables as doing so could result in losing access to important functionality.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Object = null;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#configuration)

This rule accepts a configuration object with the following properties:

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#exceptions)

type: `string[]`

default: `[]`

List of global variable names to exclude from this rule. Globals listed here can be assigned to without triggering warnings.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-global-assign": "error"
  }
}
```

bash

```
oxlint --deny no-global-assign
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_global_assign.rs)
