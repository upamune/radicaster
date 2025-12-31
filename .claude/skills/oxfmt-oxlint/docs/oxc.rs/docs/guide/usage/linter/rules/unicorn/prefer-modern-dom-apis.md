---
title: "unicorn/prefer-modern-dom-apis | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.md for this page in Markdown format

# unicorn/prefer-modern-dom-apis Style [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#unicorn-prefer-modern-dom-apis)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#what-it-does)

Enforces the use of:

* `childNode.replaceWith(newNode)` over `parentNode.replaceChild(newNode, oldNode)`
* `referenceNode.before(newNode)` over `parentNode.insertBefore(newNode, referenceNode)`
* `referenceNode.before('text')` over `referenceNode.insertAdjacentText('beforebegin', 'text')`
* `referenceNode.before(newNode)` over `referenceNode.insertAdjacentElement('beforebegin', newNode)`

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#why-is-this-bad)

There are some advantages of using the newer DOM APIs, like:

* Traversing to the parent node is not necessary.
* Appending multiple nodes at once.
* Both `DOMString` and DOM node objects can be manipulated.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
oldChildNode.replaceWith(newChildNode);
```

Examples of **correct** code for this rule:

javascript

```
parentNode.replaceChild(newChildNode, oldChildNode);
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "unicorn/prefer-modern-dom-apis": "error"
  }
}
```

bash

```
oxlint --deny unicorn/prefer-modern-dom-apis
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/unicorn/prefer_modern_dom_apis.rs)
