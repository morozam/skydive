
# Skydive

Skydive is an open source real-time network topology and protocols analyzer.
It aims to provide a comprehensive way of understanding what is happening in
the network infrastructure.

Skydive agents collect topology information and flows and forward them to a
central agent for further analysis. All the information is stored in an
Elasticsearch database.

Skydive is SDN-agnostic but provides SDN drivers in order to enhance the
topology and flows information.

![](https://github.com/skydive-project/skydive.network/raw/images/overview.gif)

## Key features

* Captures network topology and flows
* Full history of network topology and flows
* Distributed
* Ability to follow a flow along a path in the topology
* Supports VMs and Containers infrastructure
* Unified query language for topology and flows (Gremlin)
* Web and command line interfaces
* REST API
* Easy to deploy (standalone executable)
* Connectors to OpenStack, Docker, OpenContrail, Kubernetes

## Quick start

### All-in-one

The easiest way to get started is to download the latest binary and to run it using the `all-in-one` mode :

```console
curl -Lo - https://github.com/skydive-project/skydive-binaries/raw/jenkins-builds/skydive-latest.gz | gzip -d > skydive && chmod +x skydive && sudo mv skydive /usr/local/bin/

SKYDIVE_ETCD_DATA_DIR=/tmp SKYDIVE_ANALYZER_LISTEN=0.0.0.0:8082 sudo -E /usr/local/bin/skydive allinone
```

Open a browser to http://localhost:8082 to access the analyzer Web UI.

```

Open a browser to http://localhost:8082 to access the analyzer Web UI.

### Docker

```console
docker run -d --privileged --pid=host --net=host -p 8082:8082 -p 8081:8081 \
    -e SKYDIVE_ANALYZER_LISTEN=0.0.0.0:8082 \
    -v /var/run/docker.sock:/var/run/docker.sock -v /run/netns:/var/run/netns \
    skydive/skydive allinone
```

Open a browser to http://localhost:8082 to access the analyzer Web UI.



Open a browser to http://localhost:8082 to access the analyzer Web UI.



## License

This software is licensed under the Apache License, Version 2.0 (the
"License"); you may not use this software except in compliance with the
License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
