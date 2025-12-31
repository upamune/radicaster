---
title: "oxc/no-rest-spread-properties | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.md for this page in Markdown format

# oxc/no-rest-spread-properties Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#oxc-no-rest-spread-properties)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#what-it-does)

Disallow [Object Rest/Spread Properties](https://github.com/tc39/proposal-object-rest-spread#readme).

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#why-is-this-bad)

Object rest/spread properties are a relatively new JavaScript feature that may not be supported in all target environments. If you need to support older browsers or JavaScript engines that don't support these features, using them can cause runtime errors. This rule helps maintain compatibility with older environments by preventing the use of these modern syntax features.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
let { x, ...y } = z;
let z = { x, ...y };
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#configuration)

This rule accepts a configuration object with the following properties:

### objectRestMessage [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#objectrestmessage)

type: `string`

default: `""`

A message to display when object rest properties are found.

### objectSpreadMessage [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#objectspreadmessage)

type: `string`

default: `""`

A message to display when object spread properties are found.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "oxc/no-rest-spread-properties": "error"
  }
}
```

bash

```
oxlint --deny oxc/no-rest-spread-properties
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/oxc/no_rest_spread_properties.rs)
