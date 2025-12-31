---
title: "unicorn/prefer-event-target | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-event-target.md for this page in Markdown format

# unicorn/prefer-event-target Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#unicorn-prefer-event-target)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#what-it-does)

Prefers `EventTarget` over `EventEmitter`.

This rule reduces the bundle size and makes your code more cross-platform friendly.

See the [differences](https://nodejs.org/api/events.html#eventtarget-and-event-api) between `EventEmitter` and `EventTarget`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#why-is-this-bad)

While [`EventEmitter`](https://nodejs.org/api/events.html#class-eventemitter) is only available in Node.js, [`EventTarget`](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget) is also available in *Deno* and browsers.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
class Foo extends EventEmitter {}
```

Examples of **correct** code for this rule:

javascript

```
class Foo extends OtherClass {}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-event-target": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-event-target
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_event_target.rs)
