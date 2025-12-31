---
title: "react/no-set-state | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-set-state.md for this page in Markdown format

# react/no-set-state Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#react-no-set-state)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#what-it-does)

Disallow the usage of `this.setState` in React components.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#why-is-this-bad)

When using an architecture that separates your application state from your UI components (e.g. Flux), it may be desirable to forbid the use of local component state. This rule is especially helpful in read-only applications (that don't use forms), since local component state should rarely be necessary in such cases.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var Hello = createReactClass({
  getInitialState: function () {
    return {
      name: this.props.name,
    };
  },
  handleClick: function () {
    this.setState({
      name: this.props.name.toUpperCase(),
    });
  },
  render: function () {
    return <div onClick={this.handleClick.bind(this)}>Hello {this.state.name}</div>;
  },
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-set-state": "error"
  }
}
```

bash

```
oxlint --deny react/no-set-state --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_set_state.rs)
