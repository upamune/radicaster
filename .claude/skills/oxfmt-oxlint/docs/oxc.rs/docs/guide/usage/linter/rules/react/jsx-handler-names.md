---
title: "react/jsx-handler-names | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-handler-names.md for this page in Markdown format

# react/jsx-handler-names Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#react-jsx-handler-names)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#what-it-does)

Ensures that any component or prop methods used to handle events are correctly prefixed.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#why-is-this-bad)

Inconsistent naming of event handlers and props can reduce code readability and maintainability.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<MyComponent handleChange={this.handleChange} />
<MyComponent onChange={this.componentChanged} />
```

Examples of **correct** code for this rule:

jsx

```
<MyComponent onChange={this.handleChange} />
<MyComponent onChange={this.props.onFoo} />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#configuration)

This rule accepts a configuration object with the following properties:

### checkInlineFunctions [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#checkinlinefunctions)

type: `boolean`

default: `false`

Whether to check for inline functions in JSX attributes.

### checkLocalVariables [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#checklocalvariables)

type: `boolean`

default: `false`

Whether to check for local variables in JSX attributes.

### eventHandlerPrefixes [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#eventhandlerprefixes)

type: `string`

default: `"handle"`

Event handler prefixes to check against.

### eventHandlerPropPrefixes [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#eventhandlerpropprefixes)

type: `string`

default: `"on"`

Event handler prop prefixes to check against.

### eventHandlerPropRegex [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#eventhandlerpropregex)

type: `string | null`

Compiled regex for event handler prop prefixes.

### eventHandlerRegex [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#eventhandlerregex)

type: `string | null`

Compiled regex for event handler prefixes.

### ignoreComponentNames [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#ignorecomponentnames)

type: `string[]`

default: `[]`

Component names to ignore when checking for event handler prefixes.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-handler-names": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-handler-names --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_handler_names.rs)
