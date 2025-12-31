---
title: "jsx_a11y/no-access-key | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-access-key.md for this page in Markdown format

# jsx\_a11y/no-access-key Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#jsx-a11y-no-access-key)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#what-it-does)

Enforces that the `accessKey` prop is not used on any element to avoid complications with keyboard commands used by a screenreader.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#why-is-this-bad)

Access keys are HTML attributes that allow web developers to assign keyboard shortcuts to elements. Inconsistencies between keyboard shortcuts and keyboard commands used by screenreaders and keyboard-only users create accessibility complications so to avoid complications, access keys should not be used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div accessKey="h" />
```

Examples of **correct** code for this rule:

jsx

```
<div />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-access-key": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-access-key --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_access_key.rs)
