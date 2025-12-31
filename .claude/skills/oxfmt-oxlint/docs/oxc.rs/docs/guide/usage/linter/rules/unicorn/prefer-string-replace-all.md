---
title: "unicorn/prefer-string-replace-all | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.md for this page in Markdown format

# unicorn/prefer-string-replace-all Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#unicorn-prefer-string-replace-all)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#what-it-does)

Prefers [`String#replaceAll()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replaceAll) over [`String#replace()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replace) when using a regex with the global flag.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#why-is-this-bad)

The [`String#replaceAll()`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/replaceAll) method is both faster and safer as you don't have to use a regex and remember to escape it if the string is not a literal. And when used with a regex, it makes the intent clearer.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#examples)

Examples of **incorrect** code for this rule:

js

```
foo.replace(/a/g, bar);
```

Examples of **correct** code for this rule:

js

```
foo.replace(/a/, bar);
foo.replaceAll(/a/, bar);

const pattern = "not-a-regexp";
foo.replace(pattern, bar);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-string-replace-all": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-string-replace-all
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_string_replace_all.rs)
