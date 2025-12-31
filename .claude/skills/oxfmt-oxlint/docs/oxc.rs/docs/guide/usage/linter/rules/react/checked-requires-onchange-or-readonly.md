---
title: "react/checked-requires-onchange-or-readonly | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.md for this page in Markdown format

# react/checked-requires-onchange-or-readonly Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#react-checked-requires-onchange-or-readonly)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#what-it-does)

This rule enforces onChange or readonly attribute for checked property of input elements. It also warns when checked and defaultChecked properties are used together.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<input type="checkbox" checked />
<input type="checkbox" checked defaultChecked />
<input type="radio" checked defaultChecked />

React.createElement('input', { checked: false });
React.createElement('input', { type: 'checkbox', checked: true });
React.createElement('input', { type: 'checkbox', checked: true, defaultChecked: true });
```

Examples of **correct** code for this rule:

jsx

```
<input type="checkbox" checked onChange={() => {}} />
<input type="checkbox" checked readOnly />
<input type="checkbox" checked onChange readOnly />
<input type="checkbox" defaultChecked />

React.createElement('input', { type: 'checkbox', checked: true, onChange() {} });
React.createElement('input', { type: 'checkbox', checked: true, readOnly: true });
React.createElement('input', { type: 'checkbox', checked: true, onChange() {}, readOnly: true });
React.createElement('input', { type: 'checkbox', defaultChecked: true });
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#configuration)

This rule accepts a configuration object with the following properties:

### ignoreExclusiveCheckedAttribute [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#ignoreexclusivecheckedattribute)

type: `boolean`

default: `false`

Ignore the restriction that `checked` and `defaultChecked` should not be used together.

### ignoreMissingProperties [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#ignoremissingproperties)

type: `boolean`

default: `false`

Ignore the requirement to provide either `onChange` or `readOnly` when the `checked` prop is present.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/checked-requires-onchange-or-readonly": "error"
  }
}
```

bash

```
oxlint --deny react/checked-requires-onchange-or-readonly --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/checked_requires_onchange_or_readonly.rs)
