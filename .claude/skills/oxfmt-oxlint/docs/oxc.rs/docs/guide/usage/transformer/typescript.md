---
title: "TypeScript | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/transformer/typescript"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/transformer/typescript.md for this page in Markdown format

# TypeScript [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#typescript)

Oxc transformer supports transforming TypeScript to JavaScript.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.ts", sourceCode, {
  typescript: {
    jsxPragma: "React.createElement",
    jsxPragmaFrag: "React.Fragment",
    onlyRemoveTypeImports: false,
    allowNamespaces: true,
    removeClassFieldsWithoutInitializer: false,
    rewriteImportExtensions: false,
  },
});
```

## `verbatimModuleSyntax` [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#verbatimmodulesyntax)

By default, TypeScript removes unused imports in a different semantics than the JavaScript specification. The [`verbatimModuleSyntax`](https://www.typescriptlang.org/tsconfig/#verbatimModuleSyntax) option tells TypeScript to align with the JavaScript specification.

If you are using this option, make sure to set `typescript.onlyRemoveTypeImports` option to `true`.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.ts", sourceCode, {
  typescript: {
    onlyRemoveTypeImports: true,
  },
});
```

## `useDefineForClassFields` [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#usedefineforclassfields)

TypeScript used to have a different semantics for class fields than the JavaScript specification. The [`useDefineForClassFields`](https://www.typescriptlang.org/tsconfig/#useDefineForClassFields) option tells TypeScript to align with the JavaScript specification. This options is enabled by default if the `target` option in the tsconfig is set to `es2022` or higher.

If you are disabling this option, make sure to set `typescript.removeClassFieldsWithoutInitializer` option and `assumptions.setPublicClassFields` to `true`.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.ts", sourceCode, {
  typescript: {
    removeClassFieldsWithoutInitializer: true,
  },
  assumptions: {
    setPublicClassFields: true,
  },
});
```

## Decorators [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#decorators)

Oxc transformer supports transforming legacy decorators. This is called experimental decorators in TypeScript.

If you are using the [`experimentalDecorators`](https://www.typescriptlang.org/tsconfig/#experimentalDecorators) option in the tsconfig, you can use the `decorators.legacy` option. If you are using the [`emitDecoratorMetadata`](https://www.typescriptlang.org/tsconfig/#emitDecoratorMetadata) option in the tsconfig, you can use the `decorators.emitDecoratorMetadata` option.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.ts", sourceCode, {
  decorators: {
    legacy: true,
    emitDecoratorMetadata: true,
  },
});
```

Decorator Metadata: Types that requires type inference will fallback to `Object`

Due to the lack of full type inference feature, Oxc transformer will fallback to `Object` type if it cannot calculate the type of the decorator metadata.

For example, the following code will be transformed to the following code:

input.tsoutput(oxc).jsoutput(typescript\_compiler).js

ts

```
import { Something1 } from "./somewhere";

type Something2 = Exclude<string | number, string>;

export class Foo {
  @test
  foo(input1: Something1, input2: Something2) {}
}
```

js

```
// omit helper functions
import { Something1 } from "./somewhere";
var _ref;
export class Foo {
  foo(input1, input2) {}
}
_decorate(
  [
    test,
    _decorateMetadata("design:type", Function),
    _decorateMetadata("design:paramtypes", [
      typeof (_ref = typeof Something1 !== "undefined" && Something1) === "function"
        ? _ref
        : Object,
      Object,
    ]),
    _decorateMetadata("design:returntype", void 0),
  ],
  Foo.prototype,
  "foo",
  null,
);
```

js

```
// omit helper functions
var _a;
import { Something1 } from "./somewhere";
export class Foo {
  foo(input1, input2) {}
}
__decorate(
  [
    test,
    __metadata("design:type", Function),
    __metadata("design:paramtypes", [
      typeof (_a = typeof Something1 !== "undefined" && Something1) === "function" ? _a : Object,
      Number,
    ]),
    __metadata("design:returntype", void 0),
  ],
  Foo.prototype,
  "foo",
  null,
);
```

This behavior aligns with TypeScript's behavior when using a type that is external.

## TSX [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#tsx)

Transforming TSX files is supported as well. See [JSX transform](https://oxc.rs/docs/guide/usage/transformer/jsx.html) for more information.

## Rewriting import extensions [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#rewriting-import-extensions)

If you are using the [`rewriteImportExtensions`](https://www.typescriptlang.org/tsconfig/#rewriteImportExtensions) option in the tsconfig, you can use the `typescript.rewriteImportExtensions` option.

js

```
import { transform } from "oxc-transform";

const result = await transform("lib.ts", sourceCode, {
  typescript: {
    rewriteImportExtensions: "rewrite", // or "remove"
  },
});
```

## Caveats [​](https://oxc.rs/docs/guide/usage/transformer/typescript.html#caveats)

Because Oxc transformer transforms each files independently, some TypeScript features are not supported. To avoid using unsupported features, you should enable the [`isolatedModules`](https://www.typescriptlang.org/tsconfig/#isolatedModules) option in your `tsconfig.json` file.
