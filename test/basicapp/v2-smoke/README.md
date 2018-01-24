## Smoke Deployment Test

This folder contains all the changes that would be made when deploying v2 of our basicapp as a smoke test

We have a duplicate of the original deployment with only these changes:

1. Changed the version label to reflect v2
2. Changed the docker image to be the updated code base
3. Added istio default route for all normal traffic to go to v1 
4. Added istio route for all traffic with a certain header to v2

_NOTE: to make this work for services that may be several services deep, the header would have to propagate between services. It would be interesting to see if we could tie into standardized tracing headers from opentracing or the other one Derek loves so much from google. Then propagated headers/meta-data would already be instrumented._
