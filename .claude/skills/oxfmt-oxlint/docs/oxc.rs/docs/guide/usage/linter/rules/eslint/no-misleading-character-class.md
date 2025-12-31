---
title: "eslint/no-misleading-character-class | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/eslint/no-misleading-character-class.md for this page in Markdown format

# eslint/no-misleading-character-class Nursery [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#eslint-no-misleading-character-class)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#what-it-does)

This rule reports regular expressions which include multiple code point characters in character class syntax. This includes:

* Characters with combining marks (e.g., `Á` where `A` is followed by a combining acute accent)
* Characters with emoji modifiers (e.g., `👶🏻`)
* Pairs of regional indicator symbols (e.g., `🇯🇵`)
* Characters joined by zero-width joiner (ZWJ) (e.g., `👨‍👩‍👦`)
* Surrogate pairs without the Unicode flag (e.g., `/^[👍]$/`)

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#why-is-this-bad)

Unicode includes characters which are made by multiple code points. RegExp character class syntax (`/[abc]/`) cannot handle characters which are made by multiple code points as a character; those characters will be dissolved to each code point. For example, `❇️` is made by `❇` (`U+2747`) and VARIATION SELECTOR-16 (`U+FE0F`). If this character is in a RegExp character class, it will match either `❇` (`U+2747`) or VARIATION SELECTOR-16 (`U+FE0F`) rather than `❇️`.

This can lead to regular expressions that do not match what the author intended, especially for emoji, regional indicators, and characters with combining marks.

#### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#examples)

Examples of **incorrect** code for this rule:

javascript

```
/^[Á]$/u;
/^[❇️]$/u;
/^[👶🏻]$/u;
/^[🇯🇵]$/u;
/^[👨‍👩‍👦]$/u;
/^[👍]$/;
new RegExp("[🎵]");
```

Examples of **correct** code for this rule:

javascript

```
/^[abc]$/;
/^[👍]$/u;
/[\u00B7\u0300-\u036F]/u;
new RegExp("^[\u{1F1EF}\u{1F1F5}]", "u");
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#configuration)

This rule accepts a configuration object with the following properties:

### allowEscape [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#allowescape)

type: `boolean`

default: `false`

When set to `true`, the rule allows any grouping of code points inside a character class as long as they are written using escape sequences.

Examples of **incorrect** code for this rule with `{ "allowEscape": true }`:

javascript

```
/[\uD83D]/; // backslash can be omitted
new RegExp("[\ud83d" + "\udc4d]");
```

Examples of **correct** code for this rule with `{ "allowEscape": true }`:

javascript

```
/[\ud83d\udc4d]/;
/[\u00B7\u0300-\u036F]/u;
/[👨\u200d👩]/u;
new RegExp("[\x41\u0301]");
new RegExp(`[\u{1F1EF}\u{1F1F5}]`, "u");
new RegExp("[\\u{1F1EF}\\u{1F1F5}]", "u");
```

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "rules": {
    "no-misleading-character-class": "error"
  }
}
```

bash

```
oxlint --deny no-misleading-character-class
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/eslint/no_misleading_character_class.rs)
