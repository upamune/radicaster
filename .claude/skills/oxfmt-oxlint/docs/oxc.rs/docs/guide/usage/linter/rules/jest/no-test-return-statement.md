---
title: "jest/no-test-return-statement | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-test-return-statement.md for this page in Markdown format

# jest/no-test-return-statement Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#jest-no-test-return-statement)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#what-it-does)

Disallow explicitly returning from tests.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#why-is-this-bad)

Tests in Jest should be void and not return values. If you are returning Promises then you should update the test to use `async/await`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
test("one", () => {
  return expect(1).toBe(1);
});
```

Examples of **correct** code for this rule:

javascript

```
test("one", () => {
  expect(1).toBe(1);
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-test-return-statement.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-test-return-statement": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-test-return-statement": "error"
  }
}
```

bash

```
oxlint --deny jest/no-test-return-statement --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_test_return_statement.rs)
