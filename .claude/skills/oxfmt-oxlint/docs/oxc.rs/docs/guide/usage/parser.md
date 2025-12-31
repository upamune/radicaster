---
title: "Parser | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/parser"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/parser.md for this page in Markdown format

# Parser [​](https://oxc.rs/docs/guide/usage/parser.html#parser)

This is production ready.

## Features [​](https://oxc.rs/docs/guide/usage/parser.html#features)

* 3x faster than swc parser ([benchmark](https://github.com/oxc-project/bench-javascript-parser-written-in-rust)).
* Parses `.js(x)` and `.ts(x)`.
* Passes all parser tests from Test262 and 99% from Babel and TypeScript.
* Returns ESM information directly, no need for [`es-module-lexer`](https://github.com/guybedford/es-module-lexer).
* [✅ works with checker.ts](https://x.com/robpalmer2/status/1805502964435505559)

## Installation [​](https://oxc.rs/docs/guide/usage/parser.html#installation)

### Node.js [​](https://oxc.rs/docs/guide/usage/parser.html#node-js)

* Use the node binding [oxc-parser](https://www.npmjs.com/package/oxc-parser).
* Try on [stackblitz](https://stackblitz.com/edit/oxc-parser).

### Rust [​](https://oxc.rs/docs/guide/usage/parser.html#rust)

Use the umbrella crate [oxc](https://docs.rs/oxc) or the individual [oxc\_ast](https://docs.rs/oxc_ast) and [oxc\_parser](https://docs.rs/oxc_parser) crates.

Rust usage example can be found [here](https://github.com/oxc-project/oxc/blob/main/crates/oxc_parser/examples/parser.rs).

## Print [​](https://oxc.rs/docs/guide/usage/parser.html#print)

After parsing and transforming, you can print code.

Here's a direct example using [esrap](https://www.npmjs.com/package/esrap) *(`parse` in reverse!)*:

js

```
import { print } from "esrap";
import ts from "esrap/languages/ts";
import { parseSync } from "oxc-parser";

const { program } = parseSync("test.js", 'alert("hello oxc & esrap");');
const { code } = print(program, ts());

console.log(code); // alert("hello oxc & esrap");
```

INFO

Today, comments are not printed. *It will be supported thanks to [oxc-parser #13285](https://github.com/oxc-project/oxc/pull/13285).*
