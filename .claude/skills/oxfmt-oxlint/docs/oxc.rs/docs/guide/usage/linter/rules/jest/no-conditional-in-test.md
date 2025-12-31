---
title: "jest/no-conditional-in-test | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-conditional-in-test.md for this page in Markdown format

# jest/no-conditional-in-test Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#jest-no-conditional-in-test)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#what-it-does)

Disallow conditional statements in tests.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#why-is-this-bad)

Conditional statements in tests can make the test harder to read and understand. It is better to have a single test case per test function.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#examples)

Examples of **incorrect** code for this rule:

js

```
it("foo", () => {
  if (true) {
    doTheThing();
  }
});

it("bar", () => {
  switch (mode) {
    case "none":
      generateNone();
    case "single":
      generateOne();
    case "multiple":
      generateMany();
  }

  expect(fixtures.length).toBeGreaterThan(-1);
});

it("baz", async () => {
  const promiseValue = () => {
    return something instanceof Promise ? something : Promise.resolve(something);
  };

  await expect(promiseValue()).resolves.toBe(1);
});
```

Examples of **correct** code for this rule:

js

```
describe("my tests", () => {
  if (true) {
    it("foo", () => {
      doTheThing();
    });
  }
});

beforeEach(() => {
  switch (mode) {
    case "none":
      generateNone();
    case "single":
      generateOne();
    case "multiple":
      generateMany();
  }
});

it("bar", () => {
  expect(fixtures.length).toBeGreaterThan(-1);
});

const promiseValue = (something) => {
  return something instanceof Promise ? something : Promise.resolve(something);
};

it("baz", async () => {
  await expect(promiseValue()).resolves.toBe(1);
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-conditional-in-test.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-conditional-in-test": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-conditional-in-test": "error"
  }
}
```

bash

```
oxlint --deny jest/no-conditional-in-test --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_conditional_in_test.rs)
