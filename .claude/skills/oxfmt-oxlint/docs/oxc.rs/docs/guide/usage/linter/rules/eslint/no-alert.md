---
title: "eslint/no-alert | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-alert.md for this page in Markdown format

# eslint/no-alert Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#eslint-no-alert)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#what-it-does)

Disallow the use of alert, confirm, and prompt

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#why-is-this-bad)

JavaScript’s alert, confirm, and prompt functions are widely considered to be obtrusive as UI elements and should be replaced by a more appropriate custom UI implementation. Furthermore, alert is often used while debugging code, which should be removed before deployment to production.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#examples)

Examples of **incorrect** code for this rule:

js

```
alert("here!");

confirm("Are you sure?");

prompt("What's your name?", "John Doe");
```

Examples of **correct** code for this rule:

js

```
customAlert("Something happened!");

customConfirm("Are you sure?");

customPrompt("Who are you?");

function foo() {
  var alert = myCustomLib.customAlert;
  alert();
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-alert": "error"
  }
}
```

bash

```
oxlint --deny no-alert
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_alert.rs)
