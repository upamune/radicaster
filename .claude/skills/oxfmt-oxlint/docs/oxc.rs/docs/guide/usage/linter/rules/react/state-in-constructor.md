---
title: "react/state-in-constructor | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/state-in-constructor.md for this page in Markdown format

# react/state-in-constructor Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#react-state-in-constructor)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#what-it-does)

Enforces the state initialization style to be either in a constructor or with a class property.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#why-is-this-bad)

Inconsistent state initialization styles can make the codebase harder to maintain and understand. This rule enforces a consistent pattern across React class components.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#examples)

This rule has two modes: `"always"` and `"never"`.

#### `"always"` mode [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#always-mode)

Will enforce the state initialization style to be in a constructor. This is the default mode.

Examples of **incorrect** code for this rule:

jsx

```
class Foo extends React.Component {
  state = { bar: 0 };
  render() {
    return <div>Foo</div>;
  }
}
```

Examples of **correct** code for this rule:

jsx

```
class Foo extends React.Component {
  constructor(props) {
    super(props);
    this.state = { bar: 0 };
  }
  render() {
    return <div>Foo</div>;
  }
}
```

#### `"never"` mode [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#never-mode)

Will enforce the state initialization style to be with a class property.

Examples of **incorrect** code for this rule:

jsx

```
class Foo extends React.Component {
  constructor(props) {
    super(props);
    this.state = { bar: 0 };
  }
  render() {
    return <div>Foo</div>;
  }
}
```

Examples of **correct** code for this rule:

jsx

```
class Foo extends React.Component {
  state = { bar: 0 };
  render() {
    return <div>Foo</div>;
  }
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#configuration)

This rule accepts one of the following string values:

### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#always)

Enforce state initialization in the constructor.

### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#never)

Enforce state initialization with a class property.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/state-in-constructor": "error"
  }
}
```

bash

```
oxlint --deny react/state-in-constructor --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/state_in_constructor.rs)
