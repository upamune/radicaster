---
title: "vue/require-default-export | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/require-default-export.md for this page in Markdown format

# vue/require-default-export Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#vue-require-default-export)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#what-it-does)

Require components to be the default export.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#why-is-this-bad)

Using SFCs (Single File Components) without a default export is not supported in Vue 3. Components should be exported as the default export.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script>
const foo = "foo";
</script>
```

Examples of **correct** code for this rule:

vue

```
<script>
export default {
  data() {
    return {
      foo: "foo",
    };
  },
};
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/require-default-export": "error"
  }
}
```

bash

```
oxlint --deny vue/require-default-export --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/require_default_export.rs)
