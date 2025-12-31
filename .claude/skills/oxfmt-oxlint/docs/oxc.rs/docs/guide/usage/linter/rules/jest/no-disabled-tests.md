---
title: "jest/no-disabled-tests | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-disabled-tests.md for this page in Markdown format

# jest/no-disabled-tests Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#jest-no-disabled-tests)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#what-it-does)

This rule raises a warning about disabled tests.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#why-is-this-bad)

Jest has a feature that allows you to temporarily mark tests as disabled. This feature is often helpful while debugging or to create placeholders for future tests. Before committing changes we may want to check that all tests are running.

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#example)

js

```
describe.skip("foo", () => {});
it.skip("foo", () => {});
test.skip("foo", () => {});

describe["skip"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});
it["skip"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});
test["skip"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});

xdescribe("foo", () => {});
xit("foo", () => {});
xtest("foo", () => {});

it("bar");
test("bar");

it("foo", () => {
  pending();
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/no-disabled-tests.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-disabled-tests": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-disabled-tests": "error"
  }
}
```

bash

```
oxlint --deny jest/no-disabled-tests --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_disabled_tests.rs)
