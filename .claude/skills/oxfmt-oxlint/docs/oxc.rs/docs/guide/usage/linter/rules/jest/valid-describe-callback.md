---
title: "jest/valid-describe-callback | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/valid-describe-callback.md for this page in Markdown format

# jest/valid-describe-callback Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#jest-valid-describe-callback)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#what-it-does)

This rule validates that the second parameter of a `describe()` function is a callback function. This callback function:

* should not be [async](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/async_function)
* should not contain any parameters
* should not contain any `return` statements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#why-is-this-bad)

Using an improper `describe()` callback function can lead to unexpected test errors.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
// Async callback functions are not allowed
describe("myFunction()", async () => {
  // ...
});

// Callback function parameters are not allowed
describe("myFunction()", (done) => {
  // ...
});

// Returning a value from a describe block is not allowed
describe("myFunction", () =>
  it("returns a truthy value", () => {
    expect(myFunction()).toBeTruthy();
  }));
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/valid-describe-callback.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/valid-describe-callback": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/valid-describe-callback": "error"
  }
}
```

bash

```
oxlint --deny jest/valid-describe-callback --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/valid_describe_callback.rs)
