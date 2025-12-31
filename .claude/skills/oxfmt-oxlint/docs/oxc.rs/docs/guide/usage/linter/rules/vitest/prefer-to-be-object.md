---
title: "vitest/prefer-to-be-object | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/prefer-to-be-object.md for this page in Markdown format

# vitest/prefer-to-be-object Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#vitest-prefer-to-be-object)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#what-it-does)

This rule enforces using `toBeObject()` to check if a value is of type `Object`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#why-is-this-bad)

Using other methods such as `toBeInstanceOf(Object)` or `instanceof Object` can be less clear and potentially misleading. Enforcing the use of `toBeObject()` provides more explicit and readable code, making your intentions clear and improving the overall maintainability and readability of your tests.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#examples)

Examples of **incorrect** code for this rule:

js

```
expectTypeOf({}).toBeInstanceOf(Object);
expectTypeOf({} instanceof Object).toBeTruthy();
```

Examples of **correct** code for this rule:

js

```
expectTypeOf({}).toBeObject();
expectTypeOf({}).toBeObject();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/prefer-to-be-object": "error"
  }
}
```

bash

```
oxlint --deny vitest/prefer-to-be-object --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/prefer_to_be_object.rs)
