---
title: "unicorn/prefer-dom-node-dataset | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.md for this page in Markdown format

# unicorn/prefer-dom-node-dataset Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#unicorn-prefer-dom-node-dataset)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#what-it-does)

Use [`.dataset`](https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/dataset) on DOM elements over `getAttribute(…)`, `.setAttribute(…)`, `.removeAttribute(…)` and `.hasAttribute(…)`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#why-is-this-bad)

The `dataset` property is a map of strings that contains all the `data-*` attributes from the element. It is a convenient way to access all of them at once.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
element.setAttribute("data-unicorn", "🦄");
```

Examples of **correct** code for this rule:

javascript

```
element.dataset.unicorn = "🦄";
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-dom-node-dataset": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-dom-node-dataset
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_dom_node_dataset.rs)
