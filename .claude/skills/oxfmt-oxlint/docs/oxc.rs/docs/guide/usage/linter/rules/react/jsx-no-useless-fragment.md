---
title: "react/jsx-no-useless-fragment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.md for this page in Markdown format

# react/jsx-no-useless-fragment Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#react-jsx-no-useless-fragment)

💡 A suggestion is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#what-it-does)

Disallow unnecessary fragments.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#why-is-this-bad)

Fragments are a useful tool when you need to group multiple children without adding a node to the DOM tree. However, sometimes you might end up with a fragment with a single child. When this child is an element, string, or expression, it's not necessary to use a fragment.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<>foo</>
<div><>foo</></div>
```

Examples of **correct** code for this rule:

jsx

```
<>foo <div></div></>
<div>foo</div>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#configuration)

This rule accepts a configuration object with the following properties:

### allowExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#allowexpressions)

type: `boolean`

default: `false`

Allow fragments with a single expression child.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-useless-fragment": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-useless-fragment --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_useless_fragment.rs)
