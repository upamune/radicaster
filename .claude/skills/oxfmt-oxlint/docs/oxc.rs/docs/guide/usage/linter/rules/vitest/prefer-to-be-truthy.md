---
title: "vitest/prefer-to-be-truthy | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.md for this page in Markdown format

# vitest/prefer-to-be-truthy Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#vitest-prefer-to-be-truthy)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#what-it-does)

This rule warns when `toBe(true)` is used with `expect` or `expectTypeOf`. With `--fix`, it will be replaced with `toBeTruthy()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#why-is-this-bad)

Using `toBe(true)` is less flexible and may not account for other truthy values like non-empty strings or objects. `toBeTruthy()` checks for any truthy value, which makes the tests more comprehensive and robust.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(foo).toBe(true);
expectTypeOf(foo).toBe(true);
```

Examples of **correct** code for this rule:

javascript

```
expect(foo).toBeTruthy();
expectTypeOf(foo).toBeTruthy();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/prefer-to-be-truthy": "error"
  }
}
```

bash

```
oxlint --deny vitest/prefer-to-be-truthy --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/prefer_to_be_truthy.rs)
