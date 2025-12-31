---
title: "unicorn/prefer-dom-node-remove | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.md for this page in Markdown format

# unicorn/prefer-dom-node-remove Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#unicorn-prefer-dom-node-remove)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#what-it-does)

Prefers the use of `child.remove()` over `parentNode.removeChild(child)`.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#why-is-this-bad)

The DOM function [`Node#remove()`](https://developer.mozilla.org/en-US/docs/Web/API/ChildNode/remove) is preferred over the indirect removal of an object with [`Node#removeChild()`](https://developer.mozilla.org/en-US/docs/Web/API/Node/removeChild).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
parentNode.removeChild(childNode);
```

Examples of **correct** code for this rule:

javascript

```
childNode.remove();
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-dom-node-remove": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-dom-node-remove
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_dom_node_remove.rs)
