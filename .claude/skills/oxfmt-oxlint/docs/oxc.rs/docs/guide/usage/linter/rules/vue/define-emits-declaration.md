---
title: "vue/define-emits-declaration | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/define-emits-declaration.md for this page in Markdown format

# vue/define-emits-declaration Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#vue-define-emits-declaration)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#what-it-does)

This rule enforces `defineEmits` typing style which you should use `type-based`, strict `type-literal` (introduced in Vue 3.3), or `runtime` declaration. This rule only works in setup script and `lang="ts"`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#examples)

Examples of **incorrect** code for this rule:

vue

```
// "vue/define-emits-declaration": ["error", "type-based"]
<script setup lang="ts">
const emit = defineEmits(["change", "update"]);
const emit2 = defineEmits({
  change: (id) => typeof id === "number",
  update: (value) => typeof value === "string",
});
</script>

// "vue/define-emits-declaration": ["error", "type-literal"]
<script setup lang="ts">
const emit = defineEmits<{
  (e: "change", id: number): void;
  (e: "update", value: string): void;
}>();
</script>

// "vue/define-emits-declaration": ["error", "runtime"]
<script setup lang="ts">
const emit = defineEmits<{
  (e: "change", id: number): void;
  (e: "update", value: string): void;
}>();
</script>
```

Examples of **correct** code for this rule:

vue

```
// "vue/define-emits-declaration": ["error", "type-based"]
<script setup lang="ts">
const emit = defineEmits<{
  (e: "change", id: number): void;
  (e: "update", value: string): void;
}>();
const emit2 = defineEmits<{
  change: [id: number];
  update: [value: string];
}>();
</script>

// "vue/define-emits-declaration": ["error", "type-literal"]
<script setup lang="ts">
const emit = defineEmits<{
  change: [id: number];
  update: [value: string];
}>();
</script>

// "vue/define-emits-declaration": ["error", "runtime"]
<script setup lang="ts">
const emit = defineEmits<{
  (e: "change", id: number): void;
  (e: "update", value: string): void;
}>();
const emit2 = defineEmits({
  change: (id) => typeof id === "number",
  update: (value) => typeof value === "string",
});
</script>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#configuration)

This rule accepts one of the following string values:

### `"type-based"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#type-based)

Enforces the use of a named TypeScript type or interface as the argument to `defineEmits`, e.g. `defineEmits<MyEmits>()`.

### `"type-literal"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#type-literal)

Enforces the use of an inline type literal as the argument to `defineEmits`, e.g. `defineEmits<{ (event: string): void }>()`.

### `"runtime"` [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#runtime)

Enforces the use of runtime declaration, where emits are declared using an array or object, e.g. `defineEmits(['event1', 'event2'])`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/define-emits-declaration": "error"
  }
}
```

bash

```
oxlint --deny vue/define-emits-declaration --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/define_emits_declaration.rs)
