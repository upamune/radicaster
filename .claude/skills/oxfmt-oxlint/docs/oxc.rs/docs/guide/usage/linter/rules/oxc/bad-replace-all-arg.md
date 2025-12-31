---
title: "oxc/bad-replace-all-arg | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.md for this page in Markdown format

# oxc/bad-replace-all-arg Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#oxc-bad-replace-all-arg)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#what-it-does)

This rule warns when the `replaceAll` method is called with a regular expression that does not have the global flag (g).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#why-is-this-bad)

The `replaceAll` method replaces all occurrences of a string with another string. If the global flag (g) is not used in the regular expression, only the first occurrence of the string will be replaced.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
withSpaces.replaceAll(/\s+/, ",");
```

Examples of **correct** code for this rule:

javascript

```
withSpaces.replaceAll(/\s+/g, ",");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/bad-replace-all-arg": "error"
  }
}
```

bash

```
oxlint --deny oxc/bad-replace-all-arg
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/bad_replace_all_arg.rs)
