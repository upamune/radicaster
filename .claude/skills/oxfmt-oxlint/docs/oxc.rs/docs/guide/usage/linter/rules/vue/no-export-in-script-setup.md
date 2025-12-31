---
title: "vue/no-export-in-script-setup | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/no-export-in-script-setup.md for this page in Markdown format

# vue/no-export-in-script-setup Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#vue-no-export-in-script-setup)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#what-it-does)

Disallow `export` in `<script setup>`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#why-is-this-bad)

The previous version of `<script setup>` RFC used `export` to define variables used in templates, but the new `<script setup>` RFC has been updated to define without using `export`. See [Vue RFCs - 0040-script-setup](https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md) for more details.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup>
export let msg = "Hello!";
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup>
let msg = "Hello!";
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/no-export-in-script-setup": "error"
  }
}
```

bash

```
oxlint --deny vue/no-export-in-script-setup --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/no_export_in_script_setup.rs)
