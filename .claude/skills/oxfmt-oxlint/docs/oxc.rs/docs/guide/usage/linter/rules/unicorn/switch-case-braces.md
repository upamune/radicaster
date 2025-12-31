---
title: "unicorn/switch-case-braces | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/switch-case-braces.md for this page in Markdown format

# unicorn/switch-case-braces Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#unicorn-switch-case-braces)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#what-it-does)

Requires empty switch cases to omit braces, while non-empty cases must use braces. This reduces visual clutter for empty cases and enforces proper scoping for non-empty ones.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#why-is-this-bad)

Using braces unnecessarily for empty cases adds visual noise, while omitting braces in non-empty cases can lead to scoping issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
switch (num) {
  case 1: {
  }
  case 2:
    console.log("Case 2");
    break;
}
```

Examples of **correct** code for this rule:

javascript

```
switch (num) {
  case 1:
  case 2: {
    console.log("Case 2");
    break;
  }
}
```

Example config:

json

```
"unicorn/switch-case-braces": ["error", "avoid"]
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#always)

Always require braces in case clauses (except empty cases).

### `"avoid"` [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#avoid)

Allow braces only when needed for scoping (e.g., variable or function declarations).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/switch-case-braces": "error"
  }
}
```

bash

```
oxlint --deny unicorn/switch-case-braces
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/switch_case_braces.rs)
