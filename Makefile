# shared code copy
# TODO create a linux version.
# This is a hack! but it works. I tried creating a shared volume and having both
# docker containers mount it. but then VS Code doesn't know the code exists and
# can't provide help when things are borked.
shared:
	@robocopy '.\common\shared-code\' '.\api\internal\shared-code' /purge /E
	@robocopy '.\common\shared-code\' '.\web\internal\shared-code' /purge /E

# startUP the datastore service
# NOTE to see logs take off -d flag OR
# docker logs -f datastore
ds-up:
	@docker compose up --remove-orphans datastore --detach

# connect to an already running datastore service
ds-bash:
	@docker compose exec datastore bash

# Shutdown the datastore service
ds-down:
	@docker compose down datastore

# **************************************************

api-build:
#	docker compose build api --no-cache
	@docker compose build api

api-bash: shared
	@docker compose run --remove-orphans --service-ports api bash

api-watch: shared
# results in same thing. Leave in case I want later for some kind of troubleshooting.
#	@docker compose run --remove-orphans --service-ports  api bash -c "make watch"
	@docker compose up --remove-orphans api

# **************************************************

web-build:
#	docker compose build web --no-cache
	@docker compose build web

web-bash: shared
	@docker compose run --remove-orphans --service-ports web bash

web-watch: shared
# results in same thing. Leave in case I want later for some kind of troubleshooting.
#	@docker compose run --remove-orphans --service-ports  api bash -c "make watch"
	@docker compose up --remove-orphans web

# **************************************************

common-build:
	@docker compose build common

common-bash:
	@docker compose run --remove-orphans --service-ports common bash

# **************************************************

all-up:
	@docker compose up --remove-orphans

all-down:
	@docker compose down

# **************************************************

.PHONY: ds-up ds-bash ds-down api-build api-bash api-watch all-up all-down
