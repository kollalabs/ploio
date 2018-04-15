## Ploio

Safe, reliable, and automated deployments for kubernetes.

### Concepts

**Safe:** Ploio was inspired by Michael Bryzek's presentation [Testing in Production](https://www.infoq.com/presentations/testing-software-production). To deploy updates safely, testing is critical. To supplement your unit tests, Ploio allows you to deploy canary tests and smoke tests right into production to give you extreme confidence that your upates are safe. 

**Reliable:** Ploio by default abstracts Kubernetes manifest YAML away from the developer, but still keeps the developer heavily involved in making sure their application is healthy.

**Automated:** Ploio was built to be run by your CI/CD system and provides some extra automation that raw kubernetes manifests do not. For instance, Ploio will detect if an update hits a crash loop backoff and will roll back automatically. Ploio does NOT replace your CI system, it augments it. Your CI system will still do unit tests, build the application, and upload your newly built container to a container repo. Ploio can then be triggered once that is complete to manage your deployment strategies.






