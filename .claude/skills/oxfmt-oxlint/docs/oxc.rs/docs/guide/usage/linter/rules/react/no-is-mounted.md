---
title: "react/no-is-mounted | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-is-mounted.md for this page in Markdown format

# react/no-is-mounted Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#react-no-is-mounted)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#what-it-does)

This rule prevents using `isMounted` in classes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#why-is-this-bad)

`isMounted` is an anti-pattern, is not available when using classes, and it is on its way to being officially deprecated.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
class Hello extends React.Component {
  someMethod() {
    if (!this.isMounted()) {
      return;
    }
  }
  render() {
    return <div onClick={this.someMethod.bind(this)}>Hello</div>;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-is-mounted": "error"
  }
}
```

bash

```
oxlint --deny react/no-is-mounted --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_is_mounted.rs)
