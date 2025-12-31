---
title: "unicorn/number-literal-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/number-literal-case.md for this page in Markdown format

# unicorn/number-literal-case Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#unicorn-number-literal-case)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#what-it-does)

This rule enforces proper case for numeric literals.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#why-is-this-bad)

When both an identifier and a number literal are in lower case, it can be hard to differentiate between them.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const foo = 0xff;
const foo = 0xff;
const foo = 0xff;
const foo = 0xffn;

const foo = 0b10;
const foo = 0b10n;

const foo = 0o76;
const foo = 0o76n;

const foo = 2e-5;
```

Examples of **correct** code for this rule:

javascript

```
const foo = 0xff;
const foo = 0b10;
const foo = 0o76;
const foo = 0xffn;
const foo = 2e5;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/number-literal-case": "error"
  }
}
```

bash

```
oxlint --deny unicorn/number-literal-case
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/number_literal_case.rs)
