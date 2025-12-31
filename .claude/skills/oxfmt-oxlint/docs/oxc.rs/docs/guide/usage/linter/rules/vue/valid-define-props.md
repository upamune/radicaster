---
title: "vue/valid-define-props | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/valid-define-props.md for this page in Markdown format

# vue/valid-define-props Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#vue-valid-define-props)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#what-it-does)

This rule checks whether `defineProps` compiler macro is valid.

This rule reports `defineProps` compiler macros in the following cases:

* `defineProps` is referencing locally declared variables.
* `defineProps` has both a literal type and an argument. e.g. `defineProps<{ /*props*/ }>({ /*props*/ })`
* `defineProps` has been called multiple times.
* Props are defined in both `defineProps` and `export default {}`.
* Props are not defined in either `defineProps` or `export default {}`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#why-is-this-bad)

Misusing `defineProps` can lead to runtime errors, and lost type safety. Vue may still compile the code, but properties may break silently or be typed incorrectly.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup>
const def = { msg: String };
defineProps(def);
</script>
```

vue

```
<script setup lang="ts">
defineProps<{ msg?: string }>({ msg: String });
</script>
```

vue

```
<script setup>
defineProps({ msg: String });
defineProps({ count: Number });
</script>
```

vue

```
<script>
export default {
  props: { msg: String },
};
</script>
<script setup>
defineProps({ count: Number });
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup>
defineProps({ msg: String });
</script>
```

vue

```
<script setup>
defineProps(["msg"]);
</script>
```

vue

```
<script setup lang="ts">
defineProps<{ msg?: string }>();
</script>
```

vue

```
<script>
export default {
  props: { msg: String },
};
</script>
<script setup>
defineProps();
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/valid-define-props": "error"
  }
}
```

bash

```
oxlint --deny vue/valid-define-props --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/valid_define_props.rs)
