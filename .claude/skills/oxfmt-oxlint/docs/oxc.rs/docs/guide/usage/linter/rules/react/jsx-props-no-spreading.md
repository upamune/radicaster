---
title: "react/jsx-props-no-spreading | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-props-no-spreading.md for this page in Markdown format

# react/jsx-props-no-spreading Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#react-jsx-props-no-spreading)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#what-it-does)

Disallow JSX prop spreading

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#why-is-this-bad)

Enforces that there is no spreading for any JSX attribute. This enhances readability of code by being more explicit about what props are received by the component. It is also good for maintainability by avoiding passing unintentional extra props and allowing react to emit warnings when invalid HTML props are passed to HTML elements.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<App {...props} />
<MyCustomComponent {...props} some_other_prop={some_other_prop} />
<img {...props} />
```

Examples of **correct** code for this rule:

jsx

```
const {src, alt} = props;
const {one_prop, two_prop} = otherProps;
<MyCustomComponent one_prop={one_prop} two_prop={two_prop} />
<img src={src} alt={alt} />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#configuration)

This rule accepts a configuration object with the following properties:

### custom [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#custom)

type: `"ignore" | "enforce"`

default: `"enforce"`

`custom` set to `ignore` will ignore all custom jsx tags like `App`, `MyCustomComponent` etc. Default is set to `enforce`.

### exceptions [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#exceptions)

type: `string[]`

default: `[]`

Exceptions flip the enforcement behavior for specific components. For example:

* If `html` is set to `ignore`, an exception for `div` will enforce the rule on `<div>` elements.
* If `custom` is set to `enforce`, an exception for `Foo` will ignore the rule on `<Foo>` components.

This allows you to override the general setting for individual components.

### explicitSpread [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#explicitspread)

type: `"ignore" | "enforce"`

default: `"enforce"`

`explicitSpread` set to `ignore` will ignore spread operators that are explicitly listing all object properties within that spread. Default is set to `enforce`.

### html [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#html)

type: `"ignore" | "enforce"`

default: `"enforce"`

`html` set to `ignore` will ignore all html jsx tags like `div`, `img` etc. Default is set to `enforce`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-props-no-spreading": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-props-no-spreading --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_props_no_spreading.rs)
