---
title: "Isolated Declarations Emit | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/transformer/isolated-declarations"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/transformer/isolated-declarations.md for this page in Markdown format

# Isolated Declarations Emit [​](https://oxc.rs/docs/guide/usage/transformer/isolated-declarations.html#isolated-declarations-emit)

Oxc transformer supports emitting TypeScript declarations without using the TypeScript compiler for projects with [isolated declarations](https://devblogs.microsoft.com/typescript/announcing-typescript-5-5-beta/#isolated-declarations) enabled.

## Example [​](https://oxc.rs/docs/guide/usage/transformer/isolated-declarations.html#example)

**Input**:

ts

```
export function foo(a: number, b: string): number {
  return a + Number(b);
}

export enum Bar {
  a,
  b,
}
```

**Output**:

ts

```
export declare function foo(a: number, b: string): number;
export declare enum Bar {
  a = 0,
  b = 1,
}
```

## Usage [​](https://oxc.rs/docs/guide/usage/transformer/isolated-declarations.html#usage)

ts

```
import { isolatedDeclaration } from "oxc-transform";

const result = await isolatedDeclaration("lib.ts", sourceCode, {
  sourcemap: false,
  stripInternal: false,
});
```
