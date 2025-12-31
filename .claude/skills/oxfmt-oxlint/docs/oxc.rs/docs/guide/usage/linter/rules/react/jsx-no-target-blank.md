---
title: "react/jsx-no-target-blank | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/react/jsx-no-target-blank.md for this page in Markdown format

# react/jsx-no-target-blank Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#react-jsx-no-target-blank)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#what-it-does)

This rule aims to prevent user generated link hrefs and form actions from creating security vulnerabilities by requiring `rel='noreferrer'` for external link hrefs and form actions, and optionally any dynamically generated link hrefs and form actions.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#why-is-this-bad)

When creating a JSX element that has an `a` tag, it is often desired to have the link open in a new tab using the `target='_blank'` attribute. Using this attribute unaccompanied by `rel='noreferrer'`, however, is a severe security vulnerability (see [`noreferrer` docs](https://html.spec.whatwg.org/multipage/links.html#link-type-noreferrer) and [`noopener` docs](https://html.spec.whatwg.org/multipage/links.html#link-type-noopener) for more details). This rules requires that you accompany `target='_blank'` attributes with `rel='noreferrer'`.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
var Hello = <a target="_blank" href="https://example.com/"></a>;
var Hello = <a target="_blank" href={dynamicLink}></a>;
```

Examples of **correct** code for this rule:

jsx

```
/// correct
var Hello = <p target="_blank"></p>;
var Hello = <a target="_blank" rel="noreferrer" href="https://example.com"></a>;
var Hello = <a target="_blank" rel="noopener noreferrer" href="https://example.com"></a>;
var Hello = <a target="_blank" href="relative/path/in/the/host"></a>;
var Hello = <a target="_blank" href="/absolute/path/in/the/host"></a>;
var Hello = <a></a>;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#configuration)

This rule accepts a configuration object with the following properties:

### allowReferrer [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#allowreferrer)

type: `boolean`

default: `false`

Whether to allow referrers.

### enforceDynamicLinks [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#enforcedynamiclinks)

type: `"always" | "never"`

default: `"always"`

Whether to enforce dynamic links or enforce static links.

### forms [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#forms)

type: `boolean`

default: `false`

Whether to check form elements.

### links [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#links)

type: `boolean`

default: `true`

Whether to check link elements.

### warnOnSpreadAttributes [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#warnonspreadattributes)

type: `boolean`

default: `false`

Whether to warn when spread attributes are used.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["react"],
  "rules": {
    "react/jsx-no-target-blank": "error"
  }
}
```

bash

```
oxlint --deny react/jsx-no-target-blank --react-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/react/jsx_no_target_blank.rs)
