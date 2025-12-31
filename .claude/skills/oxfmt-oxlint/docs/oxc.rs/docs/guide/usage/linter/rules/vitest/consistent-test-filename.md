---
title: "vitest/consistent-test-filename | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/consistent-test-filename.md for this page in Markdown format

# vitest/consistent-test-filename Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#vitest-consistent-test-filename)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#what-it-does)

This rule triggers an error when a file is considered a test file, but its name does not match an expected filename format.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#why-is-this-bad)

Files that are tests but with an unexpected filename make it hard to distinguish between source code files and test files.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#examples)

An example of an **incorrect** file path for this rule configured as `{"allTestPattern": "__tests__", "pattern": ".*\.spec\.ts$"}`:

**tests**/2.ts

An example of a **correct** file path for this rule configured as `{"allTestPattern": "__tests__", "pattern": ".*\.spec\.ts$"}`:

**tests**/2.spec.ts

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#configuration)

This rule accepts a configuration object with the following properties:

### allTestPattern [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#alltestpattern)

type: `string`

Regex pattern to ensure we are linting only test filenames. Decides whether a file is a testing file.

### pattern [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#pattern)

type: `string`

Required regex to check if a test filename have a valid formart. Pattern doesn't have a default value, you must provide one.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/consistent-test-filename": "error"
  }
}
```

bash

```
oxlint --deny vitest/consistent-test-filename --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/consistent_test_filename.rs)
