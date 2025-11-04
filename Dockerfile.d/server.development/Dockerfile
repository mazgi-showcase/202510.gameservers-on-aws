FROM ghcr.io/mazgi/go-app.development

LABEL org.opencontainers.image.source="https://github.com/mazgi/container-images/blob/main/Dockerfile.d/go-app.development/customize-example.Dockerfile"

ARG GID=0
ARG UID=0

HEALTHCHECK --interval=2s --timeout=1s --retries=2 --start-period=5s\
 CMD true

RUN :\
  # Create a user for development who has the same UID and GID as you.
  && useradd --comment '' --create-home --gid users --uid ${UID} developer\
  && groupadd --gid ${GID} developer\
  # Append docker group
  && bash -c "test -n \"${DOCKER_GID}\" && groupadd --gid ${DOCKER_GID} docker"\
  && usermod --append --groups docker developer 2> /dev/null || true\
  && echo 'developer ALL=(ALL) NOPASSWD: ALL' > /etc/sudoers.d/grant-all-without-password-to-developer\
  && :
