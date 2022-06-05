#!/usr/bin/env bash
# Bootstrap the Vagrant environment

# Install dependencies, set locales, set the timezone
sudo apt-get update
sudo apt-get install -y --no-install-recommends \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg \
  language-pack-en language-pack-ru \
  lsb-release \
  make \
  software-properties-common
sudo update-locale LANG=en_US.UTF-8 LANGUAGE=en_US:en LC_ALL=en_US.UTF-8
sudo timedatectl set-timezone Europe/Moscow

# Add Docker’s official GPG key
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

# Install the Docker repository
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# Install the latest version of Docker
sudo apt-get update
sudo apt-get install -y --no-install-recommends docker-ce docker-ce-cli containerd.io docker-compose-plugin

# Allow executing docker without sudo
sudo usermod -aG docker ${USER}

# Install Task as Make alternative
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d

echo 'All set, rock on!'
