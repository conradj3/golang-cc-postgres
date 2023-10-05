FROM postgres:latest

COPY create_table.sql /docker-entrypoint-initdb.d/create_table.sql

ENV POSTGRES_USER user
ENV POSTGRES_PASSWORD password
ENV POSTGRES_DB queue

EXPOSE 5432