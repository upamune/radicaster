---
title: "jsx_a11y/role-supports-aria-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/role-supports-aria-props.md for this page in Markdown format

# jsx\_a11y/role-supports-aria-props Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html#jsx-a11y-role-supports-aria-props)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html#what-it-does)

Enforce that elements with explicit or implicit roles defined contain only `aria-*` properties supported by that `role`. Many ARIA attributes (states and properties) can only be used on elements with particular roles. Some elements have implicit roles, such as `<a href="#" />`, which will resolve to `role="link"`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<ul role="radiogroup" "aria-labelledby"="foo">
    <li aria-required tabIndex="-1" role="radio" aria-checked="false">Rainbow Trout</li>
    <li aria-required tabIndex="-1" role="radio" aria-checked="false">Brook Trout</li>
    <li aria-required tabIndex="0" role="radio" aria-checked="true">Lake Trout</li>
</ul>
```

Examples of **correct** code for this rule:

jsx

```
<ul role="radiogroup" aria-required "aria-labelledby"="foo">
    <li tabIndex="-1" role="radio" aria-checked="false">Rainbow Trout</li>
    <li tabIndex="-1" role="radio" aria-checked="false">Brook Trout</li>
    <li tabIndex="0" role="radio" aria-checked="true">Lake Trout</li>
</ul>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/role-supports-aria-props": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/role-supports-aria-props --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/role_supports_aria_props.rs)
