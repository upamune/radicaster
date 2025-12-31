---
title: "react/self-closing-comp | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/self-closing-comp.md for this page in Markdown format

# react/self-closing-comp Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#react-self-closing-comp)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#what-it-does)

Detects components without children which can be self-closed to avoid unnecessary extra closing tags.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#why-is-this-bad)

Components without children don't need explicit closing tags. Using self-closing syntax makes code more concise and reduces visual clutter. It also follows common React and JSX conventions for empty elements.

A self-closing component which contains whitespace is allowed except when it also contains a newline.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const elem = <Component linter="oxlint"></Component>;
const dom_elem = <div id="oxlint"></div>;
const welem = <div id="oxlint"></div>;
```

Examples of **correct** code for this rule:

jsx

```
const elem = <Component linter="oxlint" />;
const welem = <Component linter="oxlint"> </Component>;
const dom_elem = <div id="oxlint" />;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#configuration)

This rule accepts a configuration object with the following properties:

### component [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#component)

type: `boolean`

default: `true`

Whether to enforce self-closing for custom components.

### html [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#html)

type: `boolean`

default: `true`

Whether to enforce self-closing for native HTML elements.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/self-closing-comp": "error"
  }
}
```

bash

```
oxlint --deny react/self-closing-comp --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/self_closing_comp.rs)
