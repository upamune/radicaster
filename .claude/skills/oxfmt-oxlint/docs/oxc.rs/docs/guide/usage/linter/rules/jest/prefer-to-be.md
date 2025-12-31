---
title: "jest/prefer-to-be | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-to-be.md for this page in Markdown format

# jest/prefer-to-be Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#jest-prefer-to-be)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#what-it-does)

Recommends using `toBe` matcher for primitive literals and specific matchers for `null`, `undefined`, and `NaN`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#why-is-this-bad)

When asserting against primitive literals such as numbers and strings, the equality matchers all operate the same, but read slightly differently in code.

This rule recommends using the `toBe` matcher in these situations, as it forms the most grammatically natural sentence. For `null`, `undefined`, and `NaN` this rule recommends using their specific `toBe` matchers, as they give better error messages as well.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(value).not.toEqual(5);
expect(getMessage()).toStrictEqual("hello world");
expect(loadMessage()).resolves.toEqual("hello world");
```

Examples of **correct** code for this rule:

javascript

```
expect(value).not.toBe(5);
expect(getMessage()).toBe("hello world");
expect(loadMessage()).resolves.toBe("hello world");
expect(didError).not.toBe(true);
expect(catchError()).toStrictEqual({ message: "oh noes!" });
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-to-be.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-to-be": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-to-be": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-to-be --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_to_be.rs)
