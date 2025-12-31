---
title: "react/jsx-boolean-value | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-boolean-value.md for this page in Markdown format

# react/jsx-boolean-value Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#react-jsx-boolean-value)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#what-it-does)

Enforce a consistent boolean attribute style in your code.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#why-is-this-bad)

In JSX, you can set a boolean attribute to `true` or omit it. This rule will enforce a consistent style for boolean attributes.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const Hello = <Hello personal={true} />;
```

Examples of **correct** code for this rule:

jsx

```
const Hello = <Hello personal />;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#configuration)

This rule accepts a configuration object with the following properties:

### assumeUndefinedIsFalse [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#assumeundefinedisfalse)

type: `boolean`

default: `false`

If true, treats `prop={false}` as equivalent to the prop being undefined

### enforceBooleanAttribute [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#enforcebooleanattribute)

type: `"always" | "never"`

default: `"never"`

Enforce boolean attributes to always or never have a value.

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#exceptions)

type: `string[]`

default: `[]`

List of attribute names to exclude from the rule.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-boolean-value": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-boolean-value --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_boolean_value.rs)
