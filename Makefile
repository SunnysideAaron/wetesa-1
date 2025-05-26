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

api-bash:
	@docker compose run --remove-orphans --service-ports api bash

api-watch:
# results in same thing. Leave in case I want later for some kind of troubleshooting.
#	@docker compose run --remove-orphans --service-ports  api bash -c "make watch"
	@docker compose up --remove-orphans api

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
