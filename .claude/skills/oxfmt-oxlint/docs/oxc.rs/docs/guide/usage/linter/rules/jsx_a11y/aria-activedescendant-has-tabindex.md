---
title: "jsx_a11y/aria-activedescendant-has-tabindex | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/aria-activedescendant-has-tabindex.md for this page in Markdown format

# jsx\_a11y/aria-activedescendant-has-tabindex Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#jsx-a11y-aria-activedescendant-has-tabindex)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#what-it-does)

Enforce elements with aria-activedescendant are tabbable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#why-is-this-bad)

Elements with `aria-activedescendant` must be tabbable for users to navigate to them using keyboard input. Without proper tabindex, screen reader users cannot access the element through keyboard navigation, making the functionality inaccessible.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const Bad = <div aria-activedescendant={someID} />;
```

Examples of **correct** code for this rule:

jsx

```
const Good = (
  <>
    <CustomComponent />
    <CustomComponent aria-activedescendant={someID} />
    <CustomComponent aria-activedescendant={someID} tabIndex={0} />
    <CustomComponent aria-activedescendant={someID} tabIndex={-1} />
    <div />
    <input />
    <div tabIndex={0} />
    <div aria-activedescendant={someID} tabIndex={0} />
    <div aria-activedescendant={someID} tabIndex="0" />
    <div aria-activedescendant={someID} tabIndex={1} />
    <div aria-activedescendant={someID} tabIndex={-1} />
    <div aria-activedescendant={someID} tabIndex="-1" />
    <input aria-activedescendant={someID} />
    <input aria-activedescendant={someID} tabIndex={0} />
    <input aria-activedescendant={someID} tabIndex={-1} />
  </>
);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/aria-activedescendant-has-tabindex": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/aria-activedescendant-has-tabindex --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/aria_activedescendant_has_tabindex.rs)
