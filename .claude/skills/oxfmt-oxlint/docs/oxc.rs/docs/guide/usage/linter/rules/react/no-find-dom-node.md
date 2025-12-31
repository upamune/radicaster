---
title: "react/no-find-dom-node | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-find-dom-node.md for this page in Markdown format

# react/no-find-dom-node Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#react-no-find-dom-node)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#what-it-does)

This rule disallows the use of `findDOMNode`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#why-is-this-bad)

`findDOMNode` is an escape hatch used to access the underlying DOM node. In most cases, use of this escape hatch is discouraged because it pierces the component abstraction. [It has been deprecated in `StrictMode`.](https://legacy.reactjs.org/docs/strict-mode.html#warning-about-deprecated-finddomnode-usage)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
class MyComponent extends Component {
  componentDidMount() {
    findDOMNode(this).scrollIntoView();
  }
  render() {
    return <div />;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-find-dom-node": "error"
  }
}
```

bash

```
oxlint --deny react/no-find-dom-node --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_find_dom_node.rs)
