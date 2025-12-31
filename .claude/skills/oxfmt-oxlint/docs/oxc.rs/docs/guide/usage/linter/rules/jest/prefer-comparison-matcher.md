---
title: "jest/prefer-comparison-matcher | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.md for this page in Markdown format

# jest/prefer-comparison-matcher Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#jest-prefer-comparison-matcher)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#what-it-does)

This rule checks for comparisons in tests that could be replaced with one of the following built-in comparison matchers:

* `toBeGreaterThan`
* `toBeGreaterThanOrEqual`
* `toBeLessThan`
* `toBeLessThanOrEqual`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#why-is-this-bad)

Using generic matchers like `toBe(true)` with comparison expressions makes tests less readable and provides less helpful error messages when they fail. Jest's specific comparison matchers offer clearer intent and better error output that shows the actual values being compared.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#examples)

Examples of **incorrect** code for this rule:

js

```
expect(x > 5).toBe(true);
expect(x < 7).not.toEqual(true);
expect(x <= y).toStrictEqual(true);
```

Examples of **correct** code for this rule:

js

```
expect(x).toBeGreaterThan(5);
expect(x).not.toBeLessThanOrEqual(7);
expect(x).toBeLessThanOrEqual(y);
// special case - see below
expect(x < "Carl").toBe(true);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-comparison-matcher.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-comparison-matcher": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-comparison-matcher": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-comparison-matcher --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_comparison_matcher.rs)
