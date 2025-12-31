---
title: "jsx_a11y/autocomplete-valid | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/autocomplete-valid.md for this page in Markdown format

# jsx\_a11y/autocomplete-valid Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#jsx-a11y-autocomplete-valid)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#what-it-does)

Enforces that an element's autocomplete attribute must be a valid value.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#why-is-this-bad)

Incorrectly using the autocomplete attribute may decrease the accessibility of the website for users.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<input autocomplete="invalid-value" />
```

Examples of **correct** code for this rule:

jsx

```
<input autocomplete="name" />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#configuration)

This rule accepts a configuration object with the following properties:

### inputComponents [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#inputcomponents)

type: `string[]`

default: `["input"]`

List of custom component names that should be treated as input elements.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/autocomplete-valid": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/autocomplete-valid --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/autocomplete_valid.rs)
