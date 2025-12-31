---
title: "unicorn/no-abusive-eslint-disable | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.md for this page in Markdown format

# unicorn/no-abusive-eslint-disable Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#unicorn-no-abusive-eslint-disable)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#what-it-does)

Disallows `oxlint-disable` or `eslint-disable` comments without specifying rules.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#why-is-this-bad)

A general `oxlint-disable` or `eslint-disable` comment suppresses all lint errors, not just the intended one, potentially hiding useful warnings and making debugging harder.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/* eslint-disable */
console.log(message);

console.log(message); // eslint-disable-line

// eslint-disable-next-line
console.log(message);
```

javascript

```
/* oxlint-disable */
console.log(message);

console.log(message); // oxlint-disable-line

// oxlint-disable-next-line
console.log(message);
```

Examples of **correct** code for this rule:

javascript

```
/* eslint-disable no-console */
console.log(message);

console.log(message); // eslint-disable-line no-console

// eslint-disable-next-line no-console
console.log(message);
```

javascript

```
/* oxlint-disable no-console */
console.log(message);

console.log(message); // oxlint-disable-line no-console

// oxlint-disable-next-line no-console
console.log(message);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-abusive-eslint-disable": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-abusive-eslint-disable
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_abusive_eslint_disable.rs)
