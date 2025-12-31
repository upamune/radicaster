---
title: "jest/prefer-called-with | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-called-with.md for this page in Markdown format

# jest/prefer-called-with Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#jest-prefer-called-with)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#what-it-does)

Suggest using `toBeCalledWith()` or `toHaveBeenCalledWith()`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#why-is-this-bad)

When testing function calls, it's often more valuable to assert both that a function was called AND what arguments it was called with. Using `toBeCalled()` or `toHaveBeenCalled()` only verifies the function was invoked, but doesn't validate the arguments, potentially missing bugs where functions are called with incorrect parameters.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
expect(someFunction).toBeCalled();
expect(someFunction).toHaveBeenCalled();
```

Examples of **correct** code for this rule:

javascript

```
expect(noArgsFunction).toBeCalledWith();
expect(roughArgsFunction).toBeCalledWith(expect.anything(), expect.any(Date));
expect(anyArgsFunction).toBeCalledTimes(1);
expect(uncalledFunction).not.toBeCalled();
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-called-with.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-called-with": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-called-with": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-called-with --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_called_with.rs)
