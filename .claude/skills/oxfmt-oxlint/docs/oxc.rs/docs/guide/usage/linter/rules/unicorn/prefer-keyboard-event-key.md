---
title: "unicorn/prefer-keyboard-event-key | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.md for this page in Markdown format

# unicorn/prefer-keyboard-event-key Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#unicorn-prefer-keyboard-event-key)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#what-it-does)

Enforces the use of [`KeyboardEvent#key`](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key) over [`KeyboardEvent#keyCode`](https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/keyCode) which is deprecated. The `.key` property is also more semantic and readable.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#why-is-this-bad)

The `keyCode`, `which`, and `charCode` properties are deprecated and should be avoided in favor of the `key` property.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#examples)

Examples of **incorrect** code for this rule:

js

```
window.addEventListener("keydown", (event) => {
  if (event.keyCode === 8) {
    console.log("Backspace was pressed");
  }
});

window.addEventListener("keydown", (event) => {
  console.log(event.keyCode);
});
```

Examples of **correct** code for this rule:

js

```
window.addEventListener("keydown", (event) => {
  if (event.key === "Backspace") {
    console.log("Backspace was pressed");
  }
});

window.addEventListener("click", (event) => {
  console.log(event.key);
});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-keyboard-event-key": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-keyboard-event-key
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_keyboard_event_key.rs)
