---
title: "jest/max-expects | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/max-expects.md for this page in Markdown format

# jest/max-expects Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#jest-max-expects)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#what-it-does)

This rule enforces a maximum number of `expect()` calls in a single test.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#why-is-this-bad)

Tests with many different assertions are likely mixing multiple objectives. It is generally better to have a single objective per test to ensure that when a test fails, the problem is easy to identify.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
test("should not pass", () => {
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
});

it("should not pass", () => {
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
  expect(true).toBeDefined();
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/max-expects.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/max-expects": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#max)

type: `integer`

default: `5`

Maximum number of `expect()` assertion calls allowed within a single test.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/max-expects": "error"
  }
}
```

bash

```
oxlint --deny jest/max-expects --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/max_expects.rs)
