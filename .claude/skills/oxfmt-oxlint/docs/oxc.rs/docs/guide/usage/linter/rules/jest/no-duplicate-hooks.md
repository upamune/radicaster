---
title: "jest/no-duplicate-hooks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-duplicate-hooks.md for this page in Markdown format

# jest/no-duplicate-hooks Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#jest-no-duplicate-hooks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#what-it-does)

Disallows duplicate hooks in describe blocks.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#why-is-this-bad)

Having duplicate hooks in a describe block can lead to confusion and unexpected behavior. When multiple hooks of the same type exist, they all execute in order, which can make it difficult to understand the test setup flow and may result in redundant or conflicting operations. This makes tests harder to maintain and debug.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("foo", () => {
  beforeEach(() => {
    // some setup
  });
  beforeEach(() => {
    // some setup
  });
  test("foo_test", () => {
    // some test
  });
});

// Nested describe scenario
describe("foo", () => {
  beforeEach(() => {
    // some setup
  });
  test("foo_test", () => {
    // some test
  });
  describe("bar", () => {
    test("bar_test", () => {
      afterAll(() => {
        // some teardown
      });
      afterAll(() => {
        // some teardown
      });
    });
  });
});
```

Examples of **correct** code for this rule:

javascript

```
describe("foo", () => {
  beforeEach(() => {
    // some setup
  });
  test("foo_test", () => {
    // some test
  });
});

// Nested describe scenario
describe("foo", () => {
  beforeEach(() => {
    // some setup
  });
  test("foo_test", () => {
    // some test
  });
  describe("bar", () => {
    test("bar_test", () => {
      beforeEach(() => {
        // some setup
      });
    });
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-duplicate-hooks.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-duplicate-hooks": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-duplicate-hooks": "error"
  }
}
```

bash

```
oxlint --deny jest/no-duplicate-hooks --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_duplicate_hooks.rs)
