Ploio is an automated deployment pipeline tool for kubernetes services. In concept it is based on Netflix's spinnaker, but is more specialized (limited, focussed, trimmed down, crappier) in that it requires use of kubernetes, helm, and istio.


### Components

* __Cluster__ - k8s cluster to talk to (dev, stage, prod-va, prod-ut, etc)
* __Application__ - Represents an application
* __Pipeline__ - A set of steps/actions to take on a service when triggered to do so (e.g. CI server triggers it via an API call with the chart name and version)
* __Step__ - an action taken in a service pipeline (e.g. deploy to cluster, manual check, kanary deploy, green/blue deploy, run integration tests)
* __Template__ - a pre-defined pipeline config that can be reused on any service


### Step Types

* __Standard Deploy__ - Deploy a helm chart version to a cluster (configurable). Also allows to configure the overwrites for the values in a chart. For example, I add a deploy step to a my service that deploys to dev, and sets the number of instances to 1 via a values override. But production the instances is set to 3.
* __Manual Approve__ - wait for someone to manually allow it to proceed to next step
* __Kanary deploy (with a K)__ - Deploy chart to a cluster, and route X% of traffic to it on that cluster
* __Manual Test Deploy__ - deploy chart to a cluster, and route all traffic with a specific header or metadata to that service. For example, deploy a pull request to prod, but only route traffic to the new version with http header of "X-Tester=true"