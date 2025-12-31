---
title: "jest/no-confusing-set-timeout | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.md for this page in Markdown format

# jest/no-confusing-set-timeout Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#jest-no-confusing-set-timeout)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#what-it-does)

Disallow confusing usages of jest.setTimeout

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#why-is-this-bad)

* being called anywhere other than in global scope
* being called multiple times
* being called after other Jest functions like hooks, `describe`, `test`, or `it`

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#example)

All of these are invalid case:

javascript

```
escribe("test foo", () => {
  jest.setTimeout(1000);
  it("test-description", () => {
    // test logic;
  });
});

describe("test bar", () => {
  it("test-description", () => {
    jest.setTimeout(1000);
    // test logic;
  });
});

test("foo-bar", () => {
  jest.setTimeout(1000);
});

describe("unit test", () => {
  beforeEach(() => {
    jest.setTimeout(1000);
  });
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-confusing-set-timeout": "error"
  }
}
```

bash

```
oxlint --deny jest/no-confusing-set-timeout --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_confusing_set_timeout.rs)
