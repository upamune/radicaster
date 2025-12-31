---
title: "unicorn/prefer-classlist-toggle | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.md for this page in Markdown format

# unicorn/prefer-classlist-toggle Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#unicorn-prefer-classlist-toggle)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#what-it-does)

Prefers the use of `element.classList.toggle(className, condition)` over conditional add/remove patterns.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#why-is-this-bad)

The `toggle()` method is more concise and expressive than using conditional logic to switch between `add()` and `remove()`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
if (condition) {
  element.classList.add("className");
} else {
  element.classList.remove("className");
}

condition ? element.classList.add("className") : element.classList.remove("className");

element.classList[condition ? "add" : "remove"](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/"className");
```

Examples of **correct** code for this rule:

javascript

```
element.classList.toggle("className", condition);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-classlist-toggle": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-classlist-toggle
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_classlist_toggle.rs)
