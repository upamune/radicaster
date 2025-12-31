---
title: "typescript/consistent-type-imports | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/typescript/consistent-type-imports.md for this page in Markdown format

# typescript/consistent-type-imports Style [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#typescript-consistent-type-imports)

🛠️ An auto-fix is available for this rule for some violations.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#what-it-does)

Enforce consistent usage of type imports.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#why-is-this-bad)

Inconsistent usage of type imports can make the code harder to read and understand.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#examples)

Examples of **incorrect** code for this rule:

ts

```
import { Foo } from "Foo";
type T = Foo;

type S = import("Foo");
```

Examples of **correct** code for this rule:

ts

```
import type { Foo } from "Foo";
```

#### Examples with `"prefer": "type-imports"` (default) [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#examples-with-prefer-type-imports-default)

Examples of **incorrect** code:

ts

```
import { Foo } from "foo";
let foo: Foo;
```

Examples of **correct** code:

ts

```
import type { Foo } from "foo";
let foo: Foo;
```

#### Examples with `"prefer": "no-type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#examples-with-prefer-no-type-imports)

Examples of **incorrect** code:

ts

```
import type { Foo } from "foo";
let foo: Foo;
```

Examples of **correct** code:

ts

```
import { Foo } from "foo";
let foo: Foo;
```

#### Examples with `"fixStyle": "inline-type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#examples-with-fixstyle-inline-type-imports)

When fixing type imports, this option will use inline `type` modifiers:

ts

```
// Before fixing
import { A, B } from "foo";
type T = A;
const b = B;

// After fixing
import { type A, B } from "foo";
type T = A;
const b = B;
```

#### Examples with `"disallowTypeAnnotations": false` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#examples-with-disallowtypeannotations-false)

When set to `false`, allows `import()` type annotations:

ts

```
type T = import("foo").Bar;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#configuration)

This rule accepts a configuration object with the following properties:

### disallowTypeAnnotations [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#disallowtypeannotations)

type: `boolean`

default: `true`

Disallow using `import()` in type annotations, like `type T = import('foo')`

### fixStyle [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#fixstyle)

type: `"separate-type-imports" | "inline-type-imports"`

default: `"separate-type-imports"`

Control how type imports are added when auto-fixing.

#### `"separate-type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#separate-type-imports)

Will add the type keyword after the import keyword `import type { A } from '...'`

#### `"inline-type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#inline-type-imports)

Will inline the type keyword `import { type A } from '...'` (only available in TypeScript 4.5+)

### prefer [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#prefer)

type: `"type-imports" | "no-type-imports"`

default: `"type-imports"`

Control whether to enforce type imports or value imports.

#### `"type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#type-imports)

Will enforce that you always use `import type Foo from '...'` except referenced by metadata of decorators.

#### `"no-type-imports"` [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#no-type-imports)

Will enforce that you always use `import Foo from '...'`

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "typescript/consistent-type-imports": "error"
  }
}
```

bash

```
oxlint --deny typescript/consistent-type-imports
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/typescript/consistent_type_imports.rs)
