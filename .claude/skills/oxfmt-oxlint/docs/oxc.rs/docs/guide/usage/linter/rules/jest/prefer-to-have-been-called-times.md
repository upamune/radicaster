---
title: "jest/prefer-to-have-been-called-times | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.md for this page in Markdown format

# jest/prefer-to-have-been-called-times Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#jest-prefer-to-have-been-called-times)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#what-it-does)

In order to have a better failure message, [`toHaveBeenCalledTimes` should be used instead of directly checking the length of `mock.calls`](https://github.com/jest-community/eslint-plugin-jest/blob/v29.5.0/docs/rules/prefer-to-have-been-called-times.md).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#why-is-this-bad)

This rule triggers a warning if `toHaveLength` is used to assert the number of times a mock is called.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#examples)

Examples of **incorrect** code for this rule:

js

```
expect(someFunction.mock.calls).toHaveLength(1);
expect(someFunction.mock.calls).toHaveLength(0);
expect(someFunction.mock.calls).not.toHaveLength(1);
```

Examples of **correct** code for this rule:

js

```
expect(someFunction).toHaveBeenCalledTimes(1);
expect(someFunction).toHaveBeenCalledTimes(0);
expect(someFunction).not.toHaveBeenCalledTimes(0);
expect(uncalledFunction).not.toBeCalled();
expect(method.mock.calls[0][0]).toStrictEqual(value);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-to-have-been-called-times": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-to-have-been-called-times --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_to_have_been_called_times.rs)
