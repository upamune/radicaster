---
title: "Formatter (oxfmt) | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/formatter"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/formatter.md for this page in Markdown format

# Formatter (oxfmt) [​](https://oxc.rs/docs/guide/usage/formatter.html#formatter-oxfmt)

Oxfmt (`/oh-eks-for-mat/`) is a Prettier-compatible code formatter.

INFO

Oxfmt is currently in alpha.

Please join the discussion!

> RFC: Formatter · oxc-project/oxc · Discussion #13608
> <https://github.com/oxc-project/oxc/discussions/13608>

Waiting on Oxfmt to implement additional features? Consider using [@prettier/plugin-oxc](https://github.com/prettier/prettier/tree/main/packages/plugin-oxc) in Prettier to gain some parsing speed in the meantime.

## Features [​](https://oxc.rs/docs/guide/usage/formatter.html#features)

* Support many kinds of file types
  + JS/TS(X): Supported by `oxc_formatter`
  + All file types supported by Prettier by default
* Faster alternative of Prettier CLI
  + Over [30×](https://oxc.rs/blog/2025-12-01-oxfmt-alpha#performance) faster than Prettier’s experimental CLI without cache
* Experimental but usable features
  + Native sort-imports
  + Native sort-packagejson

## Supported languages [​](https://oxc.rs/docs/guide/usage/formatter.html#supported-languages)

* JS, JSX
* TS, TSX
* TOML
* JSON, JSONC, JSON5
* YAML
* HTML, Angular, Vue, MJML
* Ember, Handlebars
* CSS, SCSS, Less
* GraphQL
* Markdown, MDX

WARNING

Note that the `embeddedLanguageFormatting` option is not fully supported in JS/TS files. And for now, it is disabled by default.

## Installation [​](https://oxc.rs/docs/guide/usage/formatter.html#installation)

Run `oxfmt` directly at the root of your repository:

npmpnpmyarnbundeno

sh

```
$ npx oxfmt@latest
```

sh

```
$ pnpm dlx oxfmt@latest
```

sh

```
$ yarn dlx oxfmt@latest
```

sh

```
$ bunx oxfmt@latest
```

sh

```
$ deno run npm:oxfmt@latest
```

Or save it to your `package.json`:

npmpnpmyarnbun

sh

```
$ npm add -D oxfmt
```

sh

```
$ pnpm add -D oxfmt
```

sh

```
$ yarn add -D oxfmt
```

sh

```
$ bun add -D oxfmt
```

## Command-line Interface [​](https://oxc.rs/docs/guide/usage/formatter.html#command-line-interface)

`oxfmt` CLI works like `prettier --write .` by default.

Formatting config options like `--no-semi` are not supported via CLI flags. We recommend setting these via the configuration file instead. This will ensure that the CLI and editor integrations always use the same settings.

Globs in positional paths are not expanded. (You can rely on your shell.) But `!`-prefixed exclude paths do support glob expansion.

See more details in the [CLI reference](https://oxc.rs/docs/guide/usage/formatter/cli.html).

## Node.js API [​](https://oxc.rs/docs/guide/usage/formatter.html#node-js-api)

`oxfmt` is also available via Node.js API: `format()` function.

ts

```
import { format } from "oxfmt";
import type { FormatOptions } from "oxfmt";

const INPUT = `let a=42;`;
const options: FormatOptions = {
  semi: false,
};

const { code, errors } = await format("a.js", INPUT, options);
console.log(code); // "let a = 42"
```

## System Requirements [​](https://oxc.rs/docs/guide/usage/formatter.html#system-requirements)

* **Node.js**: >= 20.19.0 or >= 22.12.0
* **Platforms**: darwin-arm64, darwin-x64, linux-arm64, linux-x64, win32-arm64, and win32-x64

## FAQs [​](https://oxc.rs/docs/guide/usage/formatter.html#faqs)

### Are there any formatting differences with Prettier? [​](https://oxc.rs/docs/guide/usage/formatter.html#are-there-any-formatting-differences-with-prettier)

For JS/TS files, we're tested against the `v3.7.3` of Prettier.

For known differences, please see this discussion.

> `Oxfmt` differences with `Prettier` · oxc-project/oxc · Discussion #14669
> <https://github.com/oxc-project/oxc/discussions/14669>

### Are there any limitations for configuration with Prettier? [​](https://oxc.rs/docs/guide/usage/formatter.html#are-there-any-limitations-for-configuration-with-prettier)

The following are NOT currently supported:

* `prettier` field in `package.json`
* Config file format other than `.json` and `.jsonc`
* `overrides` field
* Nested configs in sub directories
* Nested `.editorconfig` in sub directories
* `experimentalTernaries` and `experimentalOperatorPosition` option

Also, if `printWidth` is not specified, its default value is `100`. This differs from Prettier's default `80`.

### Are Prettier plugins supported? [​](https://oxc.rs/docs/guide/usage/formatter.html#are-prettier-plugins-supported)

Currently, NOT supported.

However, for import sorting functionality, we provide experimental behavior based on `eslint-plugin-perfectionist/sort-imports` through the `experimentalSortImports` option.

And for `prettier-plugin-packagejson`, we have the `experimentalSortPackageJson` option, which is enabled by default.

See more details in the [Configuration file reference](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html).

### Why are nested scripts and code blocks not formatted? [​](https://oxc.rs/docs/guide/usage/formatter.html#why-are-nested-scripts-and-code-blocks-not-formatted)

Currently, the `embeddedLanguageFormatting` option is `"off"` by default. Please set it to `"auto"` in your config file.

However, even with `"auto"`, the contents inside `TaggedTemplateLiteral` in JS/TS files may not be fully formatted in some cases.
