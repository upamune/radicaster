---
title: "jsx_a11y/anchor-ambiguous-text | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/anchor-ambiguous-text.md for this page in Markdown format

# jsx\_a11y/anchor-ambiguous-text Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#jsx-a11y-anchor-ambiguous-text)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#what-it-does)

Inspects anchor link text for the use of ambiguous words.

This rule checks the text from the anchor element `aria-label` if available. In absence of an anchor `aria-label` it combines the following text of it's children:

* `aria-label` if available
* if the child is an image, the `alt` text
* the text content of the HTML element

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#why-is-this-bad)

Screen readers users rely on link text for context, ambiguous words such as "click here" do not provide enough context.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<a>link</a>
<a>click here</a>
```

Examples of **correct** code for this rule:

jsx

```
<a>read this tutorial</a>
<a aria-label="oxc linter documentation">click here</a>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#configuration)

This rule accepts a configuration object with the following properties:

### words [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#words)

type: `string[]`

default: `["click here", "here", "link", "a link", "learn more"]`

List of ambiguous words or phrases that should be flagged in anchor text.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/anchor-ambiguous-text": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/anchor-ambiguous-text --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/anchor_ambiguous_text.rs)
