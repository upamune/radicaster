---
title: "jest/no-untyped-mock-factory | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.md for this page in Markdown format

# jest/no-untyped-mock-factory Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#jest-no-untyped-mock-factory)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#what-it-does)

This rule triggers a warning if `mock()` or `doMock()` is used without a generic type parameter or return type.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#why-is-this-bad)

By default, `jest.mock` and `jest.doMock` allow any type to be returned by a mock factory. A generic type parameter can be used to enforce that the factory returns an object with the same shape as the original module, or some other strict type. Requiring a type makes it easier to use TypeScript to catch changes needed in test mocks when the source module changes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
jest.mock("../moduleName", () => {
  return jest.fn(() => 42);
});

jest.mock("./module", () => ({
  ...jest.requireActual("./module"),
  foo: jest.fn(),
}));

jest.mock("random-num", () => {
  return jest.fn(() => 42);
});
```

Examples of **correct** code for this rule:

typescript

```
// Uses typeof import()
jest.mock<typeof import("../moduleName")>("../moduleName", () => {
  return jest.fn(() => 42);
});

jest.mock<typeof import("./module")>("./module", () => ({
  ...jest.requireActual("./module"),
  foo: jest.fn(),
}));

// Uses custom type
jest.mock<() => number>("random-num", () => {
  return jest.fn(() => 42);
});

// No factory
jest.mock("random-num");

// Virtual mock
jest.mock(
  "../moduleName",
  () => {
    return jest.fn(() => 42);
  },
  { virtual: true },
);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-untyped-mock-factory": "error"
  }
}
```

bash

```
oxlint --deny jest/no-untyped-mock-factory --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_untyped_mock_factory.rs)
