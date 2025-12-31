---
title: "jest/no-jasmine-globals | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-jasmine-globals.md for this page in Markdown format

# jest/no-jasmine-globals Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#jest-no-jasmine-globals)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#what-it-does)

This rule reports on any usage of Jasmine globals, which is not ported to Jest, and suggests alternatives from Jest's own API.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#why-is-this-bad)

When migrating from Jasmine to Jest, relying on Jasmine-specific globals creates compatibility issues and prevents taking advantage of Jest's improved testing features and better error reporting.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
jasmine.DEFAULT_TIMEOUT_INTERVAL = 5000;
test("my test", () => {
  pending();
});
test("my test", () => {
  jasmine.createSpy();
});
```

Examples of **correct** code for this rule:

javascript

```
jest.setTimeout(5000);
test("my test", () => {
  // Use test.skip() instead of pending()
});
test.skip("my test", () => {
  // Skipped test
});
test("my test", () => {
  jest.fn(); // Use jest.fn() instead of jasmine.createSpy()
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-jasmine-globals": "error"
  }
}
```

bash

```
oxlint --deny jest/no-jasmine-globals --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_jasmine_globals.rs)
