---
title: "jest/no-identical-title | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-identical-title.md for this page in Markdown format

# jest/no-identical-title Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#jest-no-identical-title)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#what-it-does)

This rule looks at the title of every test and test suite. It will report when two test suites or two test cases at the same level of a test suite have the same title.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#why-is-this-bad)

Having identical titles for two different tests or test suites may create confusion. For example, when a test with the same title as another test in the same test suite fails, it is harder to know which one failed and thus harder to fix.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("baz", () => {
  //...
});

describe("baz", () => {
  // Has the same title as a previous test suite
  // ...
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/no-identical-title.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-identical-title": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-identical-title": "error"
  }
}
```

bash

```
oxlint --deny jest/no-identical-title --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_identical_title.rs)
