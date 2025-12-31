---
title: "unicorn/no-invalid-fetch-options | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.md for this page in Markdown format

# unicorn/no-invalid-fetch-options Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#unicorn-no-invalid-fetch-options)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#what-it-does)

Disallow invalid options in `fetch()` and `new Request()`. Specifically, this rule ensures that a body is not provided when the method is `GET` or `HEAD`, as it will result in a `TypeError`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#why-is-this-bad)

The `fetch()` function throws a `TypeError` when the method is `GET` or `HEAD` and a body is provided. This can lead to unexpected behavior and errors in your code. By disallowing such invalid options, the rule ensures that requests are correctly configured and prevents unnecessary errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const response = await fetch("/", { method: "GET", body: "foo=bar" });

const request = new Request("/", { method: "GET", body: "foo=bar" });
```

Examples of **correct** code for this rule:

javascript

```
const response = await fetch("/", { method: "POST", body: "foo=bar" });

const request = new Request("/", { method: "POST", body: "foo=bar" });
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-invalid-fetch-options": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-invalid-fetch-options
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_invalid_fetch_options.rs)
