---
title: "jest/padding-around-test-blocks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/padding-around-test-blocks.md for this page in Markdown format

# jest/padding-around-test-blocks Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html#jest-padding-around-test-blocks)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html#what-it-does)

This rule enforces a line of padding before and after 1 or more test/it statements

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html#examples)

Examples of **incorrect** code for this rule:

js

```
const thing = 123;
test("foo", () => {});
test("bar", () => {});
```

js

```
const thing = 123;
it("foo", () => {});
it("bar", () => {});
```

Examples of **correct** code for this rule:

js

```
const thing = 123;

test("foo", () => {});

test("bar", () => {});
```

js

```
const thing = 123;

it("foo", () => {});

it("bar", () => {});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/padding-around-test-blocks": "error"
  }
}
```

bash

```
oxlint --deny jest/padding-around-test-blocks --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/padding_around_test_blocks.rs)
