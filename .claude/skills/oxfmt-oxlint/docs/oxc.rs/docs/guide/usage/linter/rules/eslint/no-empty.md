---
title: "eslint/no-empty | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-empty.md for this page in Markdown format

# eslint/no-empty Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#eslint-no-empty)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#what-it-does)

Disallows empty block statements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#why-is-this-bad)

Empty block statements, while not technically errors, usually occur due to refactoring that wasn’t completed. They can cause confusion when reading code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (condition) {
}
```

Examples of **correct** code for this rule:

javascript

```
if (condition) {
  throw new Error("condition should be false");
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#configuration)

This rule accepts a configuration object with the following properties:

### allowEmptyCatch [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#allowemptycatch)

type: `boolean`

default: `false`

If set to `true`, allows an empty `catch` block without triggering the linter.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-empty": "error"
  }
}
```

bash

```
oxlint --deny no-empty
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_empty.rs)
