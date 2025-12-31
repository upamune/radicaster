---
title: "jsx_a11y/no-redundant-roles | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-redundant-roles.md for this page in Markdown format

# jsx\_a11y/no-redundant-roles Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#jsx-a11y-no-redundant-roles)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#what-it-does)

Enforces that the explicit `role` property is not the same as implicit/default role property on element.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#why-is-this-bad)

Redundant roles can lead to confusion and verbosity in the codebase.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<nav role="navigation" />
```

Examples of **correct** code for this rule:

jsx

```
<nav />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-redundant-roles": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-redundant-roles --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_redundant_roles.rs)
