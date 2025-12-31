---
title: "react/require-render-return | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/require-render-return.md for this page in Markdown format

# react/require-render-return Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#react-require-render-return)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#what-it-does)

Enforce ES5 or ES2015 class for returning value in render function

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#why-is-this-bad)

When writing the `render` method in a component it is easy to forget to return the JSX content. This rule will warn if the return statement is missing.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var Hello = createReactClass({
  render() {
    <div>Hello</div>;
  },
});

class Hello extends React.Component {
  render() {
    <div>Hello</div>;
  }
}
```

Examples of **correct** code for this rule:

jsx

```
var Hello = createReactClass({
  render() {
    return <div>Hello</div>;
  },
});

class Hello extends React.Component {
  render() {
    return <div>Hello</div>;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/require-render-return": "error"
  }
}
```

bash

```
oxlint --deny react/require-render-return --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/require_render_return.rs)
