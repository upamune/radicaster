---
title: "jest/prefer-hooks-in-order | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.md for this page in Markdown format

# jest/prefer-hooks-in-order Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#jest-prefer-hooks-in-order)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#what-it-does)

Ensures that hooks are in the order that they are called in.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#why-is-this-bad)

While hooks can be setup in any order, they're always called by `jest` in this specific order:

1. `beforeAll`
2. `beforeEach`
3. `afterEach`
4. `afterAll`

This rule aims to make that more obvious by enforcing grouped hooks be setup in that order within tests.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("foo", () => {
  beforeEach(() => {
    seedMyDatabase();
  });
  beforeAll(() => {
    createMyDatabase();
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
    it("accepts that input", () => {
      // ...
    });
    it("throws an error", () => {
      // ...
    });
    beforeEach(() => {
      mockLogger();
    });
    afterEach(() => {
      clearLogger();
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

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/prefer-hooks-in-order.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-hooks-in-order": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-hooks-in-order": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-hooks-in-order --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_hooks_in_order.rs)
