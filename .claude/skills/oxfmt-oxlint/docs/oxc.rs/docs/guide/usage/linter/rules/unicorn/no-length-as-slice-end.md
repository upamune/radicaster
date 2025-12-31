---
title: "unicorn/no-length-as-slice-end | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.md for this page in Markdown format

# unicorn/no-length-as-slice-end Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#unicorn-no-length-as-slice-end)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#what-it-does)

Disallow using `length` as the end argument of a `slice` call.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#why-is-this-bad)

Passing `length` as the end argument of a `slice` call is unnecessary and can be confusing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.slice(1, foo.length);
```

Examples of **correct** code for this rule:

javascript

```
foo.slice(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-length-as-slice-end": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-length-as-slice-end
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_length_as_slice_end.rs)
