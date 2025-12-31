---
title: "typescript/ban-ts-comment | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/ban-ts-comment.md for this page in Markdown format

# typescript/ban-ts-comment Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#typescript-ban-ts-comment)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#what-it-does)

This rule lets you set which directive comments you want to allow in your codebase.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#why-is-this-bad)

Using TypeScript directives to suppress TypeScript compiler errors reduces the effectiveness of TypeScript overall.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#examples)

Examples of **incorrect** code for this rule:

ts

```
if (false) {
  // @ts-ignore: Unreachable code error
  console.log("hello");
}
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#configuration)

This rule allows you to specify how different TypeScript directive comments should be handled.

For each directive (`@ts-expect-error`, `@ts-ignore`, `@ts-nocheck`, `@ts-check`), you can choose one of the following options:

* `true`: Disallow the directive entirely, preventing its use in the entire codebase.
* `false`: Allow the directive without any restrictions.
* `"allow-with-description"`: Allow the directive only if it is followed by a description explaining its use. The description must meet the minimum length specified by `minimumDescriptionLength`.
* `{ "descriptionFormat": "<regex>" }`: Allow the directive only if the description matches the specified regex pattern.

For example:

json

```
{
  "ts-expect-error": "allow-with-description",
  "ts-ignore": true,
  "ts-nocheck": { "descriptionFormat": "^: TS\\d+ because .+$" },
  "ts-check": false,
  "minimumDescriptionLength": 3
}
```

This rule accepts a configuration object with the following properties:

### minimumDescriptionLength [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#minimumdescriptionlength)

type: `integer`

default: `3`

Minimum description length required when using directives with `allow-with-description`.

### ts-check [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#ts-check)

How to handle the `@ts-check` directive.

### ts-expect-error [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#ts-expect-error)

How to handle the `@ts-expect-error` directive.

### ts-ignore [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#ts-ignore)

How to handle the `@ts-ignore` directive.

### ts-nocheck [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#ts-nocheck)

How to handle the `@ts-nocheck` directive.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/ban-ts-comment": "error"
  }
}
```

bash

```
oxlint --deny typescript/ban-ts-comment
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/ban_ts_comment.rs)
