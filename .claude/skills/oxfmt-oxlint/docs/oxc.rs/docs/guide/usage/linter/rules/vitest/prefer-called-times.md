---
title: "vitest/prefer-called-times | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vitest/prefer-called-times.md for this page in Markdown format

# vitest/prefer-called-times Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#vitest-prefer-called-times)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#what-it-does)

This rule aims to enforce the use of `toBeCalledTimes(1)` or `toHaveBeenCalledTimes(1)` over `toBeCalledOnce()` or `toHaveBeenCalledOnce()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#why-is-this-bad)

This rule aims to enforce the use of `toBeCalledTimes(1)` or `toHaveBeenCalledTimes(1)` over `toBeCalledOnce()` or `toHaveBeenCalledOnce()`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#examples)

Examples of **incorrect** code for this rule:

js

```
test("foo", () => {
  const mock = vi.fn();
  mock("foo");
  expect(mock).toBeCalledOnce();
  expect(mock).toHaveBeenCalledOnce();
});
```

Examples of **correct** code for this rule:

js

```
test("foo", () => {
  const mock = vi.fn();
  mock("foo");
  expect(mock).toBeCalledTimes(1);
  expect(mock).toHaveBeenCalledTimes(1);
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vitest"],
  "rules": {
    "vitest/prefer-called-times": "error"
  }
}
```

bash

```
oxlint --deny vitest/prefer-called-times --vitest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vitest/prefer_called_times.rs)
