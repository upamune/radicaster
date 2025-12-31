---
title: "typescript/no-meaningless-void-operator | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.md for this page in Markdown format

# typescript/no-meaningless-void-operator Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#typescript-no-meaningless-void-operator)

✅ This rule is turned on by default when type-aware linting is enabled.

💭 This rule requires [type information](https://oxc.rs/docs/guide/usage/linter/type-aware.html).

🚧 An auto-fix is planned for this rule, but not implemented at this time.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#what-it-does)

This rule disallows the void operator when its argument is already of type void or undefined.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#why-is-this-bad)

The void operator is useful when you want to execute an expression and force it to evaluate to undefined. However, using void on expressions that are already of type void or undefined is meaningless and adds unnecessary complexity to the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#examples)

Examples of **incorrect** code for this rule:

ts

```
function foo(): void {
  return;
}

void foo(); // meaningless, foo() already returns void

void undefined; // meaningless, undefined is already undefined

async function bar() {
  void (await somePromise); // meaningless if somePromise resolves to void
}
```

Examples of **correct** code for this rule:

ts

```
function getValue(): number {
  return 42;
}

void getValue(); // meaningful, converts number to void

void console.log("hello"); // meaningful, console.log returns undefined but we want to be explicit

function processData() {
  // some processing
}

processData(); // no void needed since we don't care about return value
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#configuration)

This rule accepts a configuration object with the following properties:

### checkNever [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#checknever)

type: `boolean`

default: `false`

Whether to check `void` applied to expressions of type `never`.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/no-meaningless-void-operator": "error"
  }
}
```

bash

```
oxlint --type-aware --deny typescript/no-meaningless-void-operator
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/no_meaningless_void_operator.rs)
* [Rule Source (tsgolint)](https://github.com/oxc-project/tsgolint/blob/main/internal/rules/no_meaningless_void_operator/no_meaningless_void_operator.go)
