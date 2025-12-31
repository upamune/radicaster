---
title: "react/prefer-es6-class | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/prefer-es6-class.md for this page in Markdown format

# react/prefer-es6-class Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#react-prefer-es6-class)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#what-it-does)

React offers you two ways to create traditional components: using the ES5 create-react-class module or the new ES2015 class system.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#why-is-this-bad)

This rule enforces a consistent React class style.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var Hello = createReactClass({
  render: function () {
    return <div>Hello {this.props.name}</div>;
  },
});
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#always)

Always prefer ES6 class-style components

### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#never)

Do not allow ES6 class-style

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/prefer-es6-class": "error"
  }
}
```

bash

```
oxlint --deny react/prefer-es6-class --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/prefer_es6_class.rs)
