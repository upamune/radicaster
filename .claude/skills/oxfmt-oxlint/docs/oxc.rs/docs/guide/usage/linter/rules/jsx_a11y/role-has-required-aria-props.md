---
title: "jsx_a11y/role-has-required-aria-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/role-has-required-aria-props.md for this page in Markdown format

# jsx\_a11y/role-has-required-aria-props Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#jsx-a11y-role-has-required-aria-props)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#what-it-does)

Enforces that elements with ARIA roles must have all required attributes for that role.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#why-is-this-bad)

Certain ARIA roles require specific attributes to express necessary semantics for assistive technology.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div role="checkbox" />
```

Examples of **correct** code for this rule:

jsx

```
<div role="checkbox" aria-checked="false" />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/role-has-required-aria-props": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/role-has-required-aria-props --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/role_has_required_aria_props.rs)
