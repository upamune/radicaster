---
title: "react/no-render-return-value | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-render-return-value.md for this page in Markdown format

# react/no-render-return-value Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#react-no-render-return-value)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#what-it-does)

This rule will warn you if you try to use the `ReactDOM.render()` return value.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#why-is-this-bad)

Using the return value from `ReactDOM.render()` is a legacy feature and should not be used.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
vaa inst =ReactDOM.render(<App />, document.body);
function render() {
 return ReactDOM.render(<App />, document.body);
}
```

Examples of **correct** code for this rule:

jsx

```
ReactDOM.render(<App />, document.body);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-render-return-value": "error"
  }
}
```

bash

```
oxlint --deny react/no-render-return-value --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_render_return_value.rs)
