db:
	docker run -v nails:/var/lib/postgresql/data/ -p "5432:5432" -e POSTGRES_PASSWORD=asdbnm321 -e POSTGRES_USER=kr -e POSTGRES_DB=nails -d postgres:14.2
redis:
	docker run -p 6379:6379 -d redis