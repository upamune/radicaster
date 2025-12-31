---
title: "eslint/no-labels | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-labels.md for this page in Markdown format

# eslint/no-labels Style [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#eslint-no-labels)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#what-it-does)

Disallow labeled statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#why-is-this-bad)

Labeled statements in JavaScript are used in conjunction with `break` and `continue` to control flow around multiple loops. For example:

js

```
outer: while (true) {
  while (true) {
    break outer;
  }
}
```

The `break outer` statement ensures that this code will not result in an infinite loop because control is returned to the next statement after the `outer` label was applied. If this statement was changed to be just `break`, control would flow back to the outer `while` statement and an infinite loop would result. While convenient in some cases, labels tend to be used only rarely and are frowned upon by some as a remedial form of flow control that is more error prone and harder to understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#examples)

Examples of **incorrect** code for this rule:

js

```
label: while (true) {
  // ...
}

label: while (true) {
  break label;
}

label: while (true) {
  continue label;
}

label: switch (a) {
  case 0:
    break label;
}

label: {
  break label;
}

label: if (a) {
  break label;
}
```

Examples of **correct** code for this rule:

js

```
var f = {
  label: "foo",
};

while (true) {
  break;
}

while (true) {
  continue;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#configuration)

This rule accepts a configuration object with the following properties:

### allowLoop [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#allowloop)

type: `boolean`

default: `false`

If set to `true`, this rule ignores labels which are sticking to loop statements. Examples of **correct** code with this option set to `true`:

js

```
label: while (true) {
  break label;
}
```

### allowSwitch [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#allowswitch)

type: `boolean`

default: `false`

If set to `true`, this rule ignores labels which are sticking to switch statements. Examples of **correct** code with this option set to `true`:

js

```
label: switch (a) {
  case 0:
    break label;
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-labels": "error"
  }
}
```

bash

```
oxlint --deny no-labels
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_labels.rs)
