---
title: "unicorn/prefer-query-selector | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-query-selector.md for this page in Markdown format

# unicorn/prefer-query-selector Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#unicorn-prefer-query-selector)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#what-it-does)

Prefer `.querySelector()` over `.getElementById()`, `.querySelectorAll()` over `.getElementsByClassName()` and `.getElementsByTagName()`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#why-is-this-bad)

* Using `.querySelector()` and `.querySelectorAll()` is more flexible and allows for more specific selectors.
* It's better to use the same method to query DOM elements. This helps keep consistency and it lends itself to future improvements (e.g. more specific selectors).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
document.getElementById("foo");
document.getElementsByClassName("foo bar");
document.getElementsByTagName("main");
document.getElementsByClassName(fn());
```

Examples of **correct** code for this rule:

javascript

```
document.querySelector("#foo");
document.querySelector(".bar");
document.querySelector("main #foo .bar");
document.querySelectorAll(".foo .bar");
document.querySelectorAll("li a");
document.querySelector("li").querySelectorAll("a");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-query-selector": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-query-selector
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_query_selector.rs)
