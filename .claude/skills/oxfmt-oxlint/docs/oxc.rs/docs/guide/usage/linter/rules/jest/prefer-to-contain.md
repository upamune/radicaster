---
title: "jest/prefer-to-contain | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-to-contain.md for this page in Markdown format

# jest/prefer-to-contain Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#jest-prefer-to-contain)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#what-it-does)

In order to have a better failure message, `toContain()` should be used upon asserting expectations on an array containing an object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#why-is-this-bad)

This rule triggers a warning if `toBe()`, `toEqual()` or `toStrictEqual()` is used to assert object inclusion in an array

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(a.includes(b)).toBe(true);
expect(a.includes(b)).not.toBe(true);
expect(a.includes(b)).toBe(false);
expect(a.includes(b)).toEqual(true);
expect(a.includes(b)).toStrictEqual(true);
```

Examples of **correct** code for this rule:

javascript

```
expect(a).toContain(b);
expect(a).not.toContain(b);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-to-contain.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-to-contain": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-to-contain": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-to-contain --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_to_contain.rs)
