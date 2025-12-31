---
title: "vue/require-typed-ref | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/require-typed-ref.md for this page in Markdown format

# vue/require-typed-ref Style [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#vue-require-typed-ref)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#what-it-does)

Require `ref` and `shallowRef` functions to be strongly typed.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#why-is-this-bad)

With TypeScript it is easy to prevent usage of `any` by using `noImplicitAny`. Unfortunately this rule is easily bypassed with Vue `ref()` function. Calling `ref()` function without a generic parameter or an initial value leads to ref having `Ref<any>` type.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
const count = ref();
const name = shallowRef();
```

Examples of **correct** code for this rule:

typescript

```
const count = ref<number>();
const a = ref(0);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/require-typed-ref": "error"
  }
}
```

bash

```
oxlint --deny vue/require-typed-ref --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/require_typed_ref.rs)
