---
title: "jest/consistent-test-it | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/consistent-test-it.md for this page in Markdown format

# jest/consistent-test-it Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#jest-consistent-test-it)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#what-it-does)

Jest allows you to choose how you want to define your tests, using the `it` or the `test` keywords, with multiple permutations for each:

* **it:** `it`, `xit`, `fit`, `it.only`, `it.skip`.
* **test:** `test`, `xtest`, `test.only`, `test.skip`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#why-is-this-bad)

It's a good practice to be consistent in your test suite, so that all tests are written in the same way.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#examples)

javascript

```
/*eslint jest/consistent-test-it: ["error", {"fn": "test"}]*/
test("foo"); // valid
test.only("foo"); // valid

it("foo"); // invalid
it.only("foo"); // invalid
```

javascript

```
/*eslint jest/consistent-test-it: ["error", {"fn": "it"}]*/
it("foo"); // valid
it.only("foo"); // valid
test("foo"); // invalid
test.only("foo"); // invalid
```

javascript

```
/*eslint jest/consistent-test-it: ["error", {"fn": "it", "withinDescribe": "test"}]*/
it("foo"); // valid
describe("foo", function () {
  test("bar"); // valid
});

test("foo"); // invalid
describe("foo", function () {
  it("bar"); // invalid
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/consistent-test-it.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/consistent-test-it": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#configuration)

This rule accepts a configuration object with the following properties:

### fn [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#fn)

type: `"it" | "test"`

default: `"test"`

Decides whether to use `test` or `it`.

### withinDescribe [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#withindescribe)

type: `"it" | "test"`

default: `"it"`

Decides whether to use `test` or `it` within a `describe` scope. If only `fn` is provided, this will default to the value of `fn`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/consistent-test-it": "error"
  }
}
```

bash

```
oxlint --deny jest/consistent-test-it --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/consistent_test_it.rs)
