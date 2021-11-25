FROM mcr.microsoft.com/dotnet/sdk:6.0

# Arguments
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG USER_SHELL=/bin/bash

# Create the user with sudo support
#   https://code.visualstudio.com/remote/advancedcontainers/add-nonroot-user
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID --shell $USER_SHELL -m $USERNAME \
    && apt-get update \
    && DEBIAN_FRONTEND=noninteractive apt-get install -yq sudo \
    && echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
    && chmod 0440 /etc/sudoers.d/$USERNAME

# Install tools
RUN DEBIAN_FRONTEND=noninteractive apt-get install -yq zip unzip make vim nkf tree

# Change user
USER $USERNAME

CMD [ "dotnet", "--info" ]
