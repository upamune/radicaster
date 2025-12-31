---
title: "unicorn/consistent-date-clone | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/consistent-date-clone.md for this page in Markdown format

# unicorn/consistent-date-clone Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#unicorn-consistent-date-clone)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#what-it-does)

The Date constructor can clone a `Date` object directly when passed as an argument, making timestamp conversion unnecessary. This rule enforces the use of the direct `Date` cloning instead of using `.getTime()` for conversion.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#why-is-this-bad)

Using `.getTime()` to convert a `Date` object to a timestamp and then back to a `Date` is redundant and unnecessary. Simply passing the `Date` object to the `Date` constructor is cleaner and more efficient.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#examples)

Examples of **incorrect** code for this rule:

js

```
new Date(date.getTime());
```

Examples of **correct** code for this rule:

js

```
new Date(date);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/consistent-date-clone": "error"
  }
}
```

bash

```
oxlint --deny unicorn/consistent-date-clone
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/consistent_date_clone.rs)
