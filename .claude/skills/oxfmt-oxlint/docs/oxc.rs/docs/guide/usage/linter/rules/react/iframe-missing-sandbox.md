---
title: "react/iframe-missing-sandbox | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/iframe-missing-sandbox.md for this page in Markdown format

# react/iframe-missing-sandbox Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#react-iframe-missing-sandbox)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#what-it-does)

Enforce sandbox attribute on iframe elements

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#why-is-this-bad)

The sandbox attribute enables an extra set of restrictions for the content in the iframe. Using sandbox attribute is considered a good security practice. To learn more about sandboxing, see [MDN's documentation on the `sandbox` attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe#sandbox).

This rule checks all React `<iframe>` elements and verifies that there is `sandbox` attribute and that it's value is valid. In addition to that it also reports cases where attribute contains `allow-scripts` and `allow-same-origin` at the same time as this combination allows the embedded document to remove the sandbox attribute and bypass the restrictions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<iframe />;
<iframe sandbox="invalid-value" />;
<iframe sandbox="allow-same-origin allow-scripts" />;
```

Examples of **correct** code for this rule:

jsx

```
<iframe sandbox="" />;
<iframe sandbox="allow-origin" />;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/iframe-missing-sandbox": "error"
  }
}
```

bash

```
oxlint --deny react/iframe-missing-sandbox --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/iframe_missing_sandbox.rs)
