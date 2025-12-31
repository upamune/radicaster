---
title: "jest/prefer-to-have-been-called | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.md for this page in Markdown format

# jest/prefer-to-have-been-called Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#jest-prefer-to-have-been-called)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#what-it-does)

Suggests using `toHaveBeenCalled()` or `not.toHaveBeenCalled()` over `toHaveBeenCalledTimes(0)` or `toBeCalledTimes(0)`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#why-is-this-bad)

`toHaveBeenCalled()` is more explicit and readable than `toHaveBeenCalledTimes(0)`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#examples)

Examples of **incorrect** code for this rule:

js

```
expect(mock).toHaveBeenCalledTimes(0);
expect(mock).toBeCalledTimes(0);
expect(mock).not.toHaveBeenCalledTimes(0);
```

Examples of **correct** code for this rule:

js

```
expect(mock).not.toHaveBeenCalled();
expect(mock).toHaveBeenCalled();
expect(mock).toHaveBeenCalledTimes(1);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-to-have-been-called": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-to-have-been-called --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_to_have_been_called.rs)
