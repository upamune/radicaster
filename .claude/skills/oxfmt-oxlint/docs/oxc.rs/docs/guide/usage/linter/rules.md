---
title: "Rules | The JavaScript Oxidation Compiler"
source_url: "https://oxc.rs/docs/guide/usage/linter/rules"
fetched_at: "2025-12-31T10:44:14.234719+00:00"
---



Are you an LLM? You can read better optimized documentation at /docs/guide/usage/linter/rules.md for this page in Markdown format

# Rules [​](https://oxc.rs/docs/guide/usage/linter/rules.html#rules)

The progress of all rule implementations is tracked [here](https://github.com/oxc-project/oxc/issues/481).

* Total number of rules: 645
* Rules turned on by default: 104

**Legend for 'Fixable?' column:**

* 🛠️: an auto-fix is available for this rule
* 💡: a suggestion is available for this rule
* ⚠️🛠️: a dangerous auto-fix is available for this rule
* ⚠️💡: a dangerous suggestion is available for this rule
* 🚧: an auto-fix or suggestion is possible, but currently not implemented

## Correctness (201): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#correctness-201)

Code that is outright wrong or useless.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [constructor-super](https://oxc.rs/docs/guide/usage/linter/rules/eslint/constructor-super.html) | eslint | ✅ |  |
| [for-direction](https://oxc.rs/docs/guide/usage/linter/rules/eslint/for-direction.html) | eslint | ✅ | ⚠️🛠️️ |
| [no-async-promise-executor](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-async-promise-executor.html) | eslint | ✅ |  |
| [no-caller](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-caller.html) | eslint | ✅ |  |
| [no-class-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-class-assign.html) | eslint | ✅ |  |
| [no-compare-neg-zero](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-compare-neg-zero.html) | eslint | ✅ | 🛠️💡 |
| [no-cond-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-cond-assign.html) | eslint | ✅ |  |
| [no-const-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-const-assign.html) | eslint | ✅ |  |
| [no-constant-binary-expression](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-binary-expression.html) | eslint | ✅ |  |
| [no-constant-condition](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constant-condition.html) | eslint | ✅ |  |
| [no-control-regex](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-control-regex.html) | eslint | ✅ |  |
| [no-debugger](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-debugger.html) | eslint | ✅ | 🛠️ |
| [no-delete-var](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-delete-var.html) | eslint | ✅ |  |
| [no-dupe-class-members](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-class-members.html) | eslint | ✅ |  |
| [no-dupe-else-if](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-else-if.html) | eslint | ✅ |  |
| [no-dupe-keys](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-dupe-keys.html) | eslint | ✅ |  |
| [no-duplicate-case](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-case.html) | eslint | ✅ |  |
| [no-empty-character-class](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-character-class.html) | eslint | ✅ |  |
| [no-empty-pattern](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-pattern.html) | eslint | ✅ |  |
| [no-empty-static-block](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-static-block.html) | eslint | ✅ | 💡 |
| [no-eval](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eval.html) | eslint | ✅ |  |
| [no-ex-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ex-assign.html) | eslint | ✅ |  |
| [no-extra-boolean-cast](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-boolean-cast.html) | eslint | ✅ | 🛠️💡 |
| [no-func-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-func-assign.html) | eslint | ✅ |  |
| [no-global-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-global-assign.html) | eslint | ✅ |  |
| [no-import-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-import-assign.html) | eslint | ✅ |  |
| [no-invalid-regexp](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-invalid-regexp.html) | eslint | ✅ |  |
| [no-irregular-whitespace](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-irregular-whitespace.html) | eslint | ✅ |  |
| [no-loss-of-precision](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loss-of-precision.html) | eslint | ✅ |  |
| [no-new-native-nonconstructor](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-native-nonconstructor.html) | eslint | ✅ |  |
| [no-nonoctal-decimal-escape](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nonoctal-decimal-escape.html) | eslint | ✅ | 💡 |
| [no-obj-calls](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-obj-calls.html) | eslint | ✅ |  |
| [no-self-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-assign.html) | eslint | ✅ |  |
| [no-setter-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-setter-return.html) | eslint | ✅ |  |
| [no-shadow-restricted-names](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-shadow-restricted-names.html) | eslint | ✅ |  |
| [no-sparse-arrays](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sparse-arrays.html) | eslint | ✅ |  |
| [no-this-before-super](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-this-before-super.html) | eslint | ✅ |  |
| [no-unassigned-vars](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unassigned-vars.html) | eslint | ✅ |  |
| [no-unsafe-finally](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-finally.html) | eslint | ✅ |  |
| [no-unsafe-negation](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-negation.html) | eslint | ✅ | 🛠️ |
| [no-unsafe-optional-chaining](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unsafe-optional-chaining.html) | eslint | ✅ |  |
| [no-unused-expressions](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-expressions.html) | eslint | ✅ |  |
| [no-unused-labels](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-labels.html) | eslint | ✅ | 🛠️ |
| [no-unused-private-class-members](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-private-class-members.html) | eslint | ✅ |  |
| [no-unused-vars](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unused-vars.html) | eslint | ✅ | ⚠️💡 |
| [no-useless-backreference](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-backreference.html) | eslint | ✅ |  |
| [no-useless-catch](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-catch.html) | eslint | ✅ |  |
| [no-useless-escape](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-escape.html) | eslint | ✅ | 🛠️ |
| [no-useless-rename](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-rename.html) | eslint | ✅ |  |
| [no-with](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-with.html) | eslint | ✅ |  |
| [require-yield](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-yield.html) | eslint | ✅ |  |
| [use-isnan](https://oxc.rs/docs/guide/usage/linter/rules/eslint/use-isnan.html) | eslint | ✅ | 🛠️ |
| [valid-typeof](https://oxc.rs/docs/guide/usage/linter/rules/eslint/valid-typeof.html) | eslint | ✅ | 🛠️ |
| [default](https://oxc.rs/docs/guide/usage/linter/rules/import/default.html) | import |  |  |
| [namespace](https://oxc.rs/docs/guide/usage/linter/rules/import/namespace.html) | import |  |  |
| [expect-expect](https://oxc.rs/docs/guide/usage/linter/rules/jest/expect-expect.html) | jest |  |  |
| [no-conditional-expect](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-expect.html) | jest |  |  |
| [no-disabled-tests](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-disabled-tests.html) | jest |  |  |
| [no-export](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-export.html) | jest |  |  |
| [no-focused-tests](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-focused-tests.html) | jest |  | 🛠️ |
| [no-standalone-expect](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-standalone-expect.html) | jest |  |  |
| [require-to-throw-message](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-to-throw-message.html) | jest |  |  |
| [valid-describe-callback](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-describe-callback.html) | jest |  |  |
| [valid-expect](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-expect.html) | jest |  |  |
| [valid-title](https://oxc.rs/docs/guide/usage/linter/rules/jest/valid-title.html) | jest |  | 🛠️ |
| [check-property-names](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-property-names.html) | jsdoc |  |  |
| [check-tag-names](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-tag-names.html) | jsdoc |  |  |
| [implements-on-classes](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/implements-on-classes.html) | jsdoc |  |  |
| [no-defaults](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/no-defaults.html) | jsdoc |  |  |
| [require-property](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property.html) | jsdoc |  |  |
| [require-property-description](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-description.html) | jsdoc |  |  |
| [require-property-name](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-name.html) | jsdoc |  |  |
| [require-property-type](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-property-type.html) | jsdoc |  |  |
| [require-yields](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-yields.html) | jsdoc |  |  |
| [alt-text](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/alt-text.html) | jsx\_a11y |  |  |
| [anchor-has-content](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-has-content.html) | jsx\_a11y |  | 💡 |
| [anchor-is-valid](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-is-valid.html) | jsx\_a11y |  |  |
| [aria-activedescendant-has-tabindex](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-activedescendant-has-tabindex.html) | jsx\_a11y |  |  |
| [aria-props](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-props.html) | jsx\_a11y |  | 🛠️ |
| [aria-proptypes](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-proptypes.html) | jsx\_a11y |  |  |
| [aria-role](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-role.html) | jsx\_a11y |  |  |
| [aria-unsupported-elements](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/aria-unsupported-elements.html) | jsx\_a11y |  | 🛠️ |
| [autocomplete-valid](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/autocomplete-valid.html) | jsx\_a11y |  |  |
| [click-events-have-key-events](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/click-events-have-key-events.html) | jsx\_a11y |  |  |
| [heading-has-content](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/heading-has-content.html) | jsx\_a11y |  |  |
| [html-has-lang](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/html-has-lang.html) | jsx\_a11y |  |  |
| [iframe-has-title](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/iframe-has-title.html) | jsx\_a11y |  |  |
| [img-redundant-alt](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/img-redundant-alt.html) | jsx\_a11y |  |  |
| [label-has-associated-control](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/label-has-associated-control.html) | jsx\_a11y |  |  |
| [lang](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/lang.html) | jsx\_a11y |  |  |
| [media-has-caption](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/media-has-caption.html) | jsx\_a11y |  |  |
| [mouse-events-have-key-events](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/mouse-events-have-key-events.html) | jsx\_a11y |  |  |
| [no-access-key](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-access-key.html) | jsx\_a11y |  | 💡 |
| [no-aria-hidden-on-focusable](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-aria-hidden-on-focusable.html) | jsx\_a11y |  | 🛠️ |
| [no-autofocus](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-autofocus.html) | jsx\_a11y |  | 🛠️ |
| [no-distracting-elements](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-distracting-elements.html) | jsx\_a11y |  |  |
| [no-noninteractive-tabindex](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-noninteractive-tabindex.html) | jsx\_a11y |  |  |
| [no-redundant-roles](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/no-redundant-roles.html) | jsx\_a11y |  | 🛠️ |
| [prefer-tag-over-role](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/prefer-tag-over-role.html) | jsx\_a11y |  |  |
| [role-has-required-aria-props](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-has-required-aria-props.html) | jsx\_a11y |  |  |
| [role-supports-aria-props](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/role-supports-aria-props.html) | jsx\_a11y |  |  |
| [scope](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/scope.html) | jsx\_a11y |  | 🛠️ |
| [tabindex-no-positive](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/tabindex-no-positive.html) | jsx\_a11y |  | ⚠️💡 |
| [google-font-display](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-display.html) | nextjs |  |  |
| [google-font-preconnect](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/google-font-preconnect.html) | nextjs |  |  |
| [inline-script-id](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/inline-script-id.html) | nextjs |  |  |
| [next-script-for-ga](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/next-script-for-ga.html) | nextjs |  |  |
| [no-assign-module-variable](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-assign-module-variable.html) | nextjs |  |  |
| [no-async-client-component](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-async-client-component.html) | nextjs |  |  |
| [no-before-interactive-script-outside-document](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-before-interactive-script-outside-document.html) | nextjs |  |  |
| [no-css-tags](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-css-tags.html) | nextjs |  |  |
| [no-document-import-in-page](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-document-import-in-page.html) | nextjs |  |  |
| [no-duplicate-head](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-duplicate-head.html) | nextjs |  |  |
| [no-head-element](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-element.html) | nextjs |  |  |
| [no-head-import-in-document](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-head-import-in-document.html) | nextjs |  |  |
| [no-html-link-for-pages](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-html-link-for-pages.html) | nextjs |  |  |
| [no-img-element](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-img-element.html) | nextjs |  | 🚧 |
| [no-page-custom-font](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-page-custom-font.html) | nextjs |  |  |
| [no-script-component-in-head](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-script-component-in-head.html) | nextjs |  |  |
| [no-styled-jsx-in-document](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-styled-jsx-in-document.html) | nextjs |  |  |
| [no-sync-scripts](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-sync-scripts.html) | nextjs |  |  |
| [no-title-in-document-head](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-title-in-document-head.html) | nextjs |  |  |
| [no-typos](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-typos.html) | nextjs |  | 🚧 |
| [no-unwanted-polyfillio](https://oxc.rs/docs/guide/usage/linter/rules/nextjs/no-unwanted-polyfillio.html) | nextjs |  |  |
| [bad-array-method-on-arguments](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-array-method-on-arguments.html) | oxc | ✅ |  |
| [bad-char-at-comparison](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-char-at-comparison.html) | oxc | ✅ |  |
| [bad-comparison-sequence](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-comparison-sequence.html) | oxc | ✅ |  |
| [bad-min-max-func](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-min-max-func.html) | oxc | ✅ |  |
| [bad-object-literal-comparison](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-object-literal-comparison.html) | oxc | ✅ |  |
| [bad-replace-all-arg](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-replace-all-arg.html) | oxc | ✅ |  |
| [const-comparisons](https://oxc.rs/docs/guide/usage/linter/rules/oxc/const-comparisons.html) | oxc | ✅ |  |
| [double-comparisons](https://oxc.rs/docs/guide/usage/linter/rules/oxc/double-comparisons.html) | oxc | ✅ | 🛠️ |
| [erasing-op](https://oxc.rs/docs/guide/usage/linter/rules/oxc/erasing-op.html) | oxc | ✅ | ⚠️🛠️️ |
| [missing-throw](https://oxc.rs/docs/guide/usage/linter/rules/oxc/missing-throw.html) | oxc | ✅ | 💡 |
| [number-arg-out-of-range](https://oxc.rs/docs/guide/usage/linter/rules/oxc/number-arg-out-of-range.html) | oxc | ✅ |  |
| [only-used-in-recursion](https://oxc.rs/docs/guide/usage/linter/rules/oxc/only-used-in-recursion.html) | oxc | ✅ | ⚠️🛠️️ |
| [uninvoked-array-callback](https://oxc.rs/docs/guide/usage/linter/rules/oxc/uninvoked-array-callback.html) | oxc | ✅ |  |
| [no-callback-in-promise](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-callback-in-promise.html) | promise |  |  |
| [no-new-statics](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-new-statics.html) | promise |  | 🛠️ |
| [valid-params](https://oxc.rs/docs/guide/usage/linter/rules/promise/valid-params.html) | promise |  |  |
| [exhaustive-deps](https://oxc.rs/docs/guide/usage/linter/rules/react/exhaustive-deps.html) | react |  | ⚠️🛠️️💡 |
| [forward-ref-uses-ref](https://oxc.rs/docs/guide/usage/linter/rules/react/forward-ref-uses-ref.html) | react |  | 💡 |
| [jsx-key](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-key.html) | react |  |  |
| [jsx-no-duplicate-props](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-duplicate-props.html) | react |  |  |
| [jsx-no-undef](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-undef.html) | react |  |  |
| [jsx-props-no-spread-multi](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spread-multi.html) | react |  | 🛠️ |
| [no-children-prop](https://oxc.rs/docs/guide/usage/linter/rules/react/no-children-prop.html) | react |  |  |
| [no-danger-with-children](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger-with-children.html) | react |  |  |
| [no-did-mount-set-state](https://oxc.rs/docs/guide/usage/linter/rules/react/no-did-mount-set-state.html) | react |  |  |
| [no-direct-mutation-state](https://oxc.rs/docs/guide/usage/linter/rules/react/no-direct-mutation-state.html) | react |  |  |
| [no-find-dom-node](https://oxc.rs/docs/guide/usage/linter/rules/react/no-find-dom-node.html) | react |  |  |
| [no-is-mounted](https://oxc.rs/docs/guide/usage/linter/rules/react/no-is-mounted.html) | react |  |  |
| [no-render-return-value](https://oxc.rs/docs/guide/usage/linter/rules/react/no-render-return-value.html) | react |  |  |
| [no-string-refs](https://oxc.rs/docs/guide/usage/linter/rules/react/no-string-refs.html) | react |  |  |
| [no-unsafe](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unsafe.html) | react |  |  |
| [void-dom-elements-no-children](https://oxc.rs/docs/guide/usage/linter/rules/react/void-dom-elements-no-children.html) | react |  |  |
| [await-thenable](https://oxc.rs/docs/guide/usage/linter/rules/typescript/await-thenable.html) | typescript | ✅ | 🚧 |
| [no-array-delete](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-array-delete.html) | typescript | ✅ | 🚧 |
| [no-base-to-string](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-base-to-string.html) | typescript | ✅ | 🚧 |
| [no-duplicate-enum-values](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-enum-values.html) | typescript | ✅ |  |
| [no-duplicate-type-constituents](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-duplicate-type-constituents.html) | typescript | ✅ | 🚧 |
| [no-extra-non-null-assertion](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extra-non-null-assertion.html) | typescript | ✅ |  |
| [no-floating-promises](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-floating-promises.html) | typescript | ✅ | 🚧 |
| [no-for-in-array](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-for-in-array.html) | typescript | ✅ | 🚧 |
| [no-implied-eval](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-implied-eval.html) | typescript | ✅ | 🚧 |
| [no-meaningless-void-operator](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-meaningless-void-operator.html) | typescript | ✅ | 🚧 |
| [no-misused-new](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-new.html) | typescript | ✅ |  |
| [no-misused-spread](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-spread.html) | typescript | ✅ | 🚧 |
| [no-non-null-asserted-optional-chain](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-optional-chain.html) | typescript | ✅ | 💡 |
| [no-redundant-type-constituents](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-redundant-type-constituents.html) | typescript | ✅ | 🚧 |
| [no-this-alias](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-this-alias.html) | typescript | ✅ |  |
| [no-unnecessary-parameter-property-assignment](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-parameter-property-assignment.html) | typescript | ✅ | 💡 |
| [no-unsafe-declaration-merging](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-declaration-merging.html) | typescript | ✅ |  |
| [no-unsafe-unary-minus](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-unary-minus.html) | typescript | ✅ | 🚧 |
| [no-useless-empty-export](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-useless-empty-export.html) | typescript | ✅ | 🛠️ |
| [no-wrapper-object-types](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-wrapper-object-types.html) | typescript | ✅ | 🛠️ |
| [prefer-as-const](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-as-const.html) | typescript | ✅ | 🛠️ |
| [require-array-sort-compare](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-array-sort-compare.html) | typescript | ✅ | 🚧 |
| [restrict-template-expressions](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-template-expressions.html) | typescript | ✅ | 🚧 |
| [triple-slash-reference](https://oxc.rs/docs/guide/usage/linter/rules/typescript/triple-slash-reference.html) | typescript | ✅ |  |
| [unbound-method](https://oxc.rs/docs/guide/usage/linter/rules/typescript/unbound-method.html) | typescript | ✅ | 🚧 |
| [no-await-in-promise-methods](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-in-promise-methods.html) | unicorn | ✅ |  |
| [no-empty-file](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-empty-file.html) | unicorn | ✅ |  |
| [no-invalid-fetch-options](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-fetch-options.html) | unicorn | ✅ |  |
| [no-invalid-remove-event-listener](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-invalid-remove-event-listener.html) | unicorn | ✅ |  |
| [no-new-array](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-array.html) | unicorn | ✅ | 🚧 |
| [no-single-promise-in-promise-methods](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-single-promise-in-promise-methods.html) | unicorn | ✅ | 🛠️ |
| [no-thenable](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-thenable.html) | unicorn | ✅ |  |
| [no-unnecessary-await](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-await.html) | unicorn | ✅ | 🛠️ |
| [no-useless-fallback-in-spread](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-fallback-in-spread.html) | unicorn | ✅ | 🛠️ |
| [no-useless-length-check](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-length-check.html) | unicorn | ✅ | 🚧 |
| [no-useless-spread](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-spread.html) | unicorn | ✅ | ⚠️🛠️️ |
| [prefer-set-size](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-size.html) | unicorn | ✅ | 🛠️ |
| [prefer-string-starts-ends-with](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-starts-ends-with.html) | unicorn | ✅ | 🛠️ |
| [no-conditional-tests](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-conditional-tests.html) | vitest |  |  |
| [require-local-test-context-for-concurrent-snapshots](https://oxc.rs/docs/guide/usage/linter/rules/vitest/require-local-test-context-for-concurrent-snapshots.html) | vitest |  | 🚧 |
| [no-deprecated-destroyed-lifecycle](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-deprecated-destroyed-lifecycle.html) | vue |  | 🛠️ |
| [no-export-in-script-setup](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-export-in-script-setup.html) | vue |  |  |
| [prefer-import-from-vue](https://oxc.rs/docs/guide/usage/linter/rules/vue/prefer-import-from-vue.html) | vue |  | 🛠️ |
| [valid-define-emits](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-emits.html) | vue |  | 🚧 |
| [valid-define-props](https://oxc.rs/docs/guide/usage/linter/rules/vue/valid-define-props.html) | vue |  | 🚧 |

## Perf (12): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#perf-12)

Code that can be written to run faster.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [no-await-in-loop](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-await-in-loop.html) | eslint |  |  |
| [no-useless-call](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-call.html) | eslint |  |  |
| [no-accumulating-spread](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-accumulating-spread.html) | oxc |  |  |
| [no-map-spread](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-map-spread.html) | oxc |  | 🛠️💡 |
| [no-array-index-key](https://oxc.rs/docs/guide/usage/linter/rules/react/no-array-index-key.html) | react |  |  |
| [jsx-no-jsx-as-prop](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-jsx-as-prop.html) | react\_perf |  |  |
| [jsx-no-new-array-as-prop](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-array-as-prop.html) | react\_perf |  |  |
| [jsx-no-new-function-as-prop](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-function-as-prop.html) | react\_perf |  |  |
| [jsx-no-new-object-as-prop](https://oxc.rs/docs/guide/usage/linter/rules/react_perf/jsx-no-new-object-as-prop.html) | react\_perf |  |  |
| [prefer-array-find](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-find.html) | unicorn |  | 🚧 |
| [prefer-array-flat-map](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat-map.html) | unicorn |  | 🛠️ |
| [prefer-set-has](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-set-has.html) | unicorn |  | ⚠️🛠️️ |

## Restriction (80): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#restriction-80)

Lints which prevent the use of language and library features. Must not be enabled as a whole, should be considered on a case-by-case basis before enabling.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [class-methods-use-this](https://oxc.rs/docs/guide/usage/linter/rules/eslint/class-methods-use-this.html) | eslint |  |  |
| [default-case](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case.html) | eslint |  |  |
| [no-alert](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-alert.html) | eslint |  |  |
| [no-bitwise](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-bitwise.html) | eslint |  |  |
| [no-console](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-console.html) | eslint |  | 💡 |
| [no-div-regex](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-div-regex.html) | eslint |  | 🛠️ |
| [no-empty](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty.html) | eslint |  | 💡 |
| [no-empty-function](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-empty-function.html) | eslint |  |  |
| [no-eq-null](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-eq-null.html) | eslint |  | ⚠️🛠️️ |
| [no-iterator](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-iterator.html) | eslint |  | 💡 |
| [no-param-reassign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-param-reassign.html) | eslint |  |  |
| [no-plusplus](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-plusplus.html) | eslint |  | 💡 |
| [no-proto](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-proto.html) | eslint |  | 🚧 |
| [no-regex-spaces](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-regex-spaces.html) | eslint |  | 🛠️ |
| [no-restricted-globals](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-globals.html) | eslint |  |  |
| [no-restricted-imports](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-restricted-imports.html) | eslint |  |  |
| [no-sequences](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-sequences.html) | eslint |  |  |
| [no-undefined](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undefined.html) | eslint |  |  |
| [no-var](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-var.html) | eslint |  | 🛠️ |
| [no-void](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-void.html) | eslint |  | 💡 |
| [unicode-bom](https://oxc.rs/docs/guide/usage/linter/rules/eslint/unicode-bom.html) | eslint |  | 🛠️ |
| [extensions](https://oxc.rs/docs/guide/usage/linter/rules/import/extensions.html) | import |  |  |
| [no-amd](https://oxc.rs/docs/guide/usage/linter/rules/import/no-amd.html) | import |  |  |
| [no-commonjs](https://oxc.rs/docs/guide/usage/linter/rules/import/no-commonjs) | import |  |  |
| [no-cycle](https://oxc.rs/docs/guide/usage/linter/rules/import/no-cycle.html) | import |  |  |
| [no-default-export](https://oxc.rs/docs/guide/usage/linter/rules/import/no-default-export.html) | import |  |  |
| [no-dynamic-require](https://oxc.rs/docs/guide/usage/linter/rules/import/no-dynamic-require.html) | import |  |  |
| [no-webpack-loader-syntax](https://oxc.rs/docs/guide/usage/linter/rules/import/no-webpack-loader-syntax.html) | import |  |  |
| [unambiguous](https://oxc.rs/docs/guide/usage/linter/rules/import/unambiguous.html) | import |  |  |
| [check-access](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/check-access.html) | jsdoc |  |  |
| [empty-tags](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/empty-tags.html) | jsdoc |  |  |
| [anchor-ambiguous-text](https://oxc.rs/docs/guide/usage/linter/rules/jsx_a11y/anchor-ambiguous-text.html) | jsx\_a11y |  |  |
| [no-new-require](https://oxc.rs/docs/guide/usage/linter/rules/node/no-new-require.html) | node |  |  |
| [no-process-env](https://oxc.rs/docs/guide/usage/linter/rules/node/no-process-env.html) | node |  |  |
| [bad-bitwise-operator](https://oxc.rs/docs/guide/usage/linter/rules/oxc/bad-bitwise-operator.html) | oxc |  | 💡 |
| [no-async-await](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-await.html) | oxc |  |  |
| [no-barrel-file](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-barrel-file.html) | oxc |  |  |
| [no-const-enum](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-const-enum.html) | oxc |  | 🛠️ |
| [no-optional-chaining](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-optional-chaining.html) | oxc |  |  |
| [no-rest-spread-properties](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-rest-spread-properties.html) | oxc |  |  |
| [catch-or-return](https://oxc.rs/docs/guide/usage/linter/rules/promise/catch-or-return.html) | promise |  |  |
| [spec-only](https://oxc.rs/docs/guide/usage/linter/rules/promise/spec-only.html) | promise |  |  |
| [button-has-type](https://oxc.rs/docs/guide/usage/linter/rules/react/button-has-type.html) | react |  |  |
| [forbid-dom-props](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-dom-props.html) | react |  |  |
| [forbid-elements](https://oxc.rs/docs/guide/usage/linter/rules/react/forbid-elements.html) | react |  |  |
| [jsx-filename-extension](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-filename-extension.html) | react |  | 🚧 |
| [no-danger](https://oxc.rs/docs/guide/usage/linter/rules/react/no-danger.html) | react |  |  |
| [no-unknown-property](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unknown-property.html) | react |  | 🚧 |
| [only-export-components](https://oxc.rs/docs/guide/usage/linter/rules/react/only-export-components.html) | react |  |  |
| [explicit-function-return-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-function-return-type.html) | typescript |  |  |
| [explicit-module-boundary-types](https://oxc.rs/docs/guide/usage/linter/rules/typescript/explicit-module-boundary-types.html) | typescript |  |  |
| [no-dynamic-delete](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-dynamic-delete.html) | typescript |  |  |
| [no-empty-object-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-object-type.html) | typescript |  |  |
| [no-explicit-any](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-explicit-any.html) | typescript |  | 🛠️ |
| [no-import-type-side-effects](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-import-type-side-effects.html) | typescript |  | 🛠️ |
| [no-namespace](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-namespace.html) | typescript |  |  |
| [no-non-null-asserted-nullish-coalescing](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-asserted-nullish-coalescing.html) | typescript |  |  |
| [no-non-null-assertion](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-non-null-assertion.html) | typescript |  |  |
| [no-require-imports](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-require-imports.html) | typescript |  | 🚧 |
| [no-restricted-types](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-restricted-types.html) | typescript |  | 🛠️💡 |
| [no-var-requires](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-var-requires.html) | typescript |  |  |
| [non-nullable-type-assertion-style](https://oxc.rs/docs/guide/usage/linter/rules/typescript/non-nullable-type-assertion-style.html) | typescript |  | 🚧 |
| [prefer-literal-enum-member](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-literal-enum-member.html) | typescript |  |  |
| [promise-function-async](https://oxc.rs/docs/guide/usage/linter/rules/typescript/promise-function-async.html) | typescript |  | 🚧 |
| [use-unknown-in-catch-callback-variable](https://oxc.rs/docs/guide/usage/linter/rules/typescript/use-unknown-in-catch-callback-variable.html) | typescript |  | 🚧 |
| [no-abusive-eslint-disable](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-abusive-eslint-disable.html) | unicorn |  |  |
| [no-anonymous-default-export](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-anonymous-default-export.html) | unicorn |  |  |
| [no-array-for-each](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-for-each.html) | unicorn |  | 🚧 |
| [no-array-reduce](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reduce.html) | unicorn |  |  |
| [no-document-cookie](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-document-cookie.html) | unicorn |  |  |
| [no-length-as-slice-end](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-length-as-slice-end.html) | unicorn |  | 🛠️ |
| [no-magic-array-flat-depth](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-magic-array-flat-depth.html) | unicorn |  |  |
| [no-process-exit](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-process-exit.html) | unicorn |  | 🚧 |
| [no-useless-error-capture-stack-trace](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-error-capture-stack-trace.html) | unicorn |  | 🚧 |
| [prefer-modern-math-apis](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-math-apis.html) | unicorn |  | 🚧 |
| [prefer-node-protocol](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-node-protocol.html) | unicorn |  | 🛠️ |
| [prefer-number-properties](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-number-properties.html) | unicorn |  | ⚠️🛠️️ |
| [max-props](https://oxc.rs/docs/guide/usage/linter/rules/vue/max-props.html) | vue |  |  |
| [no-import-compiler-macros](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-import-compiler-macros.html) | vue |  | ⚠️🛠️️ |
| [no-multiple-slot-args](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-multiple-slot-args.html) | vue |  | 🚧 |

## Suspicious (48): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#suspicious-48)

code that is most likely wrong or useless.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [block-scoped-var](https://oxc.rs/docs/guide/usage/linter/rules/eslint/block-scoped-var.html) | eslint |  |  |
| [no-extend-native](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extend-native.html) | eslint |  |  |
| [no-extra-bind](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-bind.html) | eslint |  | 🚧 |
| [no-new](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new.html) | eslint |  |  |
| [no-unexpected-multiline](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unexpected-multiline.html) | eslint |  | ⚠️🛠️️ |
| [no-unneeded-ternary](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unneeded-ternary.html) | eslint |  | ⚠️🛠️️ |
| [no-useless-concat](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-concat.html) | eslint |  |  |
| [no-useless-constructor](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-constructor.html) | eslint |  | 🛠️ |
| [preserve-caught-error](https://oxc.rs/docs/guide/usage/linter/rules/eslint/preserve-caught-error.html) | eslint |  | 🛠️ |
| [no-absolute-path](https://oxc.rs/docs/guide/usage/linter/rules/import/no-absolute-path.html) | import |  | 🚧 |
| [no-empty-named-blocks](https://oxc.rs/docs/guide/usage/linter/rules/import/no-empty-named-blocks.html) | import |  | 🛠️ |
| [no-named-as-default](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default.html) | import |  |  |
| [no-named-as-default-member](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-as-default-member.html) | import |  |  |
| [no-self-import](https://oxc.rs/docs/guide/usage/linter/rules/import/no-self-import.html) | import |  |  |
| [no-unassigned-import](https://oxc.rs/docs/guide/usage/linter/rules/import/no-unassigned-import.html) | import |  |  |
| [no-commented-out-tests](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-commented-out-tests.html) | jest |  |  |
| [approx-constant](https://oxc.rs/docs/guide/usage/linter/rules/oxc/approx-constant.html) | oxc |  | 💡 |
| [misrefactored-assign-op](https://oxc.rs/docs/guide/usage/linter/rules/oxc/misrefactored-assign-op.html) | oxc |  | 💡 |
| [no-async-endpoint-handlers](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-async-endpoint-handlers.html) | oxc |  |  |
| [no-this-in-exported-function](https://oxc.rs/docs/guide/usage/linter/rules/oxc/no-this-in-exported-function.html) | oxc |  |  |
| [always-return](https://oxc.rs/docs/guide/usage/linter/rules/promise/always-return.html) | promise |  |  |
| [no-multiple-resolved](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-multiple-resolved.html) | promise |  |  |
| [no-promise-in-callback](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-promise-in-callback.html) | promise |  |  |
| [iframe-missing-sandbox](https://oxc.rs/docs/guide/usage/linter/rules/react/iframe-missing-sandbox.html) | react |  | 🚧 |
| [jsx-no-comment-textnodes](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-comment-textnodes.html) | react |  |  |
| [jsx-no-script-url](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-script-url.html) | react |  | 🚧 |
| [no-namespace](https://oxc.rs/docs/guide/usage/linter/rules/react/no-namespace.html) | react |  |  |
| [react-in-jsx-scope](https://oxc.rs/docs/guide/usage/linter/rules/react/react-in-jsx-scope.html) | react |  |  |
| [style-prop-object](https://oxc.rs/docs/guide/usage/linter/rules/react/style-prop-object.html) | react |  |  |
| [no-confusing-non-null-assertion](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-non-null-assertion.html) | typescript |  | 🚧 |
| [no-extraneous-class](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-extraneous-class.html) | typescript |  | ⚠️💡 |
| [no-unnecessary-boolean-literal-compare](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-boolean-literal-compare.html) | typescript |  | 🚧 |
| [no-unnecessary-template-expression](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-template-expression.html) | typescript |  | 🚧 |
| [no-unnecessary-type-arguments](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-arguments.html) | typescript |  | 🚧 |
| [no-unnecessary-type-assertion](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-assertion.html) | typescript |  | 🚧 |
| [no-unnecessary-type-constraint](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unnecessary-type-constraint.html) | typescript |  |  |
| [no-unsafe-enum-comparison](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-enum-comparison.html) | typescript |  | 🚧 |
| [no-unsafe-type-assertion](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-type-assertion.html) | typescript |  | 🚧 |
| [consistent-function-scoping](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-function-scoping.html) | unicorn |  | 🚧 |
| [no-accessor-recursion](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-accessor-recursion.html) | unicorn |  |  |
| [no-array-reverse](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-reverse.html) | unicorn |  | 🛠️ |
| [no-array-sort](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-sort.html) | unicorn |  | 🛠️ |
| [no-instanceof-builtins](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-builtins.html) | unicorn |  | 🚧 |
| [prefer-add-event-listener](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-add-event-listener.html) | unicorn |  | 🚧 |
| [require-module-specifiers](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-specifiers.html) | unicorn |  | 🛠️ |
| [require-post-message-target-origin](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-post-message-target-origin.html) | unicorn |  | 💡 |
| [no-required-prop-with-default](https://oxc.rs/docs/guide/usage/linter/rules/vue/no-required-prop-with-default.html) | vue |  | 🚧 |
| [require-default-export](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-default-export.html) | vue |  |  |

## Pedantic (113): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#pedantic-113)

Lints which are rather strict or have occasional false positives.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [accessor-pairs](https://oxc.rs/docs/guide/usage/linter/rules/eslint/accessor-pairs.html) | eslint |  |  |
| [array-callback-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/array-callback-return.html) | eslint |  |  |
| [eqeqeq](https://oxc.rs/docs/guide/usage/linter/rules/eslint/eqeqeq.html) | eslint |  | ⚠️🛠️️ |
| [max-classes-per-file](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-classes-per-file.html) | eslint |  |  |
| [max-depth](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-depth.html) | eslint |  |  |
| [max-lines](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines.html) | eslint |  |  |
| [max-lines-per-function](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-lines-per-function.html) | eslint |  |  |
| [max-nested-callbacks](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-nested-callbacks.html) | eslint |  |  |
| [no-array-constructor](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-array-constructor.html) | eslint |  | 🛠️ |
| [no-case-declarations](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-case-declarations.html) | eslint |  |  |
| [no-constructor-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-constructor-return.html) | eslint |  |  |
| [no-else-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-else-return.html) | eslint |  | 🛠️ |
| [no-fallthrough](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-fallthrough.html) | eslint |  | 🚧 |
| [no-inline-comments](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inline-comments.html) | eslint |  |  |
| [no-inner-declarations](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-inner-declarations.html) | eslint |  |  |
| [no-lonely-if](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lonely-if.html) | eslint |  | 🚧 |
| [no-loop-func](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-loop-func.html) | eslint |  |  |
| [no-negated-condition](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-negated-condition.html) | eslint |  | 🚧 |
| [no-new-wrappers](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-wrappers.html) | eslint |  | 🛠️ |
| [no-object-constructor](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-object-constructor.html) | eslint |  | 🚧 |
| [no-promise-executor-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-promise-executor-return.html) | eslint |  |  |
| [no-prototype-builtins](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-prototype-builtins.html) | eslint |  |  |
| [no-redeclare](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-redeclare.html) | eslint |  |  |
| [no-self-compare](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-self-compare.html) | eslint |  |  |
| [no-throw-literal](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-throw-literal.html) | eslint |  | 💡 |
| [no-useless-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-return.html) | eslint |  | 🚧 |
| [no-warning-comments](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-warning-comments.html) | eslint |  |  |
| [radix](https://oxc.rs/docs/guide/usage/linter/rules/eslint/radix.html) | eslint |  | ⚠️🛠️️ |
| [require-await](https://oxc.rs/docs/guide/usage/linter/rules/eslint/require-await.html) | eslint |  | ⚠️🛠️️ |
| [sort-vars](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-vars.html) | eslint |  | 🚧 |
| [symbol-description](https://oxc.rs/docs/guide/usage/linter/rules/eslint/symbol-description.html) | eslint |  |  |
| [max-dependencies](https://oxc.rs/docs/guide/usage/linter/rules/import/max-dependencies.html) | import |  |  |
| [no-conditional-in-test](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-conditional-in-test.html) | jest |  |  |
| [require-param](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param.html) | jsdoc |  |  |
| [require-param-description](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-description.html) | jsdoc |  |  |
| [require-param-name](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-name.html) | jsdoc |  |  |
| [require-param-type](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-param-type.html) | jsdoc |  |  |
| [require-returns](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns.html) | jsdoc |  |  |
| [require-returns-description](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-description.html) | jsdoc |  |  |
| [require-returns-type](https://oxc.rs/docs/guide/usage/linter/rules/jsdoc/require-returns-type.html) | jsdoc |  |  |
| [checked-requires-onchange-or-readonly](https://oxc.rs/docs/guide/usage/linter/rules/react/checked-requires-onchange-or-readonly.html) | react |  |  |
| [jsx-no-target-blank](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-target-blank.html) | react |  |  |
| [jsx-no-useless-fragment](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-no-useless-fragment.html) | react |  | 💡 |
| [no-unescaped-entities](https://oxc.rs/docs/guide/usage/linter/rules/react/no-unescaped-entities.html) | react |  |  |
| [rules-of-hooks](https://oxc.rs/docs/guide/usage/linter/rules/react/rules-of-hooks.html) | react |  |  |
| [ban-ts-comment](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-ts-comment.html) | typescript |  | 🛠️ |
| [ban-types](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-types.html) | typescript |  | 🚧 |
| [no-confusing-void-expression](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-confusing-void-expression.html) | typescript |  | 🚧 |
| [no-deprecated](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-deprecated.html) | typescript |  |  |
| [no-misused-promises](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-misused-promises.html) | typescript |  | 🚧 |
| [no-mixed-enums](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-mixed-enums.html) | typescript |  | 🚧 |
| [no-unsafe-argument](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-argument.html) | typescript |  | 🚧 |
| [no-unsafe-assignment](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-assignment.html) | typescript |  | 🚧 |
| [no-unsafe-call](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-call.html) | typescript |  | 🚧 |
| [no-unsafe-function-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-function-type.html) | typescript |  |  |
| [no-unsafe-member-access](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-member-access.html) | typescript |  | 🚧 |
| [no-unsafe-return](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-unsafe-return.html) | typescript |  | 🚧 |
| [only-throw-error](https://oxc.rs/docs/guide/usage/linter/rules/typescript/only-throw-error.html) | typescript |  | 🚧 |
| [prefer-enum-initializers](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-enum-initializers.html) | typescript |  | 💡 |
| [prefer-includes](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-includes.html) | typescript |  | 🚧 |
| [prefer-nullish-coalescing](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-nullish-coalescing.html) | typescript |  | 🚧 |
| [prefer-promise-reject-errors](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-promise-reject-errors.html) | typescript |  | 🚧 |
| [prefer-ts-expect-error](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-ts-expect-error.html) | typescript |  | 🛠️ |
| [related-getter-setter-pairs](https://oxc.rs/docs/guide/usage/linter/rules/typescript/related-getter-setter-pairs.html) | typescript |  | 🚧 |
| [require-await](https://oxc.rs/docs/guide/usage/linter/rules/typescript/require-await.html) | typescript |  | 🚧 |
| [restrict-plus-operands](https://oxc.rs/docs/guide/usage/linter/rules/typescript/restrict-plus-operands.html) | typescript |  | 🚧 |
| [return-await](https://oxc.rs/docs/guide/usage/linter/rules/typescript/return-await.html) | typescript |  | 🚧 |
| [strict-boolean-expressions](https://oxc.rs/docs/guide/usage/linter/rules/typescript/strict-boolean-expressions.html) | typescript |  | 🚧 |
| [switch-exhaustiveness-check](https://oxc.rs/docs/guide/usage/linter/rules/typescript/switch-exhaustiveness-check.html) | typescript |  | 🚧 |
| [consistent-assert](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-assert.html) | unicorn |  | 🛠️ |
| [consistent-empty-array-spread](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-empty-array-spread.html) | unicorn |  | 💡 |
| [escape-case](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/escape-case.html) | unicorn |  | 🛠️ |
| [explicit-length-check](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/explicit-length-check.html) | unicorn |  | 🛠️ |
| [new-for-builtins](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/new-for-builtins.html) | unicorn |  |  |
| [no-array-callback-reference](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-callback-reference.html) | unicorn |  | 🚧 |
| [no-hex-escape](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-hex-escape.html) | unicorn |  | 🛠️ |
| [no-immediate-mutation](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-immediate-mutation.html) | unicorn |  | 🚧 |
| [no-instanceof-array](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-instanceof-array.html) | unicorn |  | 🛠️ |
| [no-lonely-if](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-lonely-if.html) | unicorn |  |  |
| [no-negation-in-equality-check](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-negation-in-equality-check.html) | unicorn |  | 💡 |
| [no-new-buffer](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-new-buffer.html) | unicorn |  | 🚧 |
| [no-object-as-default-parameter](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-object-as-default-parameter.html) | unicorn |  |  |
| [no-static-only-class](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-static-only-class.html) | unicorn |  | ⚠️🛠️️ |
| [no-this-assignment](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-this-assignment.html) | unicorn |  |  |
| [no-typeof-undefined](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-typeof-undefined.html) | unicorn |  | 🚧 |
| [no-unnecessary-array-flat-depth](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-flat-depth.html) | unicorn |  | 💡 |
| [no-unnecessary-array-splice-count](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-array-splice-count.html) | unicorn |  | 🛠️ |
| [no-unnecessary-slice-end](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unnecessary-slice-end.html) | unicorn |  | 🛠️ |
| [no-unreadable-iife](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-iife.html) | unicorn |  |  |
| [no-useless-promise-resolve-reject](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-promise-resolve-reject.html) | unicorn |  | 🛠️ |
| [no-useless-switch-case](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-switch-case.html) | unicorn |  | 🚧 |
| [no-useless-undefined](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-undefined.html) | unicorn |  | 🛠️ |
| [prefer-array-flat](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-flat.html) | unicorn |  | ⚠️🛠️️ |
| [prefer-array-some](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-some.html) | unicorn |  | 🛠️ |
| [prefer-at](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-at.html) | unicorn |  | ⚠️🛠️️ |
| [prefer-blob-reading-methods](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-blob-reading-methods.html) | unicorn |  | 🚧 |
| [prefer-code-point](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-code-point.html) | unicorn |  | 🛠️ |
| [prefer-date-now](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-date-now.html) | unicorn |  | 🛠️ |
| [prefer-dom-node-append](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-append.html) | unicorn |  | 🛠️ |
| [prefer-dom-node-dataset](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-dataset.html) | unicorn |  | 🚧 |
| [prefer-dom-node-remove](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-remove.html) | unicorn |  |  |
| [prefer-event-target](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-event-target.html) | unicorn |  |  |
| [prefer-math-min-max](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-min-max.html) | unicorn |  | 🛠️ |
| [prefer-math-trunc](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-math-trunc.html) | unicorn |  | 🚧 |
| [prefer-native-coercion-functions](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-native-coercion-functions.html) | unicorn |  | 🚧 |
| [prefer-prototype-methods](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-prototype-methods.html) | unicorn |  | 🛠️ |
| [prefer-query-selector](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-query-selector.html) | unicorn |  | 🛠️ |
| [prefer-regexp-test](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-regexp-test.html) | unicorn |  | 🛠️ |
| [prefer-string-replace-all](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-replace-all.html) | unicorn |  | 🛠️ |
| [prefer-string-slice](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-slice.html) | unicorn |  | 🛠️ |
| [prefer-top-level-await](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-top-level-await.html) | unicorn |  |  |
| [prefer-type-error](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-type-error.html) | unicorn |  | 🛠️ |
| [require-number-to-fixed-digits-argument](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-number-to-fixed-digits-argument.html) | unicorn |  | 🛠️ |

## Style (182): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#style-182)

Code that should be written in a more idiomatic way.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [arrow-body-style](https://oxc.rs/docs/guide/usage/linter/rules/eslint/arrow-body-style.html) | eslint |  | 🛠️ |
| [capitalized-comments](https://oxc.rs/docs/guide/usage/linter/rules/eslint/capitalized-comments.html) | eslint |  | 🛠️ |
| [curly](https://oxc.rs/docs/guide/usage/linter/rules/eslint/curly.html) | eslint |  | 🛠️ |
| [default-case-last](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-case-last.html) | eslint |  |  |
| [default-param-last](https://oxc.rs/docs/guide/usage/linter/rules/eslint/default-param-last.html) | eslint |  |  |
| [func-names](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-names.html) | eslint |  | 🛠️💡 |
| [func-style](https://oxc.rs/docs/guide/usage/linter/rules/eslint/func-style.html) | eslint |  | 🚧 |
| [grouped-accessor-pairs](https://oxc.rs/docs/guide/usage/linter/rules/eslint/grouped-accessor-pairs.html) | eslint |  | 🚧 |
| [guard-for-in](https://oxc.rs/docs/guide/usage/linter/rules/eslint/guard-for-in.html) | eslint |  |  |
| [id-length](https://oxc.rs/docs/guide/usage/linter/rules/eslint/id-length.html) | eslint |  |  |
| [init-declarations](https://oxc.rs/docs/guide/usage/linter/rules/eslint/init-declarations.html) | eslint |  |  |
| [max-params](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-params.html) | eslint |  |  |
| [max-statements](https://oxc.rs/docs/guide/usage/linter/rules/eslint/max-statements.html) | eslint |  |  |
| [new-cap](https://oxc.rs/docs/guide/usage/linter/rules/eslint/new-cap.html) | eslint |  | 🚧 |
| [no-continue](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-continue.html) | eslint |  |  |
| [no-duplicate-imports](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-duplicate-imports.html) | eslint |  | 🚧 |
| [no-extra-label](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-extra-label.html) | eslint |  | 🛠️ |
| [no-implicit-coercion](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-implicit-coercion.html) | eslint |  | 🛠️ |
| [no-label-var](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-label-var.html) | eslint |  |  |
| [no-labels](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-labels.html) | eslint |  |  |
| [no-lone-blocks](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-lone-blocks.html) | eslint |  |  |
| [no-magic-numbers](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-magic-numbers.html) | eslint |  | 🚧 |
| [no-multi-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-assign.html) | eslint |  |  |
| [no-multi-str](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-multi-str.html) | eslint |  |  |
| [no-nested-ternary](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-nested-ternary.html) | eslint |  |  |
| [no-new-func](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-new-func.html) | eslint |  |  |
| [no-return-assign](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-return-assign.html) | eslint |  | 🚧 |
| [no-script-url](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-script-url.html) | eslint |  |  |
| [no-template-curly-in-string](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-template-curly-in-string.html) | eslint |  |  |
| [no-ternary](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-ternary.html) | eslint |  |  |
| [no-useless-computed-key](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-useless-computed-key.html) | eslint |  | 🚧 |
| [operator-assignment](https://oxc.rs/docs/guide/usage/linter/rules/eslint/operator-assignment.html) | eslint |  | ⚠️🛠️️ |
| [prefer-destructuring](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-destructuring.html) | eslint |  | 🛠️ |
| [prefer-exponentiation-operator](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-exponentiation-operator.html) | eslint |  |  |
| [prefer-numeric-literals](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-numeric-literals.html) | eslint |  | 🛠️ |
| [prefer-object-has-own](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-has-own.html) | eslint |  | 🛠️ |
| [prefer-object-spread](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-object-spread.html) | eslint |  | 🛠️ |
| [prefer-promise-reject-errors](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-promise-reject-errors.html) | eslint |  |  |
| [prefer-rest-params](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-rest-params.html) | eslint |  |  |
| [prefer-spread](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-spread.html) | eslint |  |  |
| [prefer-template](https://oxc.rs/docs/guide/usage/linter/rules/eslint/prefer-template.html) | eslint |  | 🚧 |
| [sort-imports](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-imports.html) | eslint |  | 🛠️ |
| [sort-keys](https://oxc.rs/docs/guide/usage/linter/rules/eslint/sort-keys.html) | eslint |  | 🛠️ |
| [vars-on-top](https://oxc.rs/docs/guide/usage/linter/rules/eslint/vars-on-top.html) | eslint |  |  |
| [yoda](https://oxc.rs/docs/guide/usage/linter/rules/eslint/yoda.html) | eslint |  | 🛠️ |
| [consistent-type-specifier-style](https://oxc.rs/docs/guide/usage/linter/rules/import/consistent-type-specifier-style.html) | import |  | 🛠️ |
| [exports-last](https://oxc.rs/docs/guide/usage/linter/rules/import/exports-last.html) | import |  |  |
| [first](https://oxc.rs/docs/guide/usage/linter/rules/import/first.html) | import |  | 🚧 |
| [group-exports](https://oxc.rs/docs/guide/usage/linter/rules/import/group-exports.html) | import |  |  |
| [no-anonymous-default-export](https://oxc.rs/docs/guide/usage/linter/rules/import/no-anonymous-default-export.html) | import |  |  |
| [no-duplicates](https://oxc.rs/docs/guide/usage/linter/rules/import/no-duplicates.html) | import |  |  |
| [no-mutable-exports](https://oxc.rs/docs/guide/usage/linter/rules/import/no-mutable-exports.html) | import |  |  |
| [no-named-default](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-default.html) | import |  |  |
| [no-named-export](https://oxc.rs/docs/guide/usage/linter/rules/import/no-named-export.html) | import |  |  |
| [no-namespace](https://oxc.rs/docs/guide/usage/linter/rules/import/no-namespace.html) | import |  | 🚧 |
| [prefer-default-export](https://oxc.rs/docs/guide/usage/linter/rules/import/prefer-default-export.html) | import |  |  |
| [consistent-test-it](https://oxc.rs/docs/guide/usage/linter/rules/jest/consistent-test-it.html) | jest |  | 🛠️ |
| [max-expects](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-expects.html) | jest |  |  |
| [max-nested-describe](https://oxc.rs/docs/guide/usage/linter/rules/jest/max-nested-describe.html) | jest |  |  |
| [no-alias-methods](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-alias-methods.html) | jest |  | 🛠️ |
| [no-confusing-set-timeout](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-confusing-set-timeout.html) | jest |  |  |
| [no-deprecated-functions](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-deprecated-functions.html) | jest |  | 🛠️ |
| [no-done-callback](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-done-callback.html) | jest |  |  |
| [no-duplicate-hooks](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-duplicate-hooks.html) | jest |  |  |
| [no-hooks](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-hooks.html) | jest |  |  |
| [no-identical-title](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-identical-title.html) | jest |  |  |
| [no-interpolation-in-snapshots](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-interpolation-in-snapshots.html) | jest |  |  |
| [no-jasmine-globals](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-jasmine-globals.html) | jest |  | 🛠️ |
| [no-large-snapshots](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-large-snapshots.html) | jest |  |  |
| [no-mocks-import](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-mocks-import.html) | jest |  |  |
| [no-restricted-jest-methods](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-jest-methods.html) | jest |  |  |
| [no-restricted-matchers](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-restricted-matchers.html) | jest |  |  |
| [no-test-prefixes](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-prefixes.html) | jest |  | 🛠️ |
| [no-test-return-statement](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-test-return-statement.html) | jest |  |  |
| [no-untyped-mock-factory](https://oxc.rs/docs/guide/usage/linter/rules/jest/no-untyped-mock-factory.html) | jest |  | 🛠️ |
| [padding-around-test-blocks](https://oxc.rs/docs/guide/usage/linter/rules/jest/padding-around-test-blocks.html) | jest |  | 🛠️ |
| [prefer-called-with](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-called-with.html) | jest |  | 🛠️ |
| [prefer-comparison-matcher](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-comparison-matcher.html) | jest |  | 🛠️ |
| [prefer-each](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-each.html) | jest |  |  |
| [prefer-equality-matcher](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-equality-matcher.html) | jest |  |  |
| [prefer-expect-resolves](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-expect-resolves.html) | jest |  | 🛠️ |
| [prefer-hooks-in-order](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-in-order.html) | jest |  |  |
| [prefer-hooks-on-top](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-hooks-on-top.html) | jest |  |  |
| [prefer-jest-mocked](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-jest-mocked.html) | jest |  | 🛠️ |
| [prefer-lowercase-title](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-lowercase-title.html) | jest |  | 🛠️ |
| [prefer-mock-promise-shorthand](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-mock-promise-shorthand.html) | jest |  | 🛠️ |
| [prefer-spy-on](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-spy-on.html) | jest |  | 🛠️ |
| [prefer-strict-equal](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-strict-equal.html) | jest |  | 🛠️ |
| [prefer-to-be](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-be.html) | jest |  | 🛠️ |
| [prefer-to-contain](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-contain.html) | jest |  | 🛠️ |
| [prefer-to-have-been-called](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called.html) | jest |  | 🛠️ |
| [prefer-to-have-been-called-times](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-been-called-times.html) | jest |  | 🛠️ |
| [prefer-to-have-length](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-to-have-length.html) | jest |  | 🛠️ |
| [prefer-todo](https://oxc.rs/docs/guide/usage/linter/rules/jest/prefer-todo.html) | jest |  | 🛠️ |
| [require-hook](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-hook.html) | jest |  |  |
| [require-top-level-describe](https://oxc.rs/docs/guide/usage/linter/rules/jest/require-top-level-describe.html) | jest |  |  |
| [global-require](https://oxc.rs/docs/guide/usage/linter/rules/node/global-require.html) | node |  |  |
| [no-exports-assign](https://oxc.rs/docs/guide/usage/linter/rules/node/no-exports-assign.html) | node |  | 🛠️ |
| [avoid-new](https://oxc.rs/docs/guide/usage/linter/rules/promise/avoid-new.html) | promise |  |  |
| [no-nesting](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-nesting.html) | promise |  | 🚧 |
| [no-return-wrap](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-wrap.html) | promise |  | 🚧 |
| [param-names](https://oxc.rs/docs/guide/usage/linter/rules/promise/param-names.html) | promise |  |  |
| [prefer-await-to-callbacks](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-callbacks.html) | promise |  |  |
| [prefer-await-to-then](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-await-to-then.html) | promise |  |  |
| [prefer-catch](https://oxc.rs/docs/guide/usage/linter/rules/promise/prefer-catch.html) | promise |  | 🚧 |
| [jsx-boolean-value](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-boolean-value.html) | react |  | 🛠️ |
| [jsx-curly-brace-presence](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-curly-brace-presence.html) | react |  | 🛠️ |
| [jsx-fragments](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-fragments.html) | react |  | 🛠️ |
| [jsx-handler-names](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-handler-names.html) | react |  |  |
| [jsx-max-depth](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-max-depth.html) | react |  |  |
| [jsx-pascal-case](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-pascal-case.html) | react |  |  |
| [jsx-props-no-spreading](https://oxc.rs/docs/guide/usage/linter/rules/react/jsx-props-no-spreading.html) | react |  |  |
| [no-redundant-should-component-update](https://oxc.rs/docs/guide/usage/linter/rules/react/no-redundant-should-component-update.html) | react |  |  |
| [no-set-state](https://oxc.rs/docs/guide/usage/linter/rules/react/no-set-state.html) | react |  |  |
| [prefer-es6-class](https://oxc.rs/docs/guide/usage/linter/rules/react/prefer-es6-class.html) | react |  |  |
| [self-closing-comp](https://oxc.rs/docs/guide/usage/linter/rules/react/self-closing-comp.html) | react |  | 🛠️ |
| [state-in-constructor](https://oxc.rs/docs/guide/usage/linter/rules/react/state-in-constructor.html) | react |  |  |
| [adjacent-overload-signatures](https://oxc.rs/docs/guide/usage/linter/rules/typescript/adjacent-overload-signatures.html) | typescript |  |  |
| [array-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/array-type.html) | typescript |  | 🛠️ |
| [ban-tslint-comment](https://oxc.rs/docs/guide/usage/linter/rules/typescript/ban-tslint-comment.html) | typescript |  | 🛠️ |
| [consistent-generic-constructors](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-generic-constructors.html) | typescript |  | 🚧 |
| [consistent-indexed-object-style](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-indexed-object-style.html) | typescript |  | 🛠️ |
| [consistent-type-definitions](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-definitions.html) | typescript |  | ⚠️🛠️️ |
| [consistent-type-imports](https://oxc.rs/docs/guide/usage/linter/rules/typescript/consistent-type-imports.html) | typescript |  | 🛠️ |
| [no-empty-interface](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-empty-interface.html) | typescript |  |  |
| [no-inferrable-types](https://oxc.rs/docs/guide/usage/linter/rules/typescript/no-inferrable-types.html) | typescript |  | 💡 |
| [prefer-for-of](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-for-of.html) | typescript |  | 🚧 |
| [prefer-function-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-function-type.html) | typescript |  | 🛠️ |
| [prefer-namespace-keyword](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-namespace-keyword.html) | typescript |  | 🛠️ |
| [prefer-reduce-type-parameter](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-reduce-type-parameter.html) | typescript |  | 🚧 |
| [prefer-return-this-type](https://oxc.rs/docs/guide/usage/linter/rules/typescript/prefer-return-this-type.html) | typescript |  | 🚧 |
| [catch-error-name](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/catch-error-name.html) | unicorn |  | 🛠️ |
| [consistent-date-clone](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-date-clone.html) | unicorn |  | 🛠️ |
| [consistent-existence-index-check](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/consistent-existence-index-check.html) | unicorn |  | 🛠️ |
| [empty-brace-spaces](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/empty-brace-spaces.html) | unicorn |  | 🛠️ |
| [error-message](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/error-message.html) | unicorn |  |  |
| [filename-case](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/filename-case.html) | unicorn |  |  |
| [no-array-method-this-argument](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-array-method-this-argument.html) | unicorn |  | 🚧 |
| [no-await-expression-member](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-await-expression-member.html) | unicorn |  | ⚠️🛠️️ |
| [no-console-spaces](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-console-spaces.html) | unicorn |  | 🛠️ |
| [no-nested-ternary](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-nested-ternary.html) | unicorn |  | 🛠️ |
| [no-null](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-null.html) | unicorn |  | 🛠️ |
| [no-unreadable-array-destructuring](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-unreadable-array-destructuring.html) | unicorn |  |  |
| [no-useless-collection-argument](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-useless-collection-argument.html) | unicorn |  | 🚧 |
| [no-zero-fractions](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/no-zero-fractions.html) | unicorn |  | 🛠️ |
| [number-literal-case](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/number-literal-case.html) | unicorn |  | 🛠️ |
| [numeric-separators-style](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/numeric-separators-style.html) | unicorn |  | 🛠️ |
| [prefer-array-index-of](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-array-index-of.html) | unicorn |  | 🚧 |
| [prefer-bigint-literals](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-bigint-literals.html) | unicorn |  | 🛠️ |
| [prefer-class-fields](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-class-fields.html) | unicorn |  | 🛠️💡 |
| [prefer-classlist-toggle](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-classlist-toggle.html) | unicorn |  | 🛠️ |
| [prefer-default-parameters](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-default-parameters.html) | unicorn |  | 🚧 |
| [prefer-dom-node-text-content](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-dom-node-text-content.html) | unicorn |  | 🛠️ |
| [prefer-global-this](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-global-this.html) | unicorn |  | 🚧 |
| [prefer-includes](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-includes.html) | unicorn |  | 🚧 |
| [prefer-keyboard-event-key](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-keyboard-event-key.html) | unicorn |  | 🛠️ |
| [prefer-logical-operator-over-ternary](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-logical-operator-over-ternary.html) | unicorn |  | 🚧 |
| [prefer-modern-dom-apis](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-modern-dom-apis.html) | unicorn |  | 🚧 |
| [prefer-negative-index](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-negative-index.html) | unicorn |  | 🛠️ |
| [prefer-object-from-entries](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-object-from-entries.html) | unicorn |  | 🚧 |
| [prefer-optional-catch-binding](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-optional-catch-binding.html) | unicorn |  | 🛠️ |
| [prefer-reflect-apply](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-reflect-apply.html) | unicorn |  |  |
| [prefer-response-static-json](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-response-static-json) | unicorn |  | 🚧 |
| [prefer-spread](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-spread.html) | unicorn |  | 🛠️ |
| [prefer-string-raw](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-raw.html) | unicorn |  | 🛠️ |
| [prefer-string-trim-start-end](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-string-trim-start-end.html) | unicorn |  | 🛠️ |
| [prefer-structured-clone](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/prefer-structured-clone.html) | unicorn |  | 💡 |
| [require-array-join-separator](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-array-join-separator.html) | unicorn |  | 🛠️ |
| [require-module-attributes](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/require-module-attributes.html) | unicorn |  | 💡 |
| [switch-case-braces](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/switch-case-braces.html) | unicorn |  | 🛠️ |
| [text-encoding-identifier-case](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/text-encoding-identifier-case.html) | unicorn |  | 🛠️ |
| [throw-new-error](https://oxc.rs/docs/guide/usage/linter/rules/unicorn/throw-new-error.html) | unicorn |  | 🛠️ |
| [consistent-test-filename](https://oxc.rs/docs/guide/usage/linter/rules/vitest/consistent-test-filename.html) | vitest |  |  |
| [no-import-node-test](https://oxc.rs/docs/guide/usage/linter/rules/vitest/no-import-node-test.html) | vitest |  | 🛠️ |
| [prefer-called-times](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-called-times.html) | vitest |  | 🛠️ |
| [prefer-to-be-falsy](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-falsy.html) | vitest |  | 🛠️ |
| [prefer-to-be-object](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-object.html) | vitest |  | 🛠️ |
| [prefer-to-be-truthy](https://oxc.rs/docs/guide/usage/linter/rules/vitest/prefer-to-be-truthy.html) | vitest |  | 🛠️ |
| [define-emits-declaration](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-emits-declaration.html) | vue |  | 🚧 |
| [define-props-declaration](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-declaration.html) | vue |  |  |
| [define-props-destructuring](https://oxc.rs/docs/guide/usage/linter/rules/vue/define-props-destructuring.html) | vue |  |  |
| [require-typed-ref](https://oxc.rs/docs/guide/usage/linter/rules/vue/require-typed-ref.html) | vue |  |  |

## Nursery (9): [​](https://oxc.rs/docs/guide/usage/linter/rules.html#nursery-9)

New lints that are still under development.

| Rule name | Source | Default | Fixable? |
| --- | --- | --- | --- |
| [getter-return](https://oxc.rs/docs/guide/usage/linter/rules/eslint/getter-return.html) | eslint |  |  |
| [no-misleading-character-class](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-misleading-character-class.html) | eslint |  |  |
| [no-undef](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-undef.html) | eslint |  |  |
| [no-unreachable](https://oxc.rs/docs/guide/usage/linter/rules/eslint/no-unreachable.html) | eslint |  |  |
| [export](https://oxc.rs/docs/guide/usage/linter/rules/import/export.html) | import |  |  |
| [named](https://oxc.rs/docs/guide/usage/linter/rules/import/named.html) | import |  |  |
| [branches-sharing-code](https://oxc.rs/docs/guide/usage/linter/rules/oxc/branches-sharing-code.html) | oxc |  |  |
| [no-return-in-finally](https://oxc.rs/docs/guide/usage/linter/rules/promise/no-return-in-finally.html) | promise |  |  |
| [require-render-return](https://oxc.rs/docs/guide/usage/linter/rules/react/require-render-return.html) | react |  |  |
