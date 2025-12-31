---
title: "jest/require-to-throw-message | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/require-to-throw-message.md for this page in Markdown format

# jest/require-to-throw-message Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#jest-require-to-throw-message)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#what-it-does)

This rule triggers a warning if `toThrow()` or `toThrowError()` is used without an error message.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#why-is-this-bad)

Using `toThrow()` or `toThrowError()` without specifying an expected error message makes tests less specific and harder to debug. When a test only checks that an error was thrown but not what kind of error, it can pass even when the wrong error is thrown, potentially hiding bugs. Providing an expected error message or error type makes tests more precise and helps catch regressions more effectively.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
test("all the things", async () => {
  expect(() => a()).toThrow();
  expect(() => a()).toThrowError();
  await expect(a()).rejects.toThrow();
  await expect(a()).rejects.toThrowError();
});
```

Examples of **correct** code for this rule:

javascript

```
test("all the things", async () => {
  expect(() => a()).toThrow("a");
  expect(() => a()).toThrowError("a");
  await expect(a()).rejects.toThrow("a");
  await expect(a()).rejects.toThrowError("a");
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/require-to-throw-message.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/require-to-throw-message": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/require-to-throw-message": "error"
  }
}
```

bash

```
oxlint --deny jest/require-to-throw-message --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/require_to_throw_message.rs)
