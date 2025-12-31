---
title: "jest/valid-expect | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/valid-expect.md for this page in Markdown format

# jest/valid-expect Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#jest-valid-expect)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#what-it-does)

Checks that `expect()` is called correctly.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#why-is-this-bad)

`expect()` is a function that is used to assert values in tests. It should be called with a single argument, which is the value to be tested. If you call `expect()` with no arguments, or with more than one argument, it will not work as expected.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect();
expect("something");
expect(true).toBeDefined;
expect(Promise.resolve("Hi!")).resolves.toBe("Hi!");
```

Examples of **correct** code for this rule:

javascript

```
expect("something").toEqual("something");
expect(true).toBeDefined();
expect(Promise.resolve("Hi!")).resolves.toBe("Hi!");
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/v1.1.9/docs/rules/valid-expect.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/valid-expect": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#configuration)

This rule accepts a configuration object with the following properties:

### alwaysAwait [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#alwaysawait)

type: `boolean`

default: `false`

When `true`, async assertions must be awaited in all contexts (not just return statements).

### asyncMatchers [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#asyncmatchers)

type: `string[]`

default: `["toResolve", "toReject"]`

List of matchers that are considered async and therefore require awaiting (e.g. `toResolve`, `toReject`).

### maxArgs [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#maxargs)

type: `integer`

default: `1`

Maximum number of arguments `expect` should be called with.

### minArgs [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#minargs)

type: `integer`

default: `1`

Minimum number of arguments `expect` should be called with.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/valid-expect": "error"
  }
}
```

bash

```
oxlint --deny jest/valid-expect --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/valid_expect.rs)
