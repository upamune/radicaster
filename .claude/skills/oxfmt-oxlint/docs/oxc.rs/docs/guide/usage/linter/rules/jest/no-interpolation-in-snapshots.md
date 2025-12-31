---
title: "jest/no-interpolation-in-snapshots | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.md for this page in Markdown format

# jest/no-interpolation-in-snapshots Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#jest-no-interpolation-in-snapshots)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#what-it-does)

Prevents the use of string interpolations in snapshots.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#why-is-this-bad)

Interpolation prevents snapshots from being updated. Instead, properties should be overloaded with a matcher by using [property matchers](https://jestjs.io/docs/en/snapshot-testing#property-matchers).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(something).toMatchInlineSnapshot(
  `Object {
    property: ${interpolated}
  }`,
);

expect(something).toMatchInlineSnapshot(
  { other: expect.any(Number) },
  `Object {
    other: Any<Number>,
    property: ${interpolated}
  }`,
);

expect(errorThrowingFunction).toThrowErrorMatchingInlineSnapshot(`${interpolated}`);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-interpolation-in-snapshots.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-interpolation-in-snapshots": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-interpolation-in-snapshots": "error"
  }
}
```

bash

```
oxlint --deny jest/no-interpolation-in-snapshots --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_interpolation_in_snapshots.rs)
