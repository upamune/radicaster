---
title: "react/jsx-filename-extension | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-filename-extension.md for this page in Markdown format

# react/jsx-filename-extension Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#react-jsx-filename-extension)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#what-it-does)

Enforces consistent use of the `.jsx` file extension.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#why-is-this-bad)

Some bundlers or parsers need to know by the file extension that it contains JSX in order to properly handle the files.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
// filename: MyComponent.js
function MyComponent() {
  return <div />;
}
```

Examples of **correct** code for this rule:

jsx

```
// filename: MyComponent.jsx
function MyComponent() {
  return <div />;
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#allow)

type: `"always" | "as-needed"`

default: `"always"`

When to allow a JSX filename extension. By default all files may have a JSX extension. Set this to `as-needed` to only allow JSX file extensions in files that contain JSX syntax.

### extensions [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#extensions)

type: `string[]`

default: `["jsx"]`

The set of allowed file extensions. Can include or exclude the leading dot (e.g., "jsx" and ".jsx" are both valid).

### ignoreFilesWithoutCode [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#ignorefileswithoutcode)

type: `boolean`

default: `false`

If enabled, files that do not contain code (i.e. are empty, contain only whitespaces or comments) will not be rejected.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-filename-extension": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-filename-extension --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_filename_extension.rs)
