---
title: "jest/no-standalone-expect | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-standalone-expect.md for this page in Markdown format

# jest/no-standalone-expect Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#jest-no-standalone-expect)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#what-it-does)

Prevents `expect` statements outside of a `test` or `it` block. An `expect` within a helper function (but outside of a `test` or `it` block) will not trigger this rule.

Statements like `expect.hasAssertions()` will NOT trigger this rule since these calls will execute if they are not in a test block.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#why-is-this-bad)

`expect` statements outside of test blocks will not be executed by the Jest test runner, which means they won't actually test anything. This can lead to false confidence in test coverage and may hide bugs that would otherwise be caught by proper testing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
describe("a test", () => {
  expect(1).toBe(1);
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-standalone-expect.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-standalone-expect": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#configuration)

This rule accepts a configuration object with the following properties:

### additionalTestBlockFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#additionaltestblockfunctions)

type: `string[]`

default: `[]`

An array of function names that should also be treated as test blocks.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-standalone-expect": "error"
  }
}
```

bash

```
oxlint --deny jest/no-standalone-expect --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_standalone_expect/mod.rs)
