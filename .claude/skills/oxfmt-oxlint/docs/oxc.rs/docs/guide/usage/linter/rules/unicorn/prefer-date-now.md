---
title: "unicorn/prefer-date-now | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-date-now.md for this page in Markdown format

# unicorn/prefer-date-now Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#unicorn-prefer-date-now)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#what-it-does)

Prefers use of `Date.now()` over `new Date().getTime()` or `new Date().valueOf()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#why-is-this-bad)

Using `Date.now()` is shorter and nicer than `new Date().getTime()`, and avoids unnecessary instantiation of `Date` objects.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const ts = new Date().getTime();
const ts = new Date().valueOf();
```

Examples of **correct** code for this rule:

javascript

```
const ts = Date.now();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-date-now": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-date-now
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_date_now.rs)
