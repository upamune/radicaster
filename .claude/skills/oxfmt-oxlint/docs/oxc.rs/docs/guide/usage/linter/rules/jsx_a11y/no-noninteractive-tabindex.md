---
title: "jsx_a11y/no-noninteractive-tabindex | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/no-noninteractive-tabindex.md for this page in Markdown format

# jsx\_a11y/no-noninteractive-tabindex Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#jsx-a11y-no-noninteractive-tabindex)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#what-it-does)

This rule checks that non-interactive elements don't have a tabIndex which would make them interactive via keyboard navigation.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#why-is-this-bad)

Tab key navigation should be limited to elements on the page that can be interacted with. Thus it is not necessary to add a tabindex to items in an unordered list, for example, to make them navigable through assistive technology.

These applications already afford page traversal mechanisms based on the HTML of the page. Generally, we should try to reduce the size of the page's tab ring rather than increasing it.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div tabIndex="0" />
<div role="article" tabIndex="0" />
<article tabIndex="0" />
<article tabIndex={0} />
```

Examples of **correct** code for this rule:

jsx

```
<div />
<MyButton tabIndex={0} />
<button />
<button tabIndex="0" />
<button tabIndex={0} />
<div />
<div tabIndex="-1" />
<div role="button" tabIndex="0" />
<div role="article" tabIndex="-1" />
<article tabIndex="-1" />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#configuration)

This rule accepts a configuration object with the following properties:

### allowExpressionValues [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#allowexpressionvalues)

type: `boolean`

default: `true`

If `true`, allows tabIndex values to be expression values (e.g., variables, ternaries). If `false`, only string literal values are allowed.

### roles [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#roles)

type: `string[]`

default: `["tabpanel"]`

An array of ARIA roles that should be considered interactive.

### tags [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#tags)

type: `string[]`

default: `[]`

An array of custom HTML elements that should be considered interactive.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/no-noninteractive-tabindex": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/no-noninteractive-tabindex --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/no_noninteractive_tabindex.rs)
