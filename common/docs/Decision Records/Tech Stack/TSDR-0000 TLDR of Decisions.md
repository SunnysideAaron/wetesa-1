# TSDR-0000  TL;DR of Decisions

- All Decisions from wetesa-0
  - These are copied from wetesa-0 so AI can find them. Don't edit them here.
    Create a new decision record.
- TSDR-1001 ORM
  - Don't use an ORM.
- TSDR-1002 DB migrations
  - **PENDING**
- TSDR-1003 CSS
  - **PENDING**  
- TSDR-1004 UI Framework
  - **PENDING**  
- TSDR-1005 Limited Javascript
  - **PENDING**  
- TSDR-1006 Documentation
  - **PENDING**  
- TSDR-1007 UI Design - Theme
  - **PENDING**  
- TSDR-1008 Debugging Tools
  - **PENDING**  
- TSDR-1009 Monitoring
  - **PENDING**  
- TSDR-1010 Code Generation
  - Don't do it.
- TSDR-1011 Limited Javascript
  - **PENDING**  
- TSDR-1012 Javascript Libraries
  - **PENDING** 
- TSDR-1013 Template Engine
  - **PENDING**  
- TSDR-1014 Maps
  - **PENDING**  
- TSDR-1015 Caching
  - **PENDING**  
- TSDR-1016 Messaging
  - **PENDING**  
- TSDR-1017 slog pretty handler
  - Keep our home brew from Wetesa-0.  
- TSDR-1018 pgx Packages
  - **PENDING**


## TODO

 - simplerr?, other error handling packages

customer id in route? or user login?

pagination
  - what if data changes between pages? / orginal query? ie data gets changed after initial pull. do we care?

starting expensive work
  then end point returns response "runing url" / status OK
  then comes back with url to check status of process
  so kicks off work. but doesn't do work.

can all (err Error) returns be named so don't have initialize?

https://github.com/vanclief/ez
  better error handling
  different types of errors: user, app, dev debug
