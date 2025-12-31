---
title: "eslint/no-empty-character-class | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-empty-character-class.md for this page in Markdown format

# eslint/no-empty-character-class Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#eslint-no-empty-character-class)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#what-it-does)

Disallow empty character classes in regular expressions

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#why-is-this-bad)

Because empty character classes in regular expressions do not match anything, they might be typing mistakes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = /^abc[]/;
```

Examples of **correct** code for this rule:

javascript

```
var foo = /^abc/;
var foo2 = /^abc[123]/;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-empty-character-class": "error"
  }
}
```

bash

```
oxlint --deny no-empty-character-class
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_empty_character_class.rs)
