---
title: "unicorn/no-invalid-remove-event-listener | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.md for this page in Markdown format

# unicorn/no-invalid-remove-event-listener Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#unicorn-no-invalid-remove-event-listener)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#what-it-does)

It warns when you use a non-function value as the second argument of `removeEventListener`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#why-is-this-bad)

The [`removeEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener) function must be called with a reference to the same function that was passed to [`addEventListener`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener). Calling `removeEventListener` with an inline function or the result of an inline `.bind()` call is indicative of an error, and won't actually remove the listener.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
el.removeEventListener("click", () => {});
el.removeEventListener("click", function () {});
```

Examples of **correct** code for this rule:

javascript

```
el.removeEventListener("click", handler);
el.removeEventListener("click", handler.bind(this));
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/no-invalid-remove-event-listener": "error"
  }
}
```

bash

```
oxlint --deny unicorn/no-invalid-remove-event-listener
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/no_invalid_remove_event_listener.rs)
