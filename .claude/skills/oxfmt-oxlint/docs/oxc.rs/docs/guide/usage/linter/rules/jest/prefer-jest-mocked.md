---
title: "jest/prefer-jest-mocked | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-jest-mocked.md for this page in Markdown format

# jest/prefer-jest-mocked Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#jest-prefer-jest-mocked)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#what-it-does)

When working with mocks of functions using Jest, it's recommended to use the `jest.mocked()` helper function to properly type the mocked functions. This rule enforces the use of `jest.mocked()` for better type safety and readability.

Restricted types:

* `jest.Mock`
* `jest.MockedFunction`
* `jest.MockedClass`
* `jest.MockedObject`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#why-is-this-bad)

Using type assertions like `fn as jest.Mock` is a less safe approach than using `jest.mocked()`. The `jest.mocked()` helper provides better type safety by preserving the original function signature while adding mock capabilities. It also makes the code more readable and explicit about mocking intentions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
(foo as jest.Mock).mockReturnValue(1);
const mock = (foo as jest.Mock).mockReturnValue(1);
(foo as unknown as jest.Mock).mockReturnValue(1);
(Obj.foo as jest.Mock).mockReturnValue(1);
([].foo as jest.Mock).mockReturnValue(1);
```

Examples of **correct** code for this rule:

typescript

```
jest.mocked(foo).mockReturnValue(1);
const mock = jest.mocked(foo).mockReturnValue(1);
jest.mocked(Obj.foo).mockReturnValue(1);
jest.mocked([].foo).mockReturnValue(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-jest-mocked": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-jest-mocked --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_jest_mocked.rs)
