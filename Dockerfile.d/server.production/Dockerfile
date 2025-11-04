FROM ubuntu:24.04 AS builder

LABEL maintainer="docker@mazgi.com"

# https://github.com/amazon-gamelift/amazon-gamelift-servers-go-server-sdk/releases
ARG GAMELIFT_GO_SERVER_SDK_VERSION=5.2.1

# Set in non-interactive mode.
ARG DEBIAN_FRONTEND=noninteractive
# https://docs.docker.com/engine/reference/builder/#automatic-platform-args-in-the-global-scope
ARG TARGETARCH
ARG TARGETPLATFORM
ARG GID=0
ARG UID=0

RUN echo 'apt::install-recommends "false";' > /etc/apt/apt.conf.d/no-install-recommends\
  && apt-get update\
#   && :
# RUN :\
  && : Set up locales\
  && apt-get install --assume-yes locales procps dialog\
  && echo 'en_US.UTF-8 UTF-8' > /etc/locale.gen\
  && locale-gen\
#   && :
# RUN :\
  && : Install basic packages\
  && apt-get install --assume-yes apt-transport-https ca-certificates gnupg gnupg2 unzip\
  && apt-get install --assume-yes ca-certificates curl\
  && apt-get install --assume-yes golang-go\
#   && :
# RUN :\
  && : Clean up\
  && apt-get clean autoclean\
  && apt-get autoremove --yes\
  && rm -rf /tmp/*\
  && rm -rf /var/lib/apt/\
  && rm -rf /var/log/*\
  && :

COPY Dockerfile.d/server.production/rootfs /
COPY server/ /workspace/server/
RUN :\
  # Build the server binary
  && cd /workspace/server/\
  && go build -o /usr/local/bin/server\
  && :
COPY SdkGoWrapper/ /workspace/SdkGoWrapper/
RUN :\
  && curl -o /tmp/gamelift-server-sdk.zip -fsSL https://github.com/amazon-gamelift/amazon-gamelift-servers-go-server-sdk/releases/download/v${GAMELIFT_GO_SERVER_SDK_VERSION}/GameLift-Go-ServerSDK-${GAMELIFT_GO_SERVER_SDK_VERSION}.zip\
  && unzip -d /workspace/SdkGoWrapper/gamelift-server-sdk /tmp/gamelift-server-sdk.zip\
  && rm -f /tmp/gamelift-server-sdk.zip\
  # Build the wrapper binary
  && cd /workspace/SdkGoWrapper/\
  && go mod tidy\
  && go build -o /usr/local/bin/gameliftwrapper\
  && :

FROM ubuntu:22.04 AS runner
ARG DEBIAN_FRONTEND=noninteractive
ARG APP_USER_GID=10000
ARG APP_USER_UID=10000
RUN :\
  && : Create the app user: UID=${APP_USER_UID}, GID=${APP_USER_GID}\
  && groupadd --gid ${APP_USER_GID} app\
  && useradd --uid ${APP_USER_UID} --gid app --shell /bin/bash --create-home app\
  #   && :
  # RUN :\
  && : Create the working directory\
  && mkdir /workspace\
  && chown -R app:app /workspace\
  #   && :
  # RUN :\
  && : Install basic packages\
  && echo 'apt::install-recommends "false";' > /etc/apt/apt.conf.d/no-install-recommends\
  && apt-get update\
  && apt-get install --assume-yes\
  tini\
  ca-certificates\
  #   && :
  # RUN :\
  && : Clean up\
  && apt-get clean autoclean\
  && apt-get autoremove --yes\
  && rm -rf /tmp/*\
  && rm -rf /var/lib/apt/\
  && rm -rf /var/log/*\
  && :
ENTRYPOINT ["/usr/bin/tini", "--"]
WORKDIR /workspace
USER app
ENV PORT=8080
EXPOSE ${PORT}
COPY --from=builder --chown=app:app /usr/local/bin /usr/local/bin/
CMD [ "/usr/local/bin/wrapper.sh" ]
