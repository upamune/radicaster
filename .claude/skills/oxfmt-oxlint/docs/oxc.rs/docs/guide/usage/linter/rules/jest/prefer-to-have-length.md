---
title: "jest/prefer-to-have-length | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-to-have-length.md for this page in Markdown format

# jest/prefer-to-have-length Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#jest-prefer-to-have-length)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#what-it-does)

In order to have a better failure message, `toHaveLength()` should be used upon asserting expectations on objects length property.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#why-is-this-bad)

This rule triggers a warning if `toBe()`, `toEqual()` or `toStrictEqual()` is used to assert objects length property.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(files["length"]).toBe(1);
expect(files["length"]).toBe(1);
expect(files["length"])["not"].toBe(1);
```

Examples of **correct** code for this rule:

javascript

```
expect(files).toHaveLength(1);
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-to-have-length.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-to-have-length": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-to-have-length": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-to-have-length --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_to_have_length.rs)
