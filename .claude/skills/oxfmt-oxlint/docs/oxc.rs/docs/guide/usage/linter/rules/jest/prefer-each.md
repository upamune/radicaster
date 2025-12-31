---
title: "jest/prefer-each | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-each.md for this page in Markdown format

# jest/prefer-each Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#jest-prefer-each)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#what-it-does)

This rule enforces using `each` rather than manual loops.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#why-is-this-bad)

Manual loops for tests can be less readable and more error-prone. Using `each` provides a clearer and more concise way to run parameterized tests, improving readability and maintainability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#examples)

Examples of **incorrect** code for this rule:

js

```
for (const item of items) {
  describe(item, () => {
    expect(item).toBe("foo");
  });
}
```

Examples of **correct** code for this rule:

js

```
describe.each(items)("item", (item) => {
  expect(item).toBe("foo");
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-each.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-each": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-each": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-each --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_each.rs)
