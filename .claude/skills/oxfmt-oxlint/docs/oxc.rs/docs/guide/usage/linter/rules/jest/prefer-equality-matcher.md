---
title: "jest/prefer-equality-matcher | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-equality-matcher.md for this page in Markdown format

# jest/prefer-equality-matcher Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#jest-prefer-equality-matcher)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#what-it-does)

Jest has built-in matchers for expecting equality, which allow for more readable tests and error messages if an expectation fails.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#why-is-this-bad)

Testing equality expressions with generic matchers like `toBe(true)` makes tests harder to read and understand. When tests fail, the error messages are less helpful because they don't show what the actual values were. Using specific equality matchers provides clearer test intent and better debugging information.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(x === 5).toBe(true);
expect(name === "Carl").not.toEqual(true);
expect(myObj !== thatObj).toStrictEqual(true);
```

Examples of **correct** code for this rule:

javascript

```
expect(x).toBe(5);
expect(name).not.toEqual("Carl");
expect(myObj).toStrictEqual(thatObj);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-equality-matcher.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-equality-matcher": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-equality-matcher": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-equality-matcher --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_equality_matcher.rs)
