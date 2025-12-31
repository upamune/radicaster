---
title: "react/void-dom-elements-no-children | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/void-dom-elements-no-children.md for this page in Markdown format

# react/void-dom-elements-no-children Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#react-void-dom-elements-no-children)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#what-it-does)

Disallow void DOM elements (e.g. `<img />`, `<br />`) from receiving children.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#why-is-this-bad)

There are some HTML elements that are only self-closing (e.g. img, br, hr). These are collectively known as void DOM elements. This rule checks that children are not passed to void DOM elements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<br>Children</br>
<br children='Children' />
<br dangerouslySetInnerHTML={{ __html: 'HTML' }} />
React.createElement('br', undefined, 'Children')
React.createElement('br', { children: 'Children' })
React.createElement('br', { dangerouslySetInnerHTML: { __html: 'HTML' } })
```

Examples of **correct** code for this rule:

jsx

```
<div>Children</div>
<div children='Children' />
<div dangerouslySetInnerHTML={{ __html: 'HTML' }} />
React.createElement('div', undefined, 'Children')
React.createElement('div', { children: 'Children' })
React.createElement('div', { dangerouslySetInnerHTML: { __html: 'HTML' } })
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/void-dom-elements-no-children": "error"
  }
}
```

bash

```
oxlint --deny react/void-dom-elements-no-children --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/void_dom_elements_no_children.rs)
