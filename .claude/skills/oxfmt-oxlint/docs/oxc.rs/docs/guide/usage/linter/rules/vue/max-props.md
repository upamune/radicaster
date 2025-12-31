---
title: "vue/max-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/max-props.md for this page in Markdown format

# vue/max-props Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#vue-max-props)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#what-it-does)

Enforce maximum number of props in Vue component.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#why-is-this-bad)

This rule enforces a maximum number of props in a Vue SFC, in order to aid in maintainability and reduce complexity.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#examples)

Examples of **incorrect** code for this rule with the default `{ maxProps: 1 }` option:

js

```
<script setup>
defineProps({
  prop1: String,
  prop2: String,
})
</script>
```

Examples of **correct** code for this rule with the default `{ maxProps: 1 }` option:

js

```
<script setup>
defineProps({
  prop1: String,
})
</script>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#configuration)

This rule accepts a configuration object with the following properties:

### maxProps [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#maxprops)

type: `integer`

default: `1`

The maximum number of props allowed in a Vue Single File Component (SFC).

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/max-props": "error"
  }
}
```

bash

```
oxlint --deny vue/max-props --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/max_props.rs)
