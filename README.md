# large-file-processing

# ChungTM

# Setup Postgresql
1. Pull image: docker pull postgres
2. Start an instance: docker run --name debug-postgres -p 5432:5432 -e POSTGRES_PASSWORD=chungtm -d postgres
3. Access to container: docker exec -it <debug-postgres> psql -U postgres initdb