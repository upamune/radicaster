---
title: "vue/define-props-declaration | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/define-props-declaration.md for this page in Markdown format

# vue/define-props-declaration Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#vue-define-props-declaration)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#what-it-does)

This rule enforces `defineProps` typing style which you should use `type-based` or `runtime` declaration. This rule only works in setup script and `lang="ts"`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#examples)

Examples of **incorrect** code for this rule:

vue

```
// "vue/define-props-declaration": ["error", "type-based"]
<script setup lang="ts">
const props = defineProps({
  kind: { type: String },
});
</script>

// "vue/define-props-declaration": ["error", "runtime"]
<script setup lang="ts">
const props = defineProps<{
  kind: string;
}>();
</script>
```

Examples of **correct** code for this rule:

vue

```
// "vue/define-props-declaration": ["error", "type-based"]
<script setup lang="ts">
const props = defineProps<{
  kind: string;
}>();
</script>

// "vue/define-props-declaration": ["error", "runtime"]
<script setup lang="ts">
const props = defineProps({
  kind: { type: String },
});
</script>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#configuration)

This rule accepts one of the following string values:

### `"type-based"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#type-based)

Enforce type-based declaration.

### `"runtime"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#runtime)

Enforce runtime declaration.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/define-props-declaration": "error"
  }
}
```

bash

```
oxlint --deny vue/define-props-declaration --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/define_props_declaration.rs)
