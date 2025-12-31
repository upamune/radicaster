---
title: "jest/prefer-strict-equal | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-strict-equal.md for this page in Markdown format

# jest/prefer-strict-equal Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#jest-prefer-strict-equal)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#what-it-does)

This rule triggers a warning if `toEqual()` is used to assert equality.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#why-is-this-bad)

The `toEqual()` matcher performs a deep equality check but ignores `undefined` values in objects and arrays. This can lead to false positives where tests pass when they should fail. `toStrictEqual()` provides more accurate comparison by checking for `undefined` values.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect({ a: "a", b: undefined }).toEqual({ a: "a" });
```

Examples of **correct** code for this rule:

javascript

```
expect({ a: "a", b: undefined }).toStrictEqual({ a: "a" });
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-strict-equal.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-strict-equal": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-strict-equal": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-strict-equal --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_strict_equal.rs)
