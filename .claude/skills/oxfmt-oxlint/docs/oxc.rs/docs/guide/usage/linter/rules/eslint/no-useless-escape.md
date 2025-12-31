---
title: "eslint/no-useless-escape | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-useless-escape.md for this page in Markdown format

# eslint/no-useless-escape Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#eslint-no-useless-escape)

✅ This rule is turned on by default.

🛠️ An auto-fix is available for this rule.

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#what-it-does)

Disallow unnecessary escape characters.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#why-is-this-bad)

Escaping characters unnecessarily has no effect on the behavior of strings or regexes, and can make code harder to read and understand by adding unnecessary complexity. This applies to string literals, template literals, and regular expressions.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/*eslint no-useless-escape: "error"*/

"\'";
'\"';
"\#";
"\e";
`\"`;
`\"${foo}\"`;
`\#{foo}`;
/\!/;
/\@/;
/[\[]/;
/[a-z\-]/;
```

Examples of **correct** code for this rule:

javascript

```
/*eslint no-useless-escape: "error"*/

"\"";
'\'';
"\x12";
"\u00a9";
"\371";
"xs\u2111";
`\``;
`\${${foo}}`;
`$\{${foo}}`;
/\\/g;
/\t/g;
/\w\$\*\^\./;
/[[]/;
/[\]]/;
/[a-z-]/;
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#configuration)

This rule accepts a configuration object with the following properties:

### allowRegexCharacters [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#allowregexcharacters)

type: `string[]`

default: `[]`

An array of characters that are allowed to be escaped unnecessarily in regexes.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-useless-escape": "error"
  }
}
```

bash

```
oxlint --deny no-useless-escape
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_useless_escape.rs)
