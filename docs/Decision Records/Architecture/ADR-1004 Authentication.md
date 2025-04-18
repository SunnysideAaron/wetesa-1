# ADR-007 Authentication

## Status

Accepted, Proposed, Deprecated or Superseded (list DR)

## Context

## Decision



## Why / Notes

https://www.reddit.com/r/learnprogramming/comments/1ji07q2/whats_a_simple_feature_that_requires_a_lot_of/

## Consequences

https://www.reddit.com/r/golang/comments/1j4lt7o/production_ready_auth_server_examples/
- https://www.reddit.com/r/golang/comments/1jch2ts/sessionbased_authentication_in_go/

https://www.reddit.com/r/softwarearchitecture/comments/ne01eg/why_do_many_applications_sometimes_have_separate/

https://www.reddit.com/r/webdev/comments/1fsg4z6/is_a_login_system_still_a_taboo_for_amateur/

idea: auth and no auth route groups so harder to mix up.

https://www.reddit.com/r/golang/comments/1gfnijj/faq_what_are_the_best_authentication_and/


hmac based one time passwords
time based one time passwords

hotp
totp 

both have IETF RC documents

https://www.reddit.com/r/golang/comments/1jsqdnq/should_i_build_a_simple_auth_service_in_go/

## Other Options

Possibilities:
- just basic sessions
- [Awesome Go - Authentication and OAuth](https://github.com/avelino/awesome-go?tab=readme-ov-file#authentication-and-oauth)
- https://github.com/markbates/goth
  - maybe Goth for Google/Apple authentication/authorization 
- https://www.reddit.com/r/golang/comments/idsvuv/what_do_you_use_for_authentication_in_golang/
- [gotth-auth](https://github.com/lordaris/gotth-auth)
  - Example we don't use tech stack


Not an option:

