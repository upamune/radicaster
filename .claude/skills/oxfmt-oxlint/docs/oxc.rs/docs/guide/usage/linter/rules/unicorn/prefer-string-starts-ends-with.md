---
title: "unicorn/prefer-string-starts-ends-with | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.md for this page in Markdown format

# unicorn/prefer-string-starts-ends-with Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#unicorn-prefer-string-starts-ends-with)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#what-it-does)

Prefer [`String#startsWith()`](https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith) and [`String#endsWith()`](https://developer.mozilla.org/en/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith) over using a regex with `/^foo/` or `/foo$/`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#why-is-this-bad)

Using `String#startsWith()` and `String#endsWith()` is more readable and performant as it does not need to parse a regex.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = "hello";
/^abc/.test(foo);
```

Examples of **correct** code for this rule:

javascript

```
const foo = "hello";
foo.startsWith("abc");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-string-starts-ends-with": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-string-starts-ends-with
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_string_starts_ends_with.rs)
