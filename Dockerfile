FROM centos:latest
LABEL maintainer="NewMountain"

# Create a non-root user
RUN useradd app

# Create an app directory
RUN mkdir /app

# Copy the binary
COPY /binaries /app

# Change default directory
WORKDIR /app

# Make permissions more permissive for app
RUN chmod -R 777 /app

# Change running user
USER app

# Expose port
EXPOSE 1323

# Run this thing!
CMD ./app
