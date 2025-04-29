# TSDR-1004 UI Framework

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

## Context

AS of 2022-MAR-19 we want 3 UIs used 6 places.
- Bare HTML - No JS no CSS (limited small CSS so just basic pretty. WORKS IF CSS OFF)
  - for low bandwidth, low power machines
- full js or htmx or typescript site
- mobile app

develop for small mobile screen first then scale up from there is much easier

test on 
- small phone (Iphone 5s? iOS 12) (lowest android?)
- tablet
- laptop
- small desktop
- med destop
- massive screen.

Why a no Javascript no CSS site?
- employees away from infrastructure. ie low bandwidth
- using older machines

- ******** SEE ME!!!
- https://digitallytailored.github.io/Classless.css/
  - https://www.reddit.com/r/webdev/comments/1je2diy/made_a_dropin_css_framework_that_transforms_bare/

- sites have to work on desktops and mobile
- mobile has to work on android and apple.

We want:
- Web UI
  - Safari (mac)
  - Chrome
  - Firefox
  - Edge
  - Internet Explorer
  - Others?
- Mobile Web UI
- Android App
- Apple App


We do not want:
- Desktop app
  - Windows
  - Mac
  - Linux

Typescript for Javascript??
include versions of browser / operationg system / phone version in list
- https://www.reddit.com/r/typescript/comments/1jrqmjl/once_you_learn_typescript_you_never_go_back_to/


Options:
https://www.reddit.com/r/webdev/comments/mfnxnj/why_your_website_should_work_without_javascript/

https://www.reddit.com/r/golang/comments/1hwlxeq/faq_whats_the_best_way_to_do_html_templating/

https://www.reddit.com/r/golang/comments/1g1k03p/i_found_the_best_web_dev_stack_for_golang/

HTMX + Alpine.js
  - seems to be a simi popluar combo

Datastar instead of htmx?
  - too new?

htmx + tailwind?

htmx
https://www.reddit.com/r/htmx/comments/1jtv7m1/what_is_dead_may_never_die/
https://www.reddit.com/r/htmx/comments/1jvrkjv/bootstrap_htmx_is_fucking_awesome/
https://www.reddit.com/r/htmx/comments/1jvynk1/i_lost_sleep_reading_essays_on_the_htmx_website/

- https://github.com/cogentcore/core
-  [Gio](https://gioui.org/)
- https://fyne.io/
    -https://www.reddit.com/r/golang/comments/1jh60pr/we_are_developing_a_fyne_crud_gui_generator/
- wails io
https://www.reddit.com/r/golang/comments/1k0t8y6/wails_is_it_still_gaining_momentum_for_go_desktop/

- Angular
- Flutter
  - https://www.reddit.com/r/FlutterDev/comments/1j2vmxj/develop_the_business_logic_first_approach/
  - https://www.reddit.com/r/FlutterDev/comments/1jayrqx/the_final_word_on_flutter_architecture/

  https://www.reddit.com/r/FlutterDev/comments/1jtjliu/what_are_your_favorites_flutter_packages_that_you/
  https://www.reddit.com/r/FlutterDev/comments/1jwm4nn/i_made_a_hidden_inapp_debug_view_for_flutter_apps/
   https://www.reddit.com/r/FlutterDev/comments/1jxbf8q/whats_a_concept_you_understand_really_well_that/

- React
  - https://github.com/alan2207/bulletproof-react
  - https://www.reddit.com/r/FlutterDev/comments/1jrvde8/performance_showdown_flutter_vs_react_native_vs/
  [React Native Isn't as Popular as You Think](https://www.youtube.com/watch?app=desktop&v=E3Yjx0fFeaA&t=1s)
    - in end of vid has sources of where pulled some stats from. worth looking up

- React Native
- Svelt
  - https://www.reddit.com/r/sveltejs/comments/1j2h7du/is_there_anything_what_you_dont_like_in_sveltekit/
  - https://www.reddit.com/r/sveltejs/comments/1jg56is/the_best_sveltekit_codebase_ive_ever_seen/
- Svelt Native
- tmple - htmx
  - htmx is a javascript library for go? allows replacement of parts of page. 
  - essays on htmx website
  - fun to use?
  - is not json?
  - is htmx on mobile app? no?
- gomponent
  - alternative to templ. dev says it's stable. more use of go code. tmple more use of html
- Vue
- https://www.reddit.com/r/htmx/comments/1jbqdm2/dont_sleep_on_daisyui_especially_for_htmx/

UI Website login
	Which UI?
		qor	https://github.com/qor/qor
		Templ - templating in go.
			has lsp, code completion
		Fyne - material design, movern ui look feel
		Wails - ? bridge between web and go desktop, rest / vue?
		Flutter (tech stack from google)
       https://www.reddit.com/r/FlutterDev/comments/1jspg4h/whats_flutter_like_for_a_ui_newbie/

        https://www.reddit.com/r/FlutterDev/comments/1jv1iad/is_the_future_for_macpc_flutter_apps_bright_or_not/
       https://www.reddit.com/r/FlutterDev/comments/1jv2efg/coming_back_after_a_few_years_how_much_has_changed/
      https://www.reddit.com/r/FlutterDev/comments/1jw2afz/sincere_question_why_would_you_use_flutter_for/


2025-02-09 svelt 5?
		react vs vue vs flutter vs svelt
		Hotwire - html over the wire?
		Other?
		Echo
		Plush templating (2019?)
		buffalo 2019
		gorrila mocks
		
https://www.reddit.com/r/htmx/comments/1js5t7x/htmx_vs_sveltejs_what_are_pros_and_cons_of_both/
https://www.reddit.com/r/sveltejs/comments/1juaepm/svelte_and_ai_coding/

	look up what HTMX is.
		fast? use with templ, hotwire?

https://www.reddit.com/r/webdev/comments/1jjfd32/whats_your_favorite_modern_web_development_stack/

### go's html/template package

https://stackoverflow.com/questions/36617949/how-to-use-base-template-file-for-golang-html-template

### Other options

- https://www.reddit.com/r/golang/comments/1j1kvl1/any_golang_libraries_to_build_simple_crud_uis/
- https://www.reddit.com/r/Frontend/comments/1j9g6f8/all_frontend_developers_lets_make_the_most/


- [Is This Tech Dead?](https://www.isthistechdead.com/)

## Decision



## Why / Notes



## Consequences



## Other Options

Possibilities:
- htmx
  - [FULL Introduction To HTMX Using Golang](https://www.youtube.com/watch?v=x7v6SNIgJpE) by ThePrimeagen
  - specifically says he just started in go. Take with grain of salt.
- https://github.com/avelino/awesome-go?tab=readme-ov-file#template-engines




Not an option:

Desktop GUI
- https://github.com/avelino/awesome-go?tab=readme-ov-file#gui