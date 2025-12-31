---
title: "eslint/no-console | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-console.md for this page in Markdown format

# eslint/no-console Restriction [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#eslint-no-console)

💡 A suggestion is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#what-it-does)

Disallow the use of console.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#why-is-this-bad)

In JavaScript that is designed to be executed in the browser, it’s considered a best practice to avoid using methods on console. Such messages are considered to be for debugging purposes and therefore not suitable to ship to the client. In general, calls using console should be stripped before being pushed to production.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
console.log("Log a debug level message.");
console.warn("Log a warn level message.");
console.error("Log an error level message.");
console.log = foo();
```

Examples of **correct** code for this rule:

javascript

```
// custom console
Console.log("Hello world!");
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#configuration)

This rule accepts a configuration object with the following properties:

### allow [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#allow)

type: `string[]`

default: `[]`

The `allow` option permits the given list of console methods to be used as exceptions to this rule.

Say the option was configured as `{ "allow": ["info"] }` then the rule would behave as follows:

Example of **incorrect** code for this option:

javascript

```
console.log("foo");
```

Example of **correct** code for this option:

javascript

```
console.info("foo");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-console": "error"
  }
}
```

bash

```
oxlint --deny no-console
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_console.rs)
