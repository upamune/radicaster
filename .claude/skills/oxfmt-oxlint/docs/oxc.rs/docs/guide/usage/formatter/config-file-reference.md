---
title: "Configuration options for the Oxfmt. | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/formatter/config-file-reference"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/formatter/config-file-reference.md for this page in Markdown format

# Configuration options for the Oxfmt. [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#configuration-options-for-the-oxfmt)

Most options are the same as Prettier's options. See also <https://prettier.io/docs/options>

In addition, some options are our own extensions.

## arrowParens [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#arrowparens)

type: `string | null`

Include parentheses around a sole arrow function parameter. (Default: `"always"`)

## bracketSameLine [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#bracketsameline)

type: `boolean | null`

Put the `>` of a multi-line JSX element at the end of the last line instead of being alone on the next line. (Default: `false`)

## bracketSpacing [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#bracketspacing)

type: `boolean | null`

Print spaces between brackets in object literals. (Default: `true`)

## embeddedLanguageFormatting [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#embeddedlanguageformatting)

type: `string | null`

Control whether to format embedded parts in the file. (Default: `"off"`)

## endOfLine [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#endofline)

type: `string | null`

Which end of line characters to apply. (Default: `"lf"`)

## experimentalSortImports [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports)

type: `object | null`

Experimental: Sort import statements. Disabled by default.

### experimentalSortImports.groups [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-groups)

type: `array | null`

Custom groups configuration for organizing imports. Each array element represents a group, and multiple group names in the same array are treated as one. Accepts both `string` and `string[]` as group elements.

#### experimentalSortImports.groups[n] [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-groups-n)

type: `string[]`

### experimentalSortImports.ignoreCase [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-ignorecase)

type: `boolean | null`

Ignore case when sorting. (Default: `true`)

### experimentalSortImports.internalPattern [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-internalpattern)

type: `string[]`

Glob patterns to identify internal imports.

### experimentalSortImports.newlinesBetween [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-newlinesbetween)

type: `boolean | null`

Add newlines between import groups. (Default: `true`)

### experimentalSortImports.order [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-order)

type: `string | null`

Sort order. (Default: `"asc"`)

### experimentalSortImports.partitionByComment [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-partitionbycomment)

type: `boolean | null`

Partition imports by comments. (Default: `false`)

### experimentalSortImports.partitionByNewline [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-partitionbynewline)

type: `boolean | null`

Partition imports by newlines. (Default: `false`)

### experimentalSortImports.sortSideEffects [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortimports-sortsideeffects)

type: `boolean | null`

Sort side-effect imports. (Default: `false`)

## experimentalSortPackageJson [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#experimentalsortpackagejson)

type: `boolean | null`

Experimental: Sort `package.json` keys. (Default: `true`)

## ignorePatterns [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#ignorepatterns)

type: `string[]`

Ignore files matching these glob patterns. Current working directory is used as the root.

## insertFinalNewline [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#insertfinalnewline)

type: `boolean | null`

Whether to insert a final newline at the end of the file. (Default: `true`)

## jsxSingleQuote [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#jsxsinglequote)

type: `boolean | null`

Use single quotes instead of double quotes in JSX. (Default: `false`)

## objectWrap [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#objectwrap)

type: `string | null`

How to wrap object literals when they could fit on one line or span multiple lines. (Default: `"preserve"`) NOTE: In addition to Prettier's `"preserve"` and `"collapse"`, we also support `"always"`.

## printWidth [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#printwidth)

type: `integer | null`

The line length that the printer will wrap on. (Default: `100`)

## quoteProps [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#quoteprops)

type: `string | null`

Change when properties in objects are quoted. (Default: `"as-needed"`)

## semi [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#semi)

type: `boolean | null`

Print semicolons at the ends of statements. (Default: `true`)

## singleAttributePerLine [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#singleattributeperline)

type: `boolean | null`

Put each attribute on a new line in JSX. (Default: `false`)

## singleQuote [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#singlequote)

type: `boolean | null`

Use single quotes instead of double quotes. (Default: `false`)

## tabWidth [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#tabwidth)

type: `integer | null`

Number of spaces per indentation level. (Default: `2`)

## trailingComma [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#trailingcomma)

type: `string | null`

Print trailing commas wherever possible. (Default: `"all"`)

## useTabs [​](https://oxc.rs/docs/guide/usage/formatter/config-file-reference.html#usetabs)

type: `boolean | null`

Use tabs for indentation or spaces. (Default: `false`)
