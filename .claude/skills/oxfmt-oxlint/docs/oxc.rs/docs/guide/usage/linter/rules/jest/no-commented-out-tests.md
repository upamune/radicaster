---
title: "jest/no-commented-out-tests | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-commented-out-tests.md for this page in Markdown format

# jest/no-commented-out-tests Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#jest-no-commented-out-tests)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#what-it-does)

This rule raises a warning about commented out tests. It's similar to the `no-disabled-tests` rule.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#why-is-this-bad)

You may forget to uncomment some tests. This rule raises a warning about commented-out tests.

It is generally better to skip a test if it's flaky, or remove it if it's no longer needed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// describe('foo', () => {});
// it('foo', () => {});
// test('foo', () => {});

// describe.skip('foo', () => {});
// it.skip('foo', () => {});
// test.skip('foo', () => {});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/no-commented-out-tests.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-commented-out-tests": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-commented-out-tests": "error"
  }
}
```

bash

```
oxlint --deny jest/no-commented-out-tests --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_commented_out_tests.rs)
