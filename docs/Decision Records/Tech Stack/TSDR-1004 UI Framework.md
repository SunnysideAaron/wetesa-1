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

Options:
https://www.reddit.com/r/webdev/comments/mfnxnj/why_your_website_should_work_without_javascript/

-  [Gio](https://gioui.org/)
- https://fyne.io/
    -https://www.reddit.com/r/golang/comments/1jh60pr/we_are_developing_a_fyne_crud_gui_generator/
- wails io
- Angular
- Flutter
  - https://www.reddit.com/r/FlutterDev/comments/1j2vmxj/develop_the_business_logic_first_approach/
  - https://www.reddit.com/r/FlutterDev/comments/1jayrqx/the_final_word_on_flutter_architecture/
- React
  - https://github.com/alan2207/bulletproof-react
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
2025-02-09 svelt 5?
		react vs vue vs flutter vs svelt
		Hotwire - html over the wire?
		Other?
		Echo
		Plush templating (2019?)
		buffalo 2019
		gorrila mocks
		
	look up what HTMX is.
		fast? use with templ, hotwire?

https://www.reddit.com/r/webdev/comments/1jjfd32/whats_your_favorite_modern_web_development_stack/

### go's html/template package

https://stackoverflow.com/questions/36617949/how-to-use-base-template-file-for-golang-html-template

### Other options

- https://www.reddit.com/r/golang/comments/1j1kvl1/any_golang_libraries_to_build_simple_crud_uis/
- https://www.reddit.com/r/Frontend/comments/1j9g6f8/all_frontend_developers_lets_make_the_most/

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