---
title: "unicorn/explicit-length-check | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/explicit-length-check.md for this page in Markdown format

# unicorn/explicit-length-check Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#unicorn-explicit-length-check)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#what-it-does)

Enforce explicitly comparing the length or size property of a value.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#why-is-this-bad)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const isEmpty = foo.length == 0;
const isEmpty = foo.length < 1;
const isEmpty = 0 === foo.length;
const isEmpty = 0 == foo.length;
const isEmpty = 1 > foo.length;

const isEmpty = !foo.length;
const isEmpty = !(foo.length > 0);
const isEmptySet = !foo.size;
```

Examples of **correct** code for this rule:

javascript

```
const isEmpty = foo.length === 0;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#configuration)

This rule accepts a configuration object with the following properties:

### non-zero [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#non-zero)

type: `"greater-than" | "not-equal"`

default: `"greater-than"`

Configuration option to specify how non-zero length checks should be enforced.

`greater-than`: Enforces non-zero to be checked with `foo.length > 0``not-equal`: Enforces non-zero to be checked with `foo.length !== 0`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/explicit-length-check": "error"
  }
}
```

bash

```
oxlint --deny unicorn/explicit-length-check
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/explicit_length_check.rs)
