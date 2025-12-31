---
title: "eslint/no-useless-backreference | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-backreference.md for this page in Markdown format

# eslint/no-useless-backreference Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#eslint-no-useless-backreference)

✅ This rule is turned on by default.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#what-it-does)

Disallows backreferences in regular expressions that will always be ignored because the capture group they refer to has not matched and cannot match at the time the backreference is evaluated.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#why-is-this-bad)

Useless backreferences can lead to confusing or misleading regular expressions. They may give the impression that a group’s value is being reused, but due to the structure of the pattern (e.g., order of evaluation, disjunctions, or negative lookarounds), the group has not matched anything — so the reference always resolves to an empty string. This is almost always a mistake and makes patterns harder to understand and maintain.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#examples)

Examples of **incorrect** code for this rule:

js

```
/\1(a)/; // backreference appears before group
/(a|\1b)/; // group and reference are in different alternatives
/(?<=\1(a))b/; // backreference used before group in lookbehind
/\1(?!(a))/; // group is inside negative lookahead
/(a\1)/; // backreference is inside its own group
```

Examples of **correct** code for this rule:

js

```
/(a)\1/; // valid — backreference follows completed group
/(?<name>a)\k<name>/; // named group used properly
/(?:a|(b))\1/; // backreference only used when group matches
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-backreference": "error"
  }
}
```

bash

```
oxlint --deny no-useless-backreference
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_backreference.rs)
