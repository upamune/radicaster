---
title: "unicorn/prefer-global-this | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-global-this.md for this page in Markdown format

# unicorn/prefer-global-this Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#unicorn-prefer-global-this)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#what-it-does)

Enforces the use of [`globalThis`](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/globalThis) instead of environment‑specific global object aliases (`window`, `self`, or `global`). Using the standard `globalThis` makes your code portable across browsers, Web Workers, Node.js, and future JavaScript runtimes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#why-is-this-bad)

• **Portability** – `window` is only defined in browser main threads, `self` is used in Web Workers, and `global` is Node‑specific. Choosing the wrong alias causes runtime crashes when the code is executed outside of its original environment. • **Clarity** – `globalThis` clearly communicates that you are referring to the global object itself rather than a particular platform.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#examples)

Examples of **incorrect** code for this rule:

js

```
// Browser‑only
window.alert("Hi");

// Node‑only
if (typeof global.Buffer !== "undefined") {
}

// Web Worker‑only
self.postMessage("done");
```

Examples of **correct** code for this rule:

js

```
globalThis.alert("Hi");

if (typeof globalThis.Buffer !== "undefined") {
}

globalThis.postMessage("done");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-global-this": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-global-this
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_global_this.rs)
