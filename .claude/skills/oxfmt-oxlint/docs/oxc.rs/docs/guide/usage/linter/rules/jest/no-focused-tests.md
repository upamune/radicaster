---
title: "jest/no-focused-tests | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-focused-tests.md for this page in Markdown format

# jest/no-focused-tests Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#jest-no-focused-tests)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#what-it-does)

This rule reminds you to remove `.only` from your tests by raising a warning whenever you are using the exclusivity feature.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#why-is-this-bad)

Jest has a feature that allows you to focus tests by appending `.only` or prepending `f` to a test-suite or a test-case. This feature is really helpful to debug a failing test, so you don’t have to execute all of your tests. After you have fixed your test and before committing the changes you have to remove `.only` to ensure all tests are executed on your build system.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe.only("foo", () => {});
it.only("foo", () => {});
describe["only"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});
it["only"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});
test.only("foo", () => {});
test["only"](https://oxc.rs/docs/guide/usage/linter/rules/jest/"bar", () => {});
fdescribe("foo", () => {});
fit("foo", () => {});
fit.each`
  table
`();
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/no-focused-tests.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-focused-tests": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-focused-tests": "error"
  }
}
```

bash

```
oxlint --deny jest/no-focused-tests --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_focused_tests.rs)
