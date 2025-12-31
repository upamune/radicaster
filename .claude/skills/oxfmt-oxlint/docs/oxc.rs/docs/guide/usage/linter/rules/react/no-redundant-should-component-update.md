---
title: "react/no-redundant-should-component-update | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-redundant-should-component-update.md for this page in Markdown format

# react/no-redundant-should-component-update Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#react-no-redundant-should-component-update)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#what-it-does)

Disallow usage of `shouldComponentUpdate` when extending `React.PureComponent`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#why-is-this-bad)

`React.PureComponent` automatically implements `shouldComponentUpdate` with a shallow prop and state comparison. Defining `shouldComponentUpdate` in a class that extends `React.PureComponent` is redundant and defeats the purpose of using `React.PureComponent`. If you need custom comparison logic, extend `React.Component` instead.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
class Foo extends React.PureComponent {
  shouldComponentUpdate() {
    // do check
  }

  render() {
    return <div>Radical!</div>;
  }
}

function Bar() {
  return class Baz extends React.PureComponent {
    shouldComponentUpdate() {
      // do check
    }

    render() {
      return <div>Groovy!</div>;
    }
  };
}
```

Examples of **correct** code for this rule:

jsx

```
class Foo extends React.Component {
  shouldComponentUpdate() {
    // do check
  }

  render() {
    return <div>Radical!</div>;
  }
}

function Bar() {
  return class Baz extends React.Component {
    shouldComponentUpdate() {
      // do check
    }

    render() {
      return <div>Groovy!</div>;
    }
  };
}

class Qux extends React.PureComponent {
  render() {
    return <div>Tubular!</div>;
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-redundant-should-component-update": "error"
  }
}
```

bash

```
oxlint --deny react/no-redundant-should-component-update --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_redundant_should_component_update.rs)
