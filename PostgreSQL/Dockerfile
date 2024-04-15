# Usa la imagen oficial de PostgreSQL
FROM postgres:latest

# Variables de entorno para configurar la base de datos
ENV POSTGRES_USER=brainy
ENV POSTGRES_PASSWORD=brainy
ENV POSTGRES_DB=brainy

# Configuración de locales
RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8
ENV LANG en_US.utf8

# Copiar script de inicialización y ejecutarlo
COPY init.sql /docker-entrypoint-initdb.d/
