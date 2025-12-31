---
title: "react/jsx-pascal-case | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-pascal-case.md for this page in Markdown format

# react/jsx-pascal-case Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#react-jsx-pascal-case)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#what-it-does)

Enforce PascalCase for user-defined JSX components

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#why-is-this-bad)

It enforces coding style that user-defined JSX components are defined and referenced in PascalCase. Note that since React's JSX uses the upper vs. lower case convention to distinguish between local component classes and HTML tags this rule will not warn on components that start with a lower case letter.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<Test_component />
<TEST_COMPONENT />
```

Examples of **correct** code for this rule:

jsx

```
<div />

<TestComponent />

<TestComponent>
    <div />
</TestComponent>

<CSSTransitionGroup />
```

Examples of **correct** code for the "allowAllCaps" option:

jsx

```
<ALLOWED />

<TEST_COMPONENT />
```

Examples of **correct** code for the "allowNamespace" option:

jsx

```
<Allowed.div />

<TestComponent.p />
```

Examples of **correct** code for the "allowLeadingUnderscore" option:

jsx

```
<_AllowedComponent />

<_AllowedComponent>
    <div />
</_AllowedComponent>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#configuration)

This rule accepts a configuration object with the following properties:

### allowAllCaps [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#allowallcaps)

type: `boolean`

default: `false`

Whether to allow all-caps component names.

### allowLeadingUnderscore [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#allowleadingunderscore)

type: `boolean`

default: `false`

Whether to allow leading underscores in component names.

### allowNamespace [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#allownamespace)

type: `boolean`

default: `false`

Whether to allow namespaced component names.

### ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#ignore)

type: `string[]`

default: `[]`

List of component names to ignore.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-pascal-case": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-pascal-case --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_pascal_case.rs)
