# large-file-processing

# ChungTM

# Setup Postgresql
1. Pull image: docker pull postgres
2. Start an instance: docker run --name debug-postgres -p 5432:5432 -e POSTGRES_PASSWORD=chungtm -d postgres
3. Access to container: docker exec -it <debug-postgres> psql -U postgres initdb
4. Build image: docker build --tag large-file-processing . 
5. Run container: docker run -d -p 8080:8080 -p 8081:8081 large-file-processing
   Run debug mode: docker run -p 8080:8080 -p 8081:8081 large-file-processing