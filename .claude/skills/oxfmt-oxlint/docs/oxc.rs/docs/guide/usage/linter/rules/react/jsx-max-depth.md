---
title: "react/jsx-max-depth | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-max-depth.md for this page in Markdown format

# react/jsx-max-depth Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#react-jsx-max-depth)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#what-it-does)

Enforces a maximum depth for nested JSX elements and fragments.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#why-is-this-bad)

Excessively nested JSX makes components harder to read and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const Component = () => (
  <div>
    <div>
      <div>
        <span />
      </div>
    </div>
  </div>
);
```

Examples of **correct** code for this rule:

jsx

```
const Component = () => (
  <div>
    <div>
      <span />
    </div>
  </div>
);
```

### Options [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#options)

`react/jsx-max-depth: [<enabled>, { "max": <number> }]`

The `max` option defaults to `2`.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#configuration)

This rule accepts a configuration object with the following properties:

### max [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#max)

type: `integer`

default: `2`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-max-depth": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-max-depth --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_max_depth.rs)
