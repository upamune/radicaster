---
title: "vue/define-props-destructuring | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/define-props-destructuring.md for this page in Markdown format

# vue/define-props-destructuring Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#vue-define-props-destructuring)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#what-it-does)

This rule enforces a consistent style for handling Vue 3 Composition API props, allowing you to choose between requiring destructuring or prohibiting it.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#why-is-this-bad)

By default, the rule requires you to use destructuring syntax when using `defineProps` instead of storing props in a variable and warns against combining `withDefaults` with destructuring.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup lang="ts">
const props = defineProps(["foo"]);
const propsWithDefaults = withDefaults(defineProps(["foo"]), { foo: "default" });
const { baz } = withDefaults(defineProps(["baz"]), { baz: "default" });
const props = defineProps<{ foo?: string }>();
const propsWithDefaults = withDefaults(defineProps<{ foo?: string }>(), { foo: "default" });
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup lang="ts">
const { foo } = defineProps(["foo"]);
const { bar = "default" } = defineProps(["bar"]);
const { foo } = defineProps<{ foo?: string }>();
const { bar = "default" } = defineProps<{ bar?: string }>();
</script>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#configuration)

This rule accepts a configuration object with the following properties:

### destructure [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#destructure)

type: `"always" | "never"`

default: `"always"`

Require or prohibit destructuring.

#### `"always"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#always)

Requires destructuring when using `defineProps` and warns against using `withDefaults` with destructuring

#### `"never"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#never)

Requires using a variable to store props and prohibits destructuring

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/define-props-destructuring": "error"
  }
}
```

bash

```
oxlint --deny vue/define-props-destructuring --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/define_props_destructuring.rs)
