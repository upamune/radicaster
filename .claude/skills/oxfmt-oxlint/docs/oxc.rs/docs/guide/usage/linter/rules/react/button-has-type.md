---
title: "react/button-has-type | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/button-has-type.md for this page in Markdown format

# react/button-has-type Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#react-button-has-type)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#what-it-does)

Enforces explicit `type` attribute for all the `button` HTML elements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#why-is-this-bad)

The default value of `type` attribute for `button` HTML element is `"submit"` which is often not the desired behavior and may lead to unexpected page reloads.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<button />
<button type="foo" />
```

Examples of **correct** code for this rule:

jsx

```
<button type="button" />
<button type="submit" />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#configuration)

This rule accepts a configuration object with the following properties:

### button [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#button)

type: `boolean`

default: `true`

If true, allow `type="button"`.

### reset [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#reset)

type: `boolean`

default: `true`

If true, allow `type="reset"`.

### submit [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#submit)

type: `boolean`

default: `true`

If true, allow `type="submit"`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/button-has-type": "error"
  }
}
```

bash

```
oxlint --deny react/button-has-type --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/button_has_type.rs)
