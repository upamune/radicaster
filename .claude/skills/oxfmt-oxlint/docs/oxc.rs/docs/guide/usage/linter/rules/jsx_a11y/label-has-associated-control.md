---
title: "jsx_a11y/label-has-associated-control | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/label-has-associated-control.md for this page in Markdown format

# jsx\_a11y/label-has-associated-control Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#jsx-a11y-label-has-associated-control)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#what-it-does)

Enforce that a label tag has a text label and an associated control.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#why-is-this-bad)

A form label that either isn't properly associated with a form control (such as an `<input>`), or doesn't contain accessible text, hinders accessibility for users using assistive technologies such as screen readers. The user may not have enough information to understand the purpose of the form control.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
function Foo(props) {
    return <label {...props} />
}

<input type="text" />
<label>Surname</label>
```

Examples of **correct** code for this rule:

jsx

```
function Foo(props) {
  const { htmlFor, ...otherProps } = props;

  return <label htmlFor={htmlFor} {...otherProps} />;
}

<label>
  <input type="text" />
  Surname
</label>;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#configuration)

This rule accepts a configuration object with the following properties:

### assert [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#assert)

type: `"htmlFor" | "nesting" | "both" | "either"`

default: `"either"`

The type of association required between the label and the control.

### controlComponents [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#controlcomponents)

type: `string[]`

default: `[]`

Custom JSX components to be treated as form controls.

### depth [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#depth)

type: `integer`

default: `2`

Maximum depth to search for a nested control.

### labelAttributes [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#labelattributes)

type: `string[]`

default: `["alt", "aria-label", "aria-labelledby"]`

Attributes to check for accessible label text.

### labelComponents [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#labelcomponents)

type: `string[]`

default: `["label"]`

Custom JSX components to be treated as labels.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/label-has-associated-control": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/label-has-associated-control --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/label_has_associated_control.rs)
