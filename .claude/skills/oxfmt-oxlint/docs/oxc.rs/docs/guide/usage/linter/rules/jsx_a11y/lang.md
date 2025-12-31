---
title: "jsx_a11y/lang | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/lang.md for this page in Markdown format

# jsx\_a11y/lang Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#jsx-a11y-lang)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#what-it-does)

The lang prop on the `<html>` element must be a valid IETF's BCP 47 language tag.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#why-is-this-bad)

If the language of a webpage is not specified as valid, the screen reader assumes the default language set by the user. Language settings become an issue for users who speak multiple languages and access website in more than one language.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<html>
<html lang="foo">
```

Examples of **correct** code for this rule:

jsx

```
<html lang="en">
<html lang="en-US">
```

### Resources [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#resources)

* [eslint-plugin-jsx-a11y/lang](https://github.com/jsx-eslint/eslint-plugin-jsx-a11y/blob/v6.9.0/docs/rules/lang.md)
* [IANA Language Subtag Registry](https://www.iana.org/assignments/language-subtag-registry/language-subtag-registry)

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/lang": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/lang --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/lang.rs)
