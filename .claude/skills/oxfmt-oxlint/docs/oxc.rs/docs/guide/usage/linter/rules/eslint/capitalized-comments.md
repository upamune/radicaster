---
title: "eslint/capitalized-comments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/capitalized-comments.md for this page in Markdown format

# eslint/capitalized-comments Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#eslint-capitalized-comments)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#what-it-does)

Enforces or disallows capitalization of the first letter of a comment.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#why-is-this-bad)

Inconsistent capitalization of comments can make code harder to read. This rule helps enforce a consistent style across the codebase.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#examples)

Examples of **incorrect** code for this rule with the default `"always"` option:

js

```
// lowercase comment
/* lowercase block comment */
```

Examples of **correct** code for this rule with the default `"always"` option:

js

```
// Capitalized comment
/* Capitalized block comment */
// 123 - comments starting with non-letters are ignored
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#configuration)

This rule accepts a configuration object with the following properties:

### block [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#block)

type: `object | null`

#### block.ignoreConsecutiveComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#block-ignoreconsecutivecomments)

type: `boolean | null`

#### block.ignoreInlineComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#block-ignoreinlinecomments)

type: `boolean | null`

#### block.ignorePattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#block-ignorepattern)

type: `string | null`

### ignoreConsecutiveComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#ignoreconsecutivecomments)

type: `boolean | null`

### ignoreInlineComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#ignoreinlinecomments)

type: `boolean | null`

### ignorePattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#ignorepattern)

type: `string | null`

### line [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#line)

type: `object | null`

#### line.ignoreConsecutiveComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#line-ignoreconsecutivecomments)

type: `boolean | null`

#### line.ignoreInlineComments [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#line-ignoreinlinecomments)

type: `boolean | null`

#### line.ignorePattern [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#line-ignorepattern)

type: `string | null`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "capitalized-comments": "error"
  }
}
```

bash

```
oxlint --deny capitalized-comments
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/capitalized_comments.rs)
