---
title: "react/jsx-fragments | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-fragments.md for this page in Markdown format

# react/jsx-fragments Style [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#react-jsx-fragments)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#what-it-does)

Enforces the shorthand or standard form for React Fragments.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#why-is-this-bad)

Makes code using fragments more consistent one way or the other.

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#configuration)

This rule accepts one of the following string values:

### `"syntax"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#syntax)

This is the default mode. It will enforce the shorthand syntax for React fragments, with one exception. Keys or attributes are not supported by the shorthand syntax, so the rule will not warn on standard-form fragments that use those.

Examples of **incorrect** code for this rule:

jsx

```
<React.Fragment>
  <Foo />
</React.Fragment>
```

Examples of **correct** code for this rule:

jsx

```
<>
  <Foo />
</>
```

jsx

```
<React.Fragment key="key">
  <Foo />
</React.Fragment>
```

### `"element"` [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#element)

This mode enforces the standard form for React fragments.

Examples of **incorrect** code for this rule:

jsx

```
<>
  <Foo />
</>
```

Examples of **correct** code for this rule:

jsx

```
<React.Fragment>
  <Foo />
</React.Fragment>
```

jsx

```
<React.Fragment key="key">
  <Foo />
</React.Fragment>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-fragments": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-fragments --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_fragments.rs)
