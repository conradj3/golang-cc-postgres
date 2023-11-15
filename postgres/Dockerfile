# Use the latest version of the official PostgreSQL image as the base image
FROM postgres:latest

# Copy the SQL script that creates a table into the container's initialization directory
COPY create_table.sql /docker-entrypoint-initdb.d/create_table.sql

# Set environment variables for the PostgreSQL user, password, and database name
# User and password are defaulted use docker secrets or a secret way to inject in your container builds. (Demo Purposes Only)
ENV POSTGRES_USER demo
ENV POSTGRES_PASSWORD pass
ENV POSTGRES_DB queue


# Expose port 5432, which is the default port used by PostgreSQL
EXPOSE 5432