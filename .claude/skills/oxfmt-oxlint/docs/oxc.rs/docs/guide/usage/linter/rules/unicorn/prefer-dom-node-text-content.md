---
title: "unicorn/prefer-dom-node-text-content | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.md for this page in Markdown format

# unicorn/prefer-dom-node-text-content Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#unicorn-prefer-dom-node-text-content)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#what-it-does)

Enforces the use of `.textContent` over `.innerText` for DOM nodes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#why-is-this-bad)

There are some disadvantages of using .innerText.

* `.innerText` is much more performance-heavy as it requires layout information to return the result.
* `.innerText` is defined only for HTMLElement objects, while `.textContent` is defined for all Node objects.
* `.innerText` is not standard, for example, it is not present in Firefox.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
const text = foo.innerText;
```

Examples of **correct** code for this rule:

javascript

```
const text = foo.textContent;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-dom-node-text-content": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-dom-node-text-content
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_dom_node_text_content.rs)
