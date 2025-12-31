---
title: "jsx_a11y/heading-has-content | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/heading-has-content.md for this page in Markdown format

# jsx\_a11y/heading-has-content Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#jsx-a11y-heading-has-content)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#what-it-does)

Enforce that heading elements (h1, h2, etc.) have content and that the content is accessible to screen readers. Accessible means that it is not hidden using the aria-hidden prop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#why-is-this-bad)

Screen readers alert users to the presence of a heading tag. If the heading is empty or the text cannot be accessed, this could either confuse users or even prevent them from accessing information on the page's structure.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<h1 />
```

Examples of **correct** code for this rule:

jsx

```
<h1>Foo</h1>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#configuration)

This rule accepts a configuration object with the following properties:

### components [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#components)

type: `string[]`

default: `null`

Additional custom component names to treat as heading elements. These will be validated in addition to the standard h1-h6 elements.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/heading-has-content": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/heading-has-content --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/heading_has_content.rs)
