---
title: "jest/no-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-export.md for this page in Markdown format

# jest/no-export Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#jest-no-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#what-it-does)

Prevents using exports if a file has one or more tests in it.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#why-is-this-bad)

This rule aims to eliminate duplicate runs of tests by exporting things from test files. If you import from a test file, then all the tests in that file will be run in each imported instance. so bottom line, don't export from a test, but instead move helper functions into a separate file when they need to be shared across tests.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
export function myHelper() {}
describe("a test", () => {
  expect(1).toBe(1);
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-export": "error"
  }
}
```

bash

```
oxlint --deny jest/no-export --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_export.rs)
