---
title: "react/jsx-key | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-key.md for this page in Markdown format

# react/jsx-key Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#react-jsx-key)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#what-it-does)

Enforce `key` prop for elements in array

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#why-is-this-bad)

React requires a `key` prop for elements in an array to help identify which items have changed, are added, or are removed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
[1, 2, 3].map((x) => <App />);
[1, 2, 3]?.map((x) => <BabelEslintApp />);
```

Examples of **correct** code for this rule:

jsx

```
[1, 2, 3].map((x) => <App key={x} />);
[1, 2, 3]?.map((x) => <BabelEslintApp key={x} />);
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#configuration)

This rule accepts a configuration object with the following properties:

### checkFragmentShorthand [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#checkfragmentshorthand)

type: `boolean`

default: `true`

When true, check fragment shorthand `<>` for keys

### checkKeyMustBeforeSpread [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#checkkeymustbeforespread)

type: `boolean`

default: `true`

When true, require key prop to be placed before any spread props

### warnOnDuplicates [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#warnonduplicates)

type: `boolean`

default: `true`

When true, warn on duplicate key values

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-key": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-key --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_key.rs)
