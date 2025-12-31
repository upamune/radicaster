---
title: "jest/prefer-hooks-on-top | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.md for this page in Markdown format

# jest/prefer-hooks-on-top Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#jest-prefer-hooks-on-top)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#what-it-does)

While hooks can be setup anywhere in a test file, they are always called in a specific order, which means it can be confusing if they're intermixed with test cases.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#why-is-this-bad)

When hooks are mixed with test cases, it becomes harder to understand the test setup and execution order. This can lead to confusion about which hooks apply to which tests and when they run. Grouping hooks at the top of each `describe` block makes the test structure clearer and more maintainable.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("foo", () => {
  beforeEach(() => {
    seedMyDatabase();
  });

  it("accepts this input", () => {
    // ...
  });

  beforeAll(() => {
    createMyDatabase();
  });

  it("returns that value", () => {
    // ...
  });

  describe("when the database has specific values", () => {
    const specificValue = "...";
    beforeEach(() => {
      seedMyDatabase(specificValue);
    });

    it("accepts that input", () => {
      // ...
    });

    it("throws an error", () => {
      // ...
    });

    afterEach(() => {
      clearLogger();
    });

    beforeEach(() => {
      mockLogger();
    });

    it("logs a message", () => {
      // ...
    });
  });

  afterAll(() => {
    removeMyDatabase();
  });
});
```

Examples of **correct** code for this rule:

javascript

```
describe("foo", () => {
  beforeAll(() => {
    createMyDatabase();
  });

  beforeEach(() => {
    seedMyDatabase();
  });

  afterAll(() => {
    clearMyDatabase();
  });

  it("accepts this input", () => {
    // ...
  });

  it("returns that value", () => {
    // ...
  });

  describe("when the database has specific values", () => {
    const specificValue = "...";

    beforeEach(() => {
      seedMyDatabase(specificValue);
    });

    beforeEach(() => {
      mockLogger();
    });

    afterEach(() => {
      clearLogger();
    });

    it("accepts that input", () => {
      // ...
    });

    it("throws an error", () => {
      // ...
    });

    it("logs a message", () => {
      // ...
    });
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-hooks-on-top.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-hooks-on-top": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-hooks-on-top": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-hooks-on-top --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_hooks_on_top.rs)
