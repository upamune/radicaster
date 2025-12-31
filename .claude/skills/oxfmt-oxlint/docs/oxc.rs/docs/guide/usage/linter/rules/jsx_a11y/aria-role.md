---
title: "jsx_a11y/aria-role | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules/jsx\_a11y/aria-role.md for this page in Markdown format

# jsx\_a11y/aria-role Correctness [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#jsx-a11y-aria-role)

### What it does [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#what-it-does)

Elements with ARIA roles must use a valid, non-abstract ARIA role. A reference to role definitions can be found at [WAI-ARIA](https://www.w3.org/TR/wai-aria/#role_definitions) site.

### Why is this bad? [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#why-is-this-bad)

The intent of this Success Criterion is to ensure that Assistive Technologies (AT) can gather information about, activate (or set) and keep up to date on the status of user interface controls in the content(such as screen readers, screen magnifiers, and speech recognition software, used by people with disabilities).

When standard controls from accessible technologies are used, this process is straightforward. If the user interface elements are used according to specification the conditions of this provision will be met.

If custom controls are created, however, or interface elements are programmed (in code or script) to have a different role and/or function than usual, then additional measures need to be taken to ensure that the controls provide important information to assistive technologies and allow themselves to be controlled by assistive technologies. A particularly important state of a user interface control is whether or not it has focus. The focus state of a control can be programmatically determined, and notifications about change of focus are sent to user agents and assistive technology. Other examples of user interface control state are whether or not a checkbox or radio button has been selected, or whether or not a collapsible tree or list node is expanded or collapsed.

### Examples [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#examples)

Examples of **incorrect** code for this rule:

jsx

```
<div role="datepicker"></div> <!-- Bad: "datepicker" is not an ARIA role -->
<div role="range"></div>      <!-- Bad: "range" is an _abstract_ ARIA role -->
<div role=""></div>           <!-- Bad: An empty ARIA role is not allowed -->
<Foo role={role}></Foo>       <!-- Bad: ignoreNonDOM is set to false or not set -->
```

Examples of **correct** code for this rule:

jsx

```
<div role="button"></div>     <!-- Good: "button" is a valid ARIA role -->
<div role={role}></div>       <!-- Good: role is a variable & cannot be determined until runtime. -->
<div></div>                   <!-- Good: No ARIA role -->
<Foo role={role}></Foo>       <!-- Good: ignoreNonDOM is set to true -->
```

## Configuration [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#configuration)

This rule accepts a configuration object with the following properties:

### allowedInvalidRoles [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#allowedinvalidroles)

type: `string[]`

default: `[]`

Custom roles that should be allowed in addition to the ARIA spec.

### ignoreNonDOM [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#ignorenondom)

type: `boolean`

default: `false`

Determines if developer-created components are checked.

## How to use [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#how-to-use)

To **enable** this rule using the config file or in the CLI, you can use:

Config (.oxlintrc.json)CLI

json

```
{
  "plugins": ["jsx-a11y"],
  "rules": {
    "jsx-a11y/aria-role": "error"
  }
}
```

bash

```
oxlint --deny jsx-a11y/aria-role --jsx-a11y-plugin
```

## References [​](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html#references)

* [Rule Source](https://github.com/oxc-project/oxc/blob/1bf0ffc0f6859c90409a9701e62e8957ef1286cc/crates/oxc_linter/src/rules/jsx_a11y/aria_role.rs)
