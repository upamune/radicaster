---
title: "unicorn/prefer-string-trim-start-end | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.md for this page in Markdown format

# unicorn/prefer-string-trim-start-end Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#unicorn-prefer-string-trim-start-end)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#what-it-does)

[`String#trimLeft()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimLeft) and [`String#trimRight()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimRight) are aliases of [`String#trimStart()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimStart) and [`String#trimEnd()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd). This is to ensure consistency and use [direction](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Handling_different_text_directions)-independent wording.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#why-is-this-bad)

The `trimLeft` and `trimRight` names are confusing and inconsistent with the rest of the language.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
str.trimLeft();
str.trimRight();
```

Examples of **correct** code for this rule:

javascript

```
str.trimStart();
str.trimEnd();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-string-trim-start-end": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-string-trim-start-end
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_string_trim_start_end.rs)
