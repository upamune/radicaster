---
title: "react/forbid-dom-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/forbid-dom-props.md for this page in Markdown format

# react/forbid-dom-props Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#react-forbid-dom-props)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#what-it-does)

This rule prevents passing of props to elements. This rule only applies to DOM Nodes (e.g.) and not Components (e.g. ). The list of forbidden props can be customized with the forbid option.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#why-is-this-bad)

This rule checks all JSX elements and verifies that no forbidden props are used on DOM Nodes. This rule is off by default.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// [1, { "forbid": ["id"] }]
<div id='Joe' />

// [1, { "forbid": ["style"] }]
<div style={{color: 'red'}} />
```

Examples of **correct** code for this rule:

jsx

```
// [1, { "forbid": ["id"] }]
<Hello id='foo' />

// [1, { "forbid": ["id"] }]
<Hello id={{color: 'red'}} />
```

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#options)

#### forbid [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#forbid)

An array of strings, with the names of props that are forbidden. The default value of this option []. Each array element can either be a string with the property name or object specifying the property name, an optional custom message, and a DOM nodes disallowed list (e.g.)

`{"propName": "someProp", "disallowedFor": ["DOMNode", "AnotherDOMNode"], "message": "Avoid using someProp" }`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/forbid-dom-props": "error"
  }
}
```

bash

```
oxlint --deny react/forbid-dom-props --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/forbid_dom_props.rs)
