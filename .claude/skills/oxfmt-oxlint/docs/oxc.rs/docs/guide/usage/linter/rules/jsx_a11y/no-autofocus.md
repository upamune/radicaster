---
title: "jsx_a11y/no-autofocus | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-autofocus.md for this page in Markdown format

# jsx\_a11y/no-autofocus Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#jsx-a11y-no-autofocus)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#what-it-does)

Enforce that `autoFocus` prop is not used on elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#why-is-this-bad)

Autofocusing elements can cause usability issues for sighted and non-sighted users alike. It can be disorienting when focus is shifted without user input and can interfere with assistive technologies. Users should control when and where focus moves on a page.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div autoFocus />
<div autoFocus="true" />
<div autoFocus="false" />
<div autoFocus={undefined} />
```

Examples of **correct** code for this rule:

jsx

```
<div />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreNonDOM [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#ignorenondom)

type: `boolean`

default: `false`

Determines if developer-created components are checked.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-autofocus": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-autofocus --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_autofocus.rs)
