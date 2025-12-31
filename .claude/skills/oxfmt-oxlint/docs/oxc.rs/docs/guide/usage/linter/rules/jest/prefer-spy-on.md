---
title: "jest/prefer-spy-on | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-spy-on.md for this page in Markdown format

# jest/prefer-spy-on Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#jest-prefer-spy-on)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#what-it-does)

When mocking a function by overwriting a property you have to manually restore the original implementation when cleaning up. When using `jest.spyOn()` Jest keeps track of changes, and they can be restored with `jest.restoreAllMocks()`, `mockFn.mockRestore()` or by setting `restoreMocks` to `true` in the Jest config.

Note: The mock created by `jest.spyOn()` still behaves the same as the original function. The original function can be overwritten with `mockFn.mockImplementation()` or by some of the [other mock functions](https://jestjs.io/docs/en/mock-function-api).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#why-is-this-bad)

Directly overwriting properties with mock functions can lead to cleanup issues and test isolation problems. When you manually assign a mock to a property, you're responsible for restoring the original implementation, which is easy to forget and can cause tests to interfere with each other. Using `jest.spyOn()` provides automatic cleanup capabilities and makes your tests more reliable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
Date.now = jest.fn();
Date.now = jest.fn(() => 10);
```

Examples of **correct** code for this rule:

javascript

```
jest.spyOn(Date, "now");
jest.spyOn(Date, "now").mockImplementation(() => 10);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-spy-on.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-spy-on": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-spy-on": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-spy-on --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_spy_on.rs)
