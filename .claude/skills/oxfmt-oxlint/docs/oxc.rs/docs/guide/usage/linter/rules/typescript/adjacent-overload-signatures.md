---
title: "typescript/adjacent-overload-signatures | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.md for this page in Markdown format

# typescript/adjacent-overload-signatures Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#typescript-adjacent-overload-signatures)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#what-it-does)

Require that function overload signatures be consecutive.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#why-is-this-bad)

Function overload signatures represent multiple ways a function can be called, potentially with different return types. It's typical for an interface or type alias describing a function to place all overload signatures next to each other. If Signatures placed elsewhere in the type are easier to be missed by future developers reading the code.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#examples)

Examples of **incorrect** code for this rule:

typescript

```
declare namespace Foo {
  export function foo(s: string): void;
  export function foo(n: number): void;
  export function bar(): void;
  export function foo(sn: string | number): void;
}

type Foo = {
  foo(s: string): void;
  foo(n: number): void;
  bar(): void;
  foo(sn: string | number): void;
};

interface Foo {
  foo(s: string): void;
  foo(n: number): void;
  bar(): void;
  foo(sn: string | number): void;
}

class Foo {
  foo(s: string): void;
  foo(n: number): void;
  bar(): void {}
  foo(sn: string | number): void {}
}

export function foo(s: string): void;
export function foo(n: number): void;
export function bar(): void;
export function foo(sn: string | number): void;
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/adjacent-overload-signatures": "error"
  }
}
```

bash

```
oxlint --deny typescript/adjacent-overload-signatures
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/adjacent_overload_signatures.rs)
