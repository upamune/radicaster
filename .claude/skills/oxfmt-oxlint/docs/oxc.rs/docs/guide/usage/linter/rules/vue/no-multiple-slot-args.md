---
title: "vue/no-multiple-slot-args | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/vue/no-multiple-slot-args.md for this page in Markdown format

# vue/no-multiple-slot-args Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#vue-no-multiple-slot-args)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#what-it-does)

Disallow passing multiple arguments to scoped slots.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#why-is-this-bad)

Users have to use the arguments in fixed order and cannot omit the ones they don't need. e.g. if you have a slot that passes in 5 arguments but the user actually only need the last 2 of them, they will have to declare all 5 just to use the last 2.

More information can be found in [vuejs/vue#9468](https://github.com/vuejs/vue/issues/9468#issuecomment-462210146)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#examples)

Examples of **incorrect** code for this rule:

vue

```
<script>
export default {
  render(h) {
    var children = this.$scopedSlots.default(foo, bar);
    var children = this.$scopedSlots.default(...foo);
  },
};
</script>
```

Examples of **correct** code for this rule:

vue

```
<script>
export default {
  render(h) {
    var children = this.$scopedSlots.default();
    var children = this.$scopedSlots.default(foo);
    var children = this.$scopedSlots.default({ foo, bar });
  },
};
</script>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["vue"],
  "rules": {
    "vue/no-multiple-slot-args": "error"
  }
}
```

bash

```
oxlint --deny vue/no-multiple-slot-args --vue-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/vue/no_multiple_slot_args.rs)
