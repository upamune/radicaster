---
title: "jest/expect-expect | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/expect-expect.md for this page in Markdown format

# jest/expect-expect Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#jest-expect-expect)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#what-it-does)

This rule triggers when there is no call made to `expect` in a test, ensure that there is at least one `expect` call made in a test.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#why-is-this-bad)

People may forget to add assertions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
it("should be a test", () => {
  console.log("no assertion");
});
test("should assert something", () => {});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/expect-expect.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/expect-expect": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#configuration)

This rule accepts a configuration object with the following properties:

### additionalTestBlockFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#additionaltestblockfunctions)

type: `string[]`

default: `[]`

An array of function names that should also be treated as test blocks.

### assertFunctionNames [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#assertfunctionnames)

type: `string[]`

default: `["expect"]`

A list of function names that should be treated as assertion functions.

NOTE: The default value is `["expect"]` for Jest and `["expect", "expectTypeOf", "assert", "assertType"]` for Vitest.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/expect-expect": "error"
  }
}
```

bash

```
oxlint --deny jest/expect-expect --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/expect_expect.rs)
