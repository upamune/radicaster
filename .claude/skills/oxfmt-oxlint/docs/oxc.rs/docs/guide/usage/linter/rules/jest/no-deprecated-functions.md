---
title: "jest/no-deprecated-functions | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-deprecated-functions.md for this page in Markdown format

# jest/no-deprecated-functions Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-no-deprecated-functions)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#what-it-does)

Over the years Jest has accrued some debt in the form of functions that have either been renamed for clarity, or replaced with more powerful APIs.

This rule can also autofix a number of these deprecations for you.

#### `jest.resetModuleRegistry` [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-resetmoduleregistry)

This function was renamed to `resetModules` in Jest 15 and removed in Jest 27.

#### `jest.addMatchers` [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-addmatchers)

This function was replaced with `expect.extend` in Jest 17 and removed in Jest 27.

#### `require.requireActual` & `require.requireMock` [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#require-requireactual-require-requiremock)

These functions were replaced in Jest 21 and removed in Jest 26.

Originally, the `requireActual` and `requireMock` functions were placed onto the `require` function.

These functions were later moved onto the `jest` object in order to be easier for type checkers to handle, and their use via `require` deprecated. Finally, the release of Jest 26 saw them removed from the `require` function altogether.

#### `jest.runTimersToTime` [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-runtimerstotime)

This function was renamed to `advanceTimersByTime` in Jest 22 and removed in Jest 27.

#### `jest.genMockFromModule` [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-genmockfrommodule)

This function was renamed to `createMockFromModule` in Jest 26, and is scheduled for removal in Jest 30.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#why-is-this-bad)

While typically these deprecated functions are kept in the codebase for a number of majors, eventually they are removed completely.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
jest.resetModuleRegistry; // since Jest 15
jest.addMatchers; // since Jest 17
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#configuration)

This rule accepts a configuration object with the following properties:

### jest [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest)

type: `object`

Jest configuration options.

#### jest.version [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#jest-version)

type: `string`

default: `"29"`

The version of Jest being used.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-deprecated-functions": "error"
  }
}
```

bash

```
oxlint --deny jest/no-deprecated-functions --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_deprecated_functions.rs)
