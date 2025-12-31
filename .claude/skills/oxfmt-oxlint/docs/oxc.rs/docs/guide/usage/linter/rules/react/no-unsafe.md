---
title: "react/no-unsafe | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-unsafe.md for this page in Markdown format

# react/no-unsafe Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#react-no-unsafe)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#what-it-does)

This rule identifies and restricts the use of unsafe React lifecycle methods.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#why-is-this-bad)

Certain lifecycle methods (`componentWillMount`, `componentWillReceiveProps`, and `componentWillUpdate`) are considered unsafe and have been deprecated since React 16.9. They are frequently misused and cause problems in async rendering. Using their `UNSAFE_` prefixed versions or the deprecated names themselves should be avoided.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// By default, UNSAFE_ prefixed methods are flagged
class Foo extends React.Component {
  UNSAFE_componentWillMount() {}
  UNSAFE_componentWillReceiveProps() {}
  UNSAFE_componentWillUpdate() {}
}

// With checkAliases: true, non-prefixed versions are also flagged
class Bar extends React.Component {
  componentWillMount() {}
  componentWillReceiveProps() {}
  componentWillUpdate() {}
}
```

Examples of **correct** code for this rule:

jsx

```
class Foo extends React.Component {
  componentDidMount() {}
  componentDidUpdate() {}
  render() {}
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#configuration)

This rule accepts a configuration object with the following properties:

### checkAliases [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#checkaliases)

type: `boolean`

default: `false`

Whether to check for the non-prefixed lifecycle methods. If `true`, this means `componentWillMount`, `componentWillReceiveProps`, and `componentWillUpdate` will also be flagged, rather than just the UNSAFE\_ versions. It is recommended to set this to `true` to fully avoid unsafe lifecycle methods.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-unsafe": "error"
  }
}
```

bash

```
oxlint --deny react/no-unsafe --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_unsafe.rs)
