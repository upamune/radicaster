---
title: "vue/valid-define-emits | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/valid-define-emits.md for this page in Markdown format

# vue/valid-define-emits Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#vue-valid-define-emits)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#what-it-does)

This rule checks whether defineEmits compiler macro is valid.

This rule reports defineEmits compiler macros in the following cases:

* `defineEmits` is referencing locally declared variables.
* `defineEmits` has both a literal type and an argument. e.g. `defineEmits<(e: 'foo')=>void>(['bar'])`
* `defineEmits` has been called multiple times.
* Custom events are defined in both `defineEmits` and `export default {}`.
* Custom events are not defined in either `defineEmits` or `export default {}`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#why-is-this-bad)

Misusing `defineEmits` can lead to runtime errors, unclear component contracts, and lost type safety. Vue may still compile the code, but emitted events may break silently or be typed incorrectly.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script setup>
const def = { notify: null };
defineEmits(def);
</script>
```

vue

```
<script setup lang="ts">
defineEmits<(e: "notify") => void>({ submit: null });
</script>
```

vue

```
<script setup>
defineEmits({ notify: null });
defineEmits({ submit: null });
</script>
```

vue

```
<script>
export default {
  emits: ["notify"],
};
</script>
<script setup>
defineEmits({ submit: null });
</script>
```

Examples of **correct** code for this rule:

vue

```
<script setup>
defineEmits({ notify: null });
</script>
```

vue

```
<script setup>
defineEmits(["notify"]);
</script>
```

vue

```
<script setup lang="ts">
defineEmits<(e: "notify") => void>();
</script>
```

vue

```
<script>
export default {
  emits: ["notify"],
};
</script>
<script setup>
defineEmits();
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/valid-define-emits": "error"
  }
}
```

bash

```
oxlint --deny vue/valid-define-emits --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/valid_define_emits.rs)
