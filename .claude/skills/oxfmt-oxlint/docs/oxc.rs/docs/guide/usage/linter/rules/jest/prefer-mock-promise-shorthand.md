---
title: "jest/prefer-mock-promise-shorthand | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.md for this page in Markdown format

# jest/prefer-mock-promise-shorthand Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#jest-prefer-mock-promise-shorthand)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#what-it-does)

When working with mocks of functions that return promises, Jest provides some API sugar functions to reduce the amount of boilerplate you have to write. These methods should be preferred when possible.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#why-is-this-bad)

Using generic mock functions like `mockImplementation(() => Promise.resolve())` or `mockReturnValue(Promise.reject())` is more verbose and less readable than Jest's specialized promise shorthands. The shorthand methods like `mockResolvedValue()` and `mockRejectedValue()` are more expressive and make the test intent clearer.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
jest.fn().mockImplementation(() => Promise.resolve(123));
jest.spyOn(fs.promises, "readFile").mockReturnValue(Promise.reject(new Error("oh noes!")));

myFunction
  .mockReturnValueOnce(Promise.resolve(42))
  .mockImplementationOnce(() => Promise.resolve(42))
  .mockReturnValue(Promise.reject(new Error("too many calls!")));
```

Examples of **correct** code for this rule:

javascript

```
jest.fn().mockResolvedValue(123);
jest.spyOn(fs.promises, "readFile").mockRejectedValue(new Error("oh noes!"));

myFunction
  .mockResolvedValueOnce(42)
  .mockResolvedValueOnce(42)
  .mockRejectedValue(new Error("too many calls!"));
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-mock-promise-shorthand.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-mock-promise-shorthand": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-mock-promise-shorthand": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-mock-promise-shorthand --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_mock_promise_shorthand.rs)
