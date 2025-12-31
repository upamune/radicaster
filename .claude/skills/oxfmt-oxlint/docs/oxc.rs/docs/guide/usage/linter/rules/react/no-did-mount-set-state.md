---
title: "react/no-did-mount-set-state | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-did-mount-set-state.md for this page in Markdown format

# react/no-did-mount-set-state Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#react-no-did-mount-set-state)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#what-it-does)

Disallows using `setState` in the `componentDidMount` lifecycle method.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#why-is-this-bad)

Updating the state after a component mount will trigger a second `render()` call and can lead to property/layout thrashing. This can cause performance issues and unexpected behavior.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var Hello = createReactClass({
  componentDidMount: function () {
    this.setState({
      name: this.props.name.toUpperCase(),
    });
  },
  render: function () {
    return <div>Hello {this.state.name}</div>;
  },
});
```

Examples of **correct** code for this rule:

jsx

```
var Hello = createReactClass({
  componentDidMount: function () {
    this.onMount(function callback(newName) {
      this.setState({
        name: newName,
      });
    });
  },
  render: function () {
    return <div>Hello {this.state.name}</div>;
  },
});
```

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#options)

The rule accepts a string value `"disallow-in-func"`:

json

```
{
  "react/no-did-mount-set-state": ["error", "disallow-in-func"]
}
```

When set, also disallows `setState` calls in nested functions within `componentDidMount`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-did-mount-set-state": "error"
  }
}
```

bash

```
oxlint --deny react/no-did-mount-set-state --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_did_mount_set_state.rs)
