---
title: "eslint/default-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/default-case.md for this page in Markdown format

# eslint/default-case Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#eslint-default-case)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#what-it-does)

Enforces that all `switch` statements include a `default` case, unless explicitly marked with a configured comment.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#why-is-this-bad)

Without a `default` case, it is unclear whether the omission was intentional or an oversight. Adding a `default` or a special comment makes the code more explicit and reduces mistakes.

You may optionally include a `// no default` after the last case if there is no default case. The comment may be in any desired case, such as `// No Default`.

Example configuration:

json

```
{
  "default-case": ["error", { "commentPattern": "^skip\\sdefault" }]
}
```

Examples of **incorrect** code for this rule:

js

```
/* default-case: ["error"] */

switch (foo) {
  case 1:
    break;
}
```

Examples of **correct** code for this rule:

js

```
/* default-case: ["error"] */

switch (a) {
  case 1:
    break;
  default:
    break;
}

switch (a) {
  case 1:
    break;
  // no default
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#configuration)

This rule accepts a configuration object with the following properties:

### commentPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#commentpattern)

type: `string | null`

A regex pattern used to detect comments that mark the absence of a `default` case as intentional.

Default value: `no default`.

Examples of **incorrect** code for this rule with the `{ "commentPattern": "^skip\\sdefault" }` option:

js

```
/* default-case: ["error", { "commentPattern": "^skip\sdefault" }] */

switch (a) {
  case 1:
    break;
  // no default
}
```

Examples of **correct** code for this rule with the `{ "commentPattern": "^skip\\sdefault" }` option:

js

```
/* default-case: ["error", { "commentPattern": "^skip\\sdefault" }] */

switch (a) {
  case 1:
    break;
  // skip default
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "default-case": "error"
  }
}
```

bash

```
oxlint --deny default-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/default_case.rs)
