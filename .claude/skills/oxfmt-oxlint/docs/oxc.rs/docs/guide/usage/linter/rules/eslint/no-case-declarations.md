---
title: "eslint/no-case-declarations | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-case-declarations.md for this page in Markdown format

# eslint/no-case-declarations Pedantic [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#eslint-no-case-declarations)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#what-it-does)

Disallow lexical declarations in case clauses.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#why-is-this-bad)

The reason is that the lexical declaration is visible in the entire switch block but it only gets initialized when it is assigned, which will only happen if the case where it is defined is reached.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
switch (foo) {
  case 1:
    let x = 1;
    break;
  case 2:
    const y = 2;
    break;
  case 3:
    function f() {}
    break;
  default:
    class C {}
}
```

Examples of **correct** code for this rule:

javascript

```
switch (foo) {
  case 1: {
    let x = 1;
    break;
  }
  case 2: {
    const y = 2;
    break;
  }
  case 3: {
    function f() {}
    break;
  }
  default: {
    class C {}
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-case-declarations": "error"
  }
}
```

bash

```
oxlint --deny no-case-declarations
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_case_declarations.rs)
