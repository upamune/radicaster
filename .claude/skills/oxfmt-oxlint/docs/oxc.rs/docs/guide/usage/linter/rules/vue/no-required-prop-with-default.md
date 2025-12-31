---
title: "vue/no-required-prop-with-default | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/no-required-prop-with-default.md for this page in Markdown format

# vue/no-required-prop-with-default Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#vue-no-required-prop-with-default)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#what-it-does)

Enforce props with default values to be optional.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#why-is-this-bad)

If a prop is declared with a default value, whether it is required or not, we can always skip it in actual use. In that situation, the default value would be applied. So, a required prop with a default value is essentially the same as an optional prop.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    name: string | number;
    age?: number;
  }>(),
  {
    name: "Foo",
  },
);
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup lang="ts">
const props = withDefaults(
  defineProps<{
    name?: string | number;
    age?: number;
  }>(),
  {
    name: "Foo",
  },
);
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/no-required-prop-with-default": "error"
  }
}
```

bash

```
oxlint --deny vue/no-required-prop-with-default --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/no_required_prop_with_default.rs)
