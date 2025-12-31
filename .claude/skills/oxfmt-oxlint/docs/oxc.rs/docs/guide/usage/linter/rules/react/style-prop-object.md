---
title: "react/style-prop-object | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/style-prop-object.md for this page in Markdown format

# react/style-prop-object Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#react-style-prop-object)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#what-it-does)

Require that the value of the prop `style` be an object or a variable that is an object.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#why-is-this-bad)

The `style` prop expects an object mapping from style properties to values when using JSX.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div style="color: 'red'" />
<div style={true} />
<Hello style={true} />
const styles = true;
<div style={styles} />

React.createElement("div", { style: "color: 'red'" });
React.createElement("div", { style: true });
React.createElement("Hello", { style: true });
const styles = true;
React.createElement("div", { style: styles });
```

Examples of **correct** code for this rule:

jsx

```
<div style={{ color: "red" }} />
<Hello style={{ color: "red" }} />
const styles = { color: "red" };
<div style={styles} />

React.createElement("div", { style: { color: 'red' }});
React.createElement("Hello", { style: { color: 'red' }});
const styles = { height: '100px' };
React.createElement("div", { style: styles });
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#allow)

type: `string[]`

default: `[]`

List of component names on which to allow `style` prop values of any type.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/style-prop-object": "error"
  }
}
```

bash

```
oxlint --deny react/style-prop-object --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/style_prop_object.rs)
