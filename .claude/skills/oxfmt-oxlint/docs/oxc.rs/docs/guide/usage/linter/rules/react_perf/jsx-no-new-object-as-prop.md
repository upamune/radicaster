---
title: "react_perf/jsx-no-new-object-as-prop | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react\_perf/jsx-no-new-object-as-prop.md for this page in Markdown format

# react\_perf/jsx-no-new-object-as-prop Perf [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#react-perf-jsx-no-new-object-as-prop)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#what-it-does)

Prevent Objects that are local to the current method from being used as values of JSX props.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#why-is-this-bad)

Using locally defined Objects as values for props can lead to unintentional re-renders and performance issues. Every time the parent component renders, a new instance of the Object is created, causing unnecessary re-renders of child components. This also leads to harder-to-maintain code as the component's props are not passed consistently.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<Item config={{}} />
<Item config={new Object()} />
<Item config={Object()} />
<Item config={this.props.config || {}} />
<Item config={this.props.config ? this.props.config : {}} />
<div style={{display: 'none'}} />
```

Examples of **correct** code for this rule:

jsx

```
<Item config={staticConfig} />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react-perf"],
  "rules": {
    "react-perf/jsx-no-new-object-as-prop": "error"
  }
}
```

bash

```
oxlint --deny react-perf/jsx-no-new-object-as-prop --react-perf-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react_perf/jsx_no_new_object_as_prop.rs)
