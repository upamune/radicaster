---
title: "react/forward-ref-uses-ref | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/forward-ref-uses-ref.md for this page in Markdown format

# react/forward-ref-uses-ref Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#react-forward-ref-uses-ref)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#what-it-does)

Requires that components wrapped with `forwardRef` must have a `ref` parameter. Omitting the `ref` argument is usually a bug, and components not using `ref` don't need to be wrapped by `forwardRef`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#why-is-this-bad)

Omitting the `ref` argument makes the `forwardRef` wrapper unnecessary, and can lead to confusion.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var React = require("react");

var Component = React.forwardRef((props) => <div />);
```

Examples of **correct** code for this rule:

jsx

```
var React = require("react");

var Component = React.forwardRef((props, ref) => <div ref={ref} />);

var Component = React.forwardRef((props, ref) => <div />);

function Component(props) {
  return <div />;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/forward-ref-uses-ref": "error"
  }
}
```

bash

```
oxlint --deny react/forward-ref-uses-ref --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/forward_ref_uses_ref.rs)
