---
title: "vue/no-import-compiler-macros | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/no-import-compiler-macros.md for this page in Markdown format

# vue/no-import-compiler-macros Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#vue-no-import-compiler-macros)

⚠️🛠️️ A dangerous auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#what-it-does)

Disallow importing Vue compiler macros.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#why-is-this-bad)

Compiler Macros like:

* `defineProps`
* `defineEmits`
* `defineExpose`
* `withDefaults`
* `defineModel`
* `defineOptions`
* `defineSlots`

are globally available in Vue 3's `<script setup>` and do not require explicit imports.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup>
import { defineProps, withDefaults } from "vue";
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup>
import { ref } from "vue";
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/no-import-compiler-macros": "error"
  }
}
```

bash

```
oxlint --deny vue/no-import-compiler-macros --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/no_import_compiler_macros.rs)
