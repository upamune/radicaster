---
title: "jest/prefer-expect-resolves | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-expect-resolves.md for this page in Markdown format

# jest/prefer-expect-resolves Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#jest-prefer-expect-resolves)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#what-it-does)

Prefer `await expect(...).resolves` over `expect(await ...)` when testing promises.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#why-is-this-bad)

When working with promises, there are two primary ways you can test the resolved value:

1. use the `resolve` modifier on `expect` (`await expect(...).resolves.<matcher>` style)
2. `await` the promise and assert against its result (`expect(await ...).<matcher>` style)

While the second style is arguably less dependent on `jest`, if the promise rejects it will be treated as a general error, resulting in less predictable behaviour and output from `jest`.

Additionally, favoring the first style ensures consistency with its `rejects` counterpart, as there is no way of "awaiting" a rejection.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
it("passes", async () => {
  expect(await someValue()).toBe(true);
});
it("is true", async () => {
  const myPromise = Promise.resolve(true);
  expect(await myPromise).toBe(true);
});
```

Examples of **correct** code for this rule:

javascript

```
it("passes", async () => {
  await expect(someValue()).resolves.toBe(true);
});
it("is true", async () => {
  const myPromise = Promise.resolve(true);

  await expect(myPromise).resolves.toBe(true);
});
it("errors", async () => {
  await expect(Promise.reject(new Error("oh noes!"))).rejects.toThrowError("oh noes!");
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-expect-resolves.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-expect-resolves": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-expect-resolves": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-expect-resolves --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_expect_resolves.rs)
