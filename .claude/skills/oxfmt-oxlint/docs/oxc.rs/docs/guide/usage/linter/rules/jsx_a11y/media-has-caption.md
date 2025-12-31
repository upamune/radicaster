---
title: "jsx_a11y/media-has-caption | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/media-has-caption.md for this page in Markdown format

# jsx\_a11y/media-has-caption Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#jsx-a11y-media-has-caption)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#what-it-does)

Checks if `<audio>` and `<video>` elements have a `<track>` element for captions. This ensures media content is accessible to all users, including those with hearing impairments.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#why-is-this-bad)

Without captions, audio and video content is not accessible to users who are deaf or hard of hearing. Captions are also useful for users in noisy environments or where audio is not available.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<audio></audio>
<video></video>
```

Examples of **correct** code for this rule:

jsx

```
<audio><track kind="captions" src="caption_file.vtt" /></audio>
<video><track kind="captions" src="caption_file.vtt" /></video>
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#configuration)

This rule accepts a configuration object with the following properties:

### audio [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#audio)

type: `string[]`

default: `["audio"]`

Element names to treat as `<audio>` elements

### track [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#track)

type: `string[]`

default: `["track"]`

Element names to treat as `<track>` elements

### video [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#video)

type: `string[]`

default: `["video"]`

Element names to treat as `<video>` elements

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/media-has-caption": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/media-has-caption --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/media_has_caption.rs)
