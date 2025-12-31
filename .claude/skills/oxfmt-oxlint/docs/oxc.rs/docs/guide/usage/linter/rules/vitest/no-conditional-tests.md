---
title: "vitest/no-conditional-tests | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/no-conditional-tests.md for this page in Markdown format

# vitest/no-conditional-tests Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#vitest-no-conditional-tests)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#what-it-does)

The rule disallows the use of conditional statements within test cases to ensure that tests are deterministic and clearly readable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#why-is-this-bad)

Conditional statements in test cases can make tests unpredictable and harder to understand. Tests should be consistent and straightforward to ensure reliable results and maintainability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#examples)

Examples of **incorrect** code for this rule:

js

```
describe("my tests", () => {
  if (true) {
    it("is awesome", () => {
      doTheThing();
    });
  }
});
```

Examples of **correct** code for this rule:

js

```
describe("my tests", () => {
  it("is awesome", () => {
    doTheThing();
  });
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/no-conditional-tests": "error"
  }
}
```

bash

```
oxlint --deny vitest/no-conditional-tests --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/no_conditional_tests.rs)
