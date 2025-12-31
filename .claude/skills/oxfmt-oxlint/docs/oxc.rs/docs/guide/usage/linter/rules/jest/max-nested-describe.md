---
title: "jest/max-nested-describe | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/max-nested-describe.md for this page in Markdown format

# jest/max-nested-describe Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#jest-max-nested-describe)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#what-it-does)

This rule enforces a maximum depth to nested `describe()` calls.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#why-is-this-bad)

Nesting `describe()` blocks too deeply can make the test suite hard to read and understand.

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#example)

The following patterns are considered warnings (with the default option of `{ "max": 5 }` ):

Examples of **incorrect** code for this rule:

javascript

```
describe("foo", () => {
  describe("bar", () => {
    describe("baz", () => {
      describe("qux", () => {
        describe("quxx", () => {
          describe("too many", () => {
            it("should get something", () => {
              expect(getSomething()).toBe("Something");
            });
          });
        });
      });
    });
  });
});

describe("foo", function () {
  describe("bar", function () {
    describe("baz", function () {
      describe("qux", function () {
        describe("quxx", function () {
          describe("too many", function () {
            it("should get something", () => {
              expect(getSomething()).toBe("Something");
            });
          });
        });
      });
    });
  });
});
```

Examples of **correct** code for this rule:

ts

```
describe("foo", () => {
  describe("bar", () => {
    it("should get something", () => {
      expect(getSomething()).toBe("Something");
    });
  });
  describe("qux", () => {
    it("should get something", () => {
      expect(getSomething()).toBe("Something");
    });
  });
});

describe("foo2", function () {
  it("should get something", () => {
    expect(getSomething()).toBe("Something");
  });
});

describe("foo", function () {
  describe("bar", function () {
    describe("baz", function () {
      describe("qux", function () {
        describe("this is the limit", function () {
          it("should get something", () => {
            expect(getSomething()).toBe("Something");
          });
        });
      });
    });
  });
});
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/max-nested-describe.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/max-nested-describe": "error"
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#max)

type: `integer`

default: `5`

Maximum allowed depth of nested describe calls.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/max-nested-describe": "error"
  }
}
```

bash

```
oxlint --deny jest/max-nested-describe --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/max_nested_describe.rs)
