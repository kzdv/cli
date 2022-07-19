# ZDV CLI

This CLI provides a wrapper for kubectl that integrates with our OIDC authentication system.

## Usage

### Container environment

You can use the following to run the cli in a containerized environment. This assumes using docker, but 
any container runtime should work.

```bash
docker run -it -p 12297:12297 -v /tmp/config:/root/.config denartcc/cli:latest bash
```

You can change the latest tag to a specific version if needed. They will be formatted as `v1.0.0`. The port 
forwarding is used for OAuth2 Authentication.  The volume mounting is optional, but if you need to share 
yaml files into or from the environment this will be useful.

The container also sets the following aliases:

* k = "zdv kubectl"
* kg = "zdv kubectl get"
* kgp = "zdv kubectl get pods"
* kgn = "zdv kubectl get nodes"
* kd = "zdv kubectl describe"
* krm = "zdv kubectl delete"
* ka = "zdv kubectl apply"

### Install

The wrapper will likely cause a conflict with other kubectl installations in its current form. It is recommended to use 
the container image.  If you do wish to install the cli locally, you can do so by running the following:

```bash
curl -sfL https://raw.githubusercontent.com/kzdv/cli/main/install.sh | bash -
```

### Login

To login, run `zdv login`.  This will open a web browser to the OIDC provider to initiate login.

### Refresh Login

Kubectl will automatically fetch a new Identity Token if it expires. However, the Refresh Token is only 
valid for 7 days. If this expires as well, you will need to login again.

### Kubectl

Access to the cluster is automatically handled by assigned roles in the cluster. If your roles change after 
you login, you will need to login again to get the new roles.

We automatically symlink kubectl to the `zdv` command.

## Scripts

scripts/test -- Runs unit tests
scripts/build -- Builds the cli binary packages

## License

Copyright 2022 Daniel Hawton, and the VATUSA KZDV Web Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.