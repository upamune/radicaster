---
title: "jest/no-done-callback | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-done-callback.md for this page in Markdown format

# jest/no-done-callback Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#jest-no-done-callback)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#what-it-does)

This rule checks the function parameter of hooks & tests for use of the done argument, suggesting you return a promise instead.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#why-is-this-bad)

When calling asynchronous code in hooks and tests, jest needs to know when the asynchronous work is complete to progress the current run. Originally the most common pattern to achieve this was to use callbacks:

javascript

```
test("the data is peanut butter", (done) => {
  function callback(data) {
    try {
      expect(data).toBe("peanut butter");
      done();
    } catch (error) {
      done(error);
    }
  }

  fetchData(callback);
});
```

This can be very error-prone however, as it requires careful understanding of how assertions work in tests or otherwise tests won't behave as expected.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
beforeEach((done) => {
  // ...
});

test("myFunction()", (done) => {
  // ...
});

test("myFunction()", function (done) {
  // ...
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-done-callback": "error"
  }
}
```

bash

```
oxlint --deny jest/no-done-callback --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_done_callback.rs)
