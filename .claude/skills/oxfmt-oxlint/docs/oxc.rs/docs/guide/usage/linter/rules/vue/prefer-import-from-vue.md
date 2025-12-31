---
title: "vue/prefer-import-from-vue | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/prefer-import-from-vue.md for this page in Markdown format

# vue/prefer-import-from-vue Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#vue-prefer-import-from-vue)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#what-it-does)

Enforce `import from 'vue'` instead of `import from '@vue/*'`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#why-is-this-bad)

Imports from the following modules are almost always wrong. You should import from vue instead.

* `@vue/runtime-dom`
* `@vue/runtime-core`
* `@vue/reactivity`
* `@vue/shared`

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#examples)

Examples of **incorrect** code for this rule:

js

```
import { createApp } from "@vue/runtime-dom";
import { Component } from "@vue/runtime-core";
import { ref } from "@vue/reactivity";
```

Examples of **correct** code for this rule:

js

```
import { createApp, ref, Component } from "vue";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/prefer-import-from-vue": "error"
  }
}
```

bash

```
oxlint --deny vue/prefer-import-from-vue --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/prefer_import_from_vue.rs)
