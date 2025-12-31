---
title: "jest/no-restricted-jest-methods | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.md for this page in Markdown format

# jest/no-restricted-jest-methods Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#jest-no-restricted-jest-methods)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#what-it-does)

Restrict the use of specific `jest` and `vi` methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#why-is-this-bad)

Certain Jest or Vitest methods may be deprecated, discouraged in specific contexts, or incompatible with your testing environment. Restricting them helps maintain consistent and reliable test practices.

By default, no methods are restricted by this rule. You must configure the rule for it to disable anything.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
jest.useFakeTimers();
it("calls the callback after 1 second via advanceTimersByTime", () => {
  // ...

  jest.advanceTimersByTime(1000);

  // ...
});

test("plays video", () => {
  const spy = jest.spyOn(video, "play");

  // ...
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-restricted-vi-methods.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-restricted-vi-methods": [
      "error",
      { "badFunction": "Don't use `badFunction`, it is bad." }
    ]
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#configuration)

This rule accepts a configuration object with the following properties:

### restrictedJestMethods [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#restrictedjestmethods)

type: `Record<string, string>`

default: `{}`

A mapping of restricted Jest method names to custom messages - or `null`, for a generic message.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-restricted-jest-methods": "error"
  }
}
```

bash

```
oxlint --deny jest/no-restricted-jest-methods --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_restricted_jest_methods.rs)
