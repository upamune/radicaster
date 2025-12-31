---
title: "eslint/no-extra-boolean-cast | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.md for this page in Markdown format

# eslint/no-extra-boolean-cast Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#eslint-no-extra-boolean-cast)

✅ This rule is turned on by default.

🛠️💡 An auto-fix and a suggestion are available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#what-it-does)

This rule disallows unnecessary boolean casts.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#why-is-this-bad)

In contexts such as an if statement's test where the result of the expression will already be coerced to a Boolean, casting to a Boolean via double negation (`!!`) or a `Boolean` call is unnecessary.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
var foo = !!!bar;
var foo = Boolean(!!bar);

if (!!foo) {
}
if (Boolean(foo)) {
}

// with "enforceForInnerExpressions" option enabled
if (!!foo || bar) {
}
```

Examples of **correct** code for this rule:

javascript

```
var foo = !bar;
var foo = Boolean(bar);

if (foo) {
}
if (foo) {
}

// with "enforceForInnerExpressions" option enabled
if (foo || bar) {
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#configuration)

This rule accepts a configuration object with the following properties:

### enforceForInnerExpressions [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#enforceforinnerexpressions)

type: `boolean`

default: `false`

when set to `true`, in addition to checking default contexts, checks whether extra boolean casts are present in expressions whose result is used in a boolean context. See examples below. Default is `false`, meaning that this rule by default does not warn about extra booleans cast inside inner expressions.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-extra-boolean-cast": "error"
  }
}
```

bash

```
oxlint --deny no-extra-boolean-cast
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_extra_boolean_cast.rs)
