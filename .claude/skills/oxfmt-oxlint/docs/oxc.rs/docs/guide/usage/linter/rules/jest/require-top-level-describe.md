---
title: "jest/require-top-level-describe | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/require-top-level-describe.md for this page in Markdown format

# jest/require-top-level-describe Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#jest-require-top-level-describe)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#what-it-does)

Requires test cases and hooks to be inside a top-level `describe` block.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#why-is-this-bad)

Having tests and hooks organized within `describe` blocks provides better structure and grouping for test suites. It makes test output more readable and helps with test organization, especially in larger codebases.

This rule triggers a warning if a test case (`test` and `it`) or a hook (`beforeAll`, `beforeEach`, `afterEach`, `afterAll`) is not located in a top-level `describe` block.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Above a describe block
test("my test", () => {});
describe("test suite", () => {
  it("test", () => {});
});

// Below a describe block
describe("test suite", () => {});
test("my test", () => {});

// Same for hooks
beforeAll("my beforeAll", () => {});
describe("test suite", () => {});
afterEach("my afterEach", () => {});
```

Examples of **correct** code for this rule:

javascript

```
// Above a describe block
// In a describe block
describe("test suite", () => {
  test("my test", () => {});
});

// In a nested describe block
describe("test suite", () => {
  test("my test", () => {});
  describe("another test suite", () => {
    test("my other test", () => {});
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/require-top-level-describe.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/require-top-level-describe": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#configuration)

This rule accepts a configuration object with the following properties:

### maxNumberOfTopLevelDescribes [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#maxnumberoftopleveldescribes)

type: `integer`

default: `18446744073709551615`

The maximum number of top-level `describe` blocks allowed in a test file.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/require-top-level-describe": "error"
  }
}
```

bash

```
oxlint --deny jest/require-top-level-describe --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/require_top_level_describe.rs)
