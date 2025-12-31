---
title: "unicorn/prefer-add-event-listener | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.md for this page in Markdown format

# unicorn/prefer-add-event-listener Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#unicorn-prefer-add-event-listener)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#what-it-does)

Enforces the use of `.addEventListener()` and `.removeEventListener()` over their `on`-function counterparts.

For example, `foo.addEventListener('click', handler);` is preferred over `foo.onclick = handler;` for HTML DOM Events.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#why-is-this-bad)

There are [numerous advantages of using `addEventListener`](https://stackoverflow.com/questions/6348494/addeventlistener-vs-onclick/35093997#35093997). Some of these advantages include registering unlimited event handlers and optionally having the event handler invoked only once.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.onclick = () => {};
```

Examples of **correct** code for this rule:

javascript

```
foo.addEventListener("click", () => {});
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-add-event-listener": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-add-event-listener
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_add_event_listener.rs)
