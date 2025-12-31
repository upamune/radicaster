---
title: "react/no-children-prop | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-children-prop.md for this page in Markdown format

# react/no-children-prop Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#react-no-children-prop)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#what-it-does)

Checks that children are not passed using a prop.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#why-is-this-bad)

Children should always be actual children, not passed in as a prop. When using JSX, the children should be nested between the opening and closing tags. When not using JSX, the children should be passed as additional arguments to `React.createElement`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div children='Children' />

<MyComponent children={<AnotherComponent />} />
<MyComponent children={['Child 1', 'Child 2']} />
React.createElement("div", { children: 'Children' })
```

Examples of **correct** code for this rule:

jsx

```
<div>Children</div>
<MyComponent>Children</MyComponent>

<MyComponent>
  <span>Child 1</span>
  <span>Child 2</span>
</MyComponent>

React.createElement("div", {}, 'Children')
React.createElement("div", 'Child 1', 'Child 2')
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-children-prop": "error"
  }
}
```

bash

```
oxlint --deny react/no-children-prop --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_children_prop.rs)
