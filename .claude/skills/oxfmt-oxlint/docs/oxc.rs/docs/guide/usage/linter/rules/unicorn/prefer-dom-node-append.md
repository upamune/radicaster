---
title: "unicorn/prefer-dom-node-append | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.md for this page in Markdown format

# unicorn/prefer-dom-node-append Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#unicorn-prefer-dom-node-append)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#what-it-does)

Enforces the use of, for example, `document.body.append(div);` over `document.body.appendChild(div);` for DOM nodes.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#why-is-this-bad)

There are [some advantages of using `Node#append()`](https://developer.mozilla.org/en-US/docs/Web/API/ParentNode/append), like the ability to append multiple nodes and to append both [`DOMString`](https://developer.mozilla.org/en-US/docs/Web/API/DOMString) and DOM node objects.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
foo.appendChild(bar);
```

Examples of **correct** code for this rule:

javascript

```
foo.append(bar);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-dom-node-append": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-dom-node-append
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_dom_node_append.rs)
