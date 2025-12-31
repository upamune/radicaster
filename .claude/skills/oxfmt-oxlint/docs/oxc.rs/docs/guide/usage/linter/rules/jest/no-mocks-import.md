---
title: "jest/no-mocks-import | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/no-mocks-import.md for this page in Markdown format

# jest/no-mocks-import Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#jest-no-mocks-import)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#what-it-does)

This rule reports imports from a path containing a **mocks** component.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#why-is-this-bad)

Manually importing mocks from a `__mocks__` directory can lead to unexpected behavior and breaks Jest's automatic mocking system. Jest is designed to automatically resolve and use mocks from `__mocks__` directories when `jest.mock()` is called. Directly importing from these directories bypasses Jest's module resolution system and can cause inconsistencies between test and production environments.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#examples)

Examples of **incorrect** code for this rule:

ts

```
import thing from "./__mocks__/index";
require("./__mocks__/index");
```

Examples of **correct** code for this rule:

ts

```
import thing from "thing";
require("thing");
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/no-mocks-import.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/no-mocks-import": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/no-mocks-import": "error"
  }
}
```

bash

```
oxlint --deny jest/no-mocks-import --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/no_mocks_import.rs)
