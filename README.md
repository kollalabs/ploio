## Ploio

Ploio (PLEE-oh) is an automated deployment pipeline tool for kubernetes

### Concepts

Ploio's mission is to make your application deployments safe, reliable, and fast. 

Ploio borrows many concepts from Spinnaker, the amazing deployment tool from Netflix. However, Spinnaker was built in the age of VMs and a huge portion of its functionality is dedictated to provisioning VMs and other AWS utilities. With modern cloud-native applications on kubernetes, the environment you are deploying to has already been provisioned (just spin up more workers!), so Ploio is focussed on just the deployment of your container-based apps to kubernetes.

Ploio makes use of Istio and Linkerd to do things such as Canary deployments, Smoke Deployments, and latency/error checks to validate successful deployments. 

Ploio abstracts the Kubernetes YAML away from the developer, but keeps the developer heavily involved in making sure their application is healthy.

Ploio does NOT replace your CI system, it augments it. Your CI system will still do unit tests, build the application, and upload your newly built container to a container repo. Ploio can then be triggered once that is complete to manage your deployment strategies.


### Components

* __Cluster__ - k8s cluster to talk to (dev, stage, prod-va, prod-ut, etc)
* __Application__ - Represents an application, the unit that can be deployed
* __Pipeline__ - A set of steps/actions to take on a service when triggered to do so (e.g. CI server triggers it via an API call with the chart name and version)
* __Step__ - an action taken in a service pipeline (e.g. deploy to cluster, manual check, kanary deploy, green/blue deploy, run integration tests)
* __Template__ - a pre-defined pipeline config that can be reused on any service


### Step Types

* __Deployment__ - Deploy your application to a cluster. Strategies available are Smoke, Canary, Blue/Green, Rolling
* __Check__ - Check your deployment to make sure all is well. Strategies are: Manual or Metrics (latency, error rates)
* __Rollback/Teardown__ - Rollback to a previous version, or teardown a current smoke/canary deployment



