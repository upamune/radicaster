---
title: "react/no-unescaped-entities | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/no-unescaped-entities.md for this page in Markdown format

# react/no-unescaped-entities Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#react-no-unescaped-entities)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#what-it-does)

This rule prevents characters that you may have meant as JSX escape characters from being accidentally injected as a text node in JSX statements.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#why-is-this-bad)

JSX escape characters are used to inject characters into JSX statements that would otherwise be interpreted as code.

### Example [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#example)

Incorrect

jsx

```
<div> > </div>
```

Correct

jsx

```
<div> &gt; </div>
```

jsx

```
<div> {">"} </div>
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/no-unescaped-entities": "error"
  }
}
```

bash

```
oxlint --deny react/no-unescaped-entities --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/no_unescaped_entities.rs)
