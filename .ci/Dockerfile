#
#   Copyright 2021-2023 Red Hat, Inc.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
FROM golang:1.18-alpine AS builder

# Install dependencies
RUN apk add --no-cache git bash curl zip

# Install yq
RUN curl -sL -O https://github.com/mikefarah/yq/releases/download/v4.9.5/yq_linux_amd64 -o /usr/local/bin/yq && mv ./yq_linux_amd64 /usr/local/bin/yq && chmod +x /usr/local/bin/yq
RUN yq

# Copy the registry build tools
COPY build-tools /build-tools
# Copy the index generator source
COPY index/ /index
COPY tests/registry /registry

# Download all the offline parent devfiles
RUN bash /build-tools/dl_parent_devfiles.sh

# Download offline starter projects
RUN bash /build-tools/dl_starter_projects.sh go-starter community

# Run the registry build tools
RUN bash /build-tools/build.sh /registry /build

FROM devfile-index-base

COPY --from=builder /build /registry