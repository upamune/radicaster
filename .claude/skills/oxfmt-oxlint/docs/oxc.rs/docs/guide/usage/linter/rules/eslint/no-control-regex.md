---
title: "eslint/no-control-regex | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-control-regex.md for this page in Markdown format

# eslint/no-control-regex Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#eslint-no-control-regex)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#what-it-does)

Disallows control characters and some escape sequences that match control characters in regular expressions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#why-is-this-bad)

Control characters are special, invisible characters in the ASCII range 0-31. These characters are rarely used in JavaScript strings so a regular expression containing elements that explicitly match these characters is most likely a mistake.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var pattern1 = /\x00/;
var pattern2 = /\x0C/;
var pattern3 = /\x1F/;
var pattern4 = /\u000C/;
var pattern5 = /\u{C}/u;
var pattern6 = new RegExp("\x0C"); // raw U+000C character in the pattern
var pattern7 = new RegExp("\\x0C"); // \x0C pattern
```

Examples of **correct** code for this rule:

javascript

```
var pattern1 = /\x20/;
var pattern2 = /\u0020/;
var pattern3 = /\u{20}/u;
var pattern4 = /\t/;
var pattern5 = /\n/;
var pattern6 = new RegExp("\x20");
var pattern7 = new RegExp("\\t");
var pattern8 = new RegExp("\\n");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-control-regex": "error"
  }
}
```

bash

```
oxlint --deny no-control-regex
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_control_regex.rs)
