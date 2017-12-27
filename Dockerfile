#
# Test runner environment that includes an installation of the Chrome
# headless browser
#

FROM golang:1.9

# System setup
ENV DEBIAN_FRONTEND=noninteractive \
    TERM=xterm \
    TIMEZONE=UTC

    # Basic .bashrc
RUN echo 'alias ll="ls -laF"' >> /root/.bashrc \
    && echo 'alias e="exit"' >> /root/.bashrc \
    && echo 'alias cls="clear"' >> /root/.bashrc

    # System software
RUN apt-get -qqy update \
    && apt-get -qqy --no-install-recommends install \
        ca-certificates \
        gnupg \
        libcap2-bin \
        tzdata \
        wget

    # System configuration
RUN echo $TIMEZONE > /etc/timezone \
    && DEBCONF_NONINTERACTIVE_SEEN=true dpkg-reconfigure --frontend noninteractive tzdata \
    && go get -u github.com/golang/dep/cmd/dep

# Install Chrome
ARG CHROME_VERSION="google-chrome-stable"
ARG CHROME_DRIVER_VERSION="latest"
RUN groupadd -r chrome \
    && wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - \
    && echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list \
    && apt-get -qqy update \
    && apt-get -qqy --no-install-recommends install ${CHROME_VERSION:-google-chrome-stable}

# Cleanup
RUN rm /etc/apt/sources.list.d/google.list \
    && rm -rf /var/lib/apt/lists/* /var/cache/apt/*

# Install htmltox
COPY ./ /go/src/github.com/mkenney/go-chrome
