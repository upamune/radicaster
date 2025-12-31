---
title: "react/forbid-elements | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/forbid-elements.md for this page in Markdown format

# react/forbid-elements Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#react-forbid-elements)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#what-it-does)

Allows you to configure a list of forbidden elements and to specify their desired replacements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#why-is-this-bad)

You may want to forbid usage of certain elements in favor of others, (e.g. forbid alland use  instead)

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// [1, { "forbid": ["button"] }]
<button />;
React.createElement("button");

// [1, { "forbid": ["Modal"] }]
<Modal />;
React.createElement(Modal);

// [1, { "forbid": ["Namespaced.Element"] }]
<Namespaced.Element />;
React.createElement(Namespaced.Element);

// [1, { "forbid": [{ "element": "button", "message": "use <Button> instead" }, "input"] }]
<div>
  <button />
  <input />
</div>;
React.createElement("div", {}, React.createElement("button", {}, React.createElement("input")));
```

Examples of **correct** code for this rule:

jsx

```
// [1, { "forbid": ["button"] }]
<Button />

// [1, { "forbid": [{ "element": "button" }] }]
<Button />
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/forbid-elements": "error"
  }
}
```

bash

```
oxlint --deny react/forbid-elements --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/forbid_elements.rs)
