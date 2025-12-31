---
title: "jest/no-conditional-expect | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-conditional-expect.md for this page in Markdown format

# jest/no-conditional-expect Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#jest-no-conditional-expect)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#what-it-does)

This rule prevents the use of expect in conditional blocks, such as ifs & catch(s). This includes using expect in callbacks to functions named catch, which are assumed to be promises.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#why-is-this-bad)

Jest only considers a test to have failed if it throws an error, meaning if calls to assertion functions like expect occur in conditional code such as a catch statement, tests can end up passing but not actually test anything. Additionally, conditionals tend to make tests more brittle and complex, as they increase the amount of mental thinking needed to understand what is actually being tested.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#examples)

Examples of **incorrect** code for this rule:

js

```
it("foo", () => {
  doTest && expect(1).toBe(2);
});

it("bar", () => {
  if (!skipTest) {
    expect(1).toEqual(2);
  }
});

it("baz", async () => {
  try {
    await foo();
  } catch (err) {
    expect(err).toMatchObject({ code: "MODULE_NOT_FOUND" });
  }
});

it("throws an error", async () => {
  await foo().catch((error) => expect(error).toBeInstanceOf(error));
});
```

Examples of **correct** code for this rule:

js

```
it("foo", () => {
  expect(!value).toBe(false);
});

function getValue() {
  if (process.env.FAIL) {
    return 1;
  }
  return 2;
}

it("foo", () => {
  expect(getValue()).toBe(2);
});

it("validates the request", () => {
  try {
    processRequest(request);
  } catch {
  } finally {
    expect(validRequest).toHaveBeenCalledWith(request);
  }
});

it("throws an error", async () => {
  await expect(foo).rejects.toThrow(Error);
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-conditional-expect.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-conditional-expect": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-conditional-expect": "error"
  }
}
```

bash

```
oxlint --deny jest/no-conditional-expect --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_conditional_expect.rs)
