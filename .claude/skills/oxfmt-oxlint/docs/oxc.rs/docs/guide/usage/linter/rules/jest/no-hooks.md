---
title: "jest/no-hooks | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-hooks.md for this page in Markdown format

# jest/no-hooks Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#jest-no-hooks)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#what-it-does)

Disallows Jest setup and teardown hooks, such as `beforeAll`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#why-is-this-bad)

Jest provides global functions for setup and teardown tasks, which are called before/after each test case and each test suite. The use of these hooks promotes shared state between tests.

This rule reports for the following function calls:

* `beforeAll`
* `beforeEach`
* `afterAll`
* `afterEach`

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
function setupFoo(options) {
  /* ... */
}
function setupBar(options) {
  /* ... */
}

describe("foo", () => {
  let foo;
  beforeEach(() => {
    foo = setupFoo();
  });
  afterEach(() => {
    foo = null;
  });
  it("does something", () => {
    expect(foo.doesSomething()).toBe(true);
  });
  describe("with bar", () => {
    let bar;
    beforeEach(() => {
      bar = setupBar();
    });
    afterEach(() => {
      bar = null;
    });
    it("does something with bar", () => {
      expect(foo.doesSomething(bar)).toBe(true);
    });
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-hooks.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-hooks": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#allow)

type: `string[]`

default: `[]`

An array of hook function names that are permitted for use.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-hooks": "error"
  }
}
```

bash

```
oxlint --deny jest/no-hooks --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_hooks.rs)
