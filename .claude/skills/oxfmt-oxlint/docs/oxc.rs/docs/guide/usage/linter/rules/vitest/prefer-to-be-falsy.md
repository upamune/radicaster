---
title: "vitest/prefer-to-be-falsy | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.md for this page in Markdown format

# vitest/prefer-to-be-falsy Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#vitest-prefer-to-be-falsy)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#what-it-does)

This rule warns when `toBe(false)` is used with `expect` or `expectTypeOf`. With `--fix`, it will be replaced with `toBeFalsy()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#why-is-this-bad)

Using `toBe(false)` is less expressive and may not account for other falsy values like `0`, `null`, or `undefined`. `toBeFalsy()` provides a more comprehensive check for any falsy value, improving the robustness of the tests.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(foo).toBe(false);
expectTypeOf(foo).toBe(false);
```

Examples of **correct** code for this rule:

javascript

```
expect(foo).toBeFalsy();
expectTypeOf(foo).toBeFalsy();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/prefer-to-be-falsy": "error"
  }
}
```

bash

```
oxlint --deny vitest/prefer-to-be-falsy --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/prefer_to_be_falsy.rs)
