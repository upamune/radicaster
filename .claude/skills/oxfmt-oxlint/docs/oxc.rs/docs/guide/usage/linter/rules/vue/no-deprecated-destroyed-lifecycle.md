---
title: "vue/no-deprecated-destroyed-lifecycle | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.md for this page in Markdown format

# vue/no-deprecated-destroyed-lifecycle Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#vue-no-deprecated-destroyed-lifecycle)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#what-it-does)

Disallow using deprecated `destroyed` and `beforeDestroy` lifecycle hooks in Vue.js 3.0.0+.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#why-is-this-bad)

In Vue.js 3.0.0+, the `destroyed` and `beforeDestroy` lifecycle hooks have been renamed to `unmounted` and `beforeUnmount` respectively. Using the old names is deprecated and may cause confusion or compatibility issues.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script>
export default {
  beforeDestroy() {},
  destroyed() {},
};
</script>
```

Examples of **correct** code for this rule:

vue

```
<script>
export default {
  beforeUnmount() {},
  unmounted() {},
};
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/no-deprecated-destroyed-lifecycle": "error"
  }
}
```

bash

```
oxlint --deny vue/no-deprecated-destroyed-lifecycle --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/no_deprecated_destroyed_lifecycle.rs)
