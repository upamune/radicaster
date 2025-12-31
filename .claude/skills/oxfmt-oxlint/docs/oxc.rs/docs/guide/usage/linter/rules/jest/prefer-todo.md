---
title: "jest/prefer-todo | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jest/prefer-todo.md for this page in Markdown format

# jest/prefer-todo Style [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#jest-prefer-todo)

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#what-it-does)

When test cases are empty then it is better to mark them as `test.todo` as it will be highlighted in the summary output.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#why-is-this-bad)

This rule triggers a warning if empty test cases are used without 'test.todo'.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
test("i need to write this test"); // invalid
test("i need to write this test", () => {}); // invalid
test.skip("i need to write this test", () => {}); // invalid
```

Examples of **correct** code for this rule:

javascript

```
test.todo("i need to write this test");
```

This rule is compatible with [eslint-plugin-vitest](https://github.com/vitest-dev/eslint-plugin-vitest/blob/main/docs/rules/prefer-todo.md), to use it, add the following configuration to your `.oxlintrc.json`:

json

```
{
  "rules": {
    "vitest/prefer-todo": "error"
  }
}
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jest"],
  "rules": {
    "jest/prefer-todo": "error"
  }
}
```

bash

```
oxlint --deny jest/prefer-todo --jest-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jest/prefer_todo.rs)
