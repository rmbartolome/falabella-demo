FROM postgres:12.3
MAINTAINER Roberto M Bartolome

ENV TZ=America/Caracas

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=root.demo
ENV POSTGRES_DB=falabella

COPY scripts /docker-entrypoint-initdb.d/
RUN chmod -R 755 /docker-entrypoint-initdb.d/*