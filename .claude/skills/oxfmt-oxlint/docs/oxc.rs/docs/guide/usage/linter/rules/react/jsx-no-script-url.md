---
title: "react/jsx-no-script-url | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-script-url.md for this page in Markdown format

# react/jsx-no-script-url Suspicious [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#react-jsx-no-script-url)

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#what-it-does)

Disallow usage of `javascript:` URLs.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#why-is-this-bad)

URLs starting with `javascript:` are a dangerous attack surface because it’s easy to accidentally include unsanitized output in a tag like `<a href>` and create a security hole.

Starting in React 16.9, any URLs starting with `javascript:` log a warning.

In React 19, `javascript:` URLs are [disallowed entirely](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#other-breaking-changes).

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<a href="javascript:void(0)">Test</a>
```

Examples of **correct** code for this rule:

jsx

```
<Foo test="javascript:void(0)" />
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#configuration)

This rule accepts a configuration object with the following properties:

### components [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#components)

type: `Record<string, array>`

default: `{}`

Additional components to check.

### includeFromSettings [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#includefromsettings)

type: `boolean`

default: `false`

Whether to include components from settings.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-script-url": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-script-url --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_script_url.rs)
