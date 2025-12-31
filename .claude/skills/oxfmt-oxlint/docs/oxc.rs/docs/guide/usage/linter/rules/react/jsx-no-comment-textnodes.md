---
title: "react/jsx-no-comment-textnodes | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.md for this page in Markdown format

# react/jsx-no-comment-textnodes Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#react-jsx-no-comment-textnodes)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#what-it-does)

This rule prevents comment strings (e.g. beginning with `//` or `/*`) from being accidentally injected as a text node in JSX statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#why-is-this-bad)

In JSX, any text node that is not wrapped in curly braces is considered a literal string to be rendered. This can lead to unexpected behavior when the text contains a comment.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
const Hello = () => {
  return <div>// empty div</div>;
};

const Hello = () => {
  return <div>/* empty div */</div>;
};
```

Examples of **correct** code for this rule:

jsx

```
const Hello = () => {
  return <div>{/* empty div */}</div>;
};
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-comment-textnodes": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-comment-textnodes --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_comment_textnodes.rs)
