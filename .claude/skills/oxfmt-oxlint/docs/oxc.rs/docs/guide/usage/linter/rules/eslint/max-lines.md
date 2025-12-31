---
title: "eslint/max-lines | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/max-lines.md for this page in Markdown format

# eslint/max-lines Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#eslint-max-lines)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#what-it-does)

Enforce a maximum number of lines per file.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#why-is-this-bad)

Some people consider large files a code smell. Large files tend to do a lot of things and can make it hard following what’s going. While there is not an objective maximum number of lines considered acceptable in a file, most people would agree it should not be in the thousands. Recommendations usually range from 100 to 500 lines.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#max)

type: `integer`

default: `300`

Maximum number of lines allowed per file.

### skipBlankLines [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#skipblanklines)

type: `boolean`

default: `false`

Whether to ignore blank lines when counting.

### skipComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#skipcomments)

type: `boolean`

default: `false`

Whether to ignore comments when counting.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "max-lines": "error"
  }
}
```

bash

```
oxlint --deny max-lines
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/max_lines.rs)
