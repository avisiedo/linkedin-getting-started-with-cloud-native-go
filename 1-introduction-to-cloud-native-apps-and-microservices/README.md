# Introduction to Cloud Native Apps and Microservices

## Motivation for cloud native apps

- They are dynamic
- They won't brake
- They scale as they need.

Antifragility: This means cloud native application won't break anymore.

Continuous delivery and devops: Deliver new features and new applications.

Opex savings: automation and utilization. You don't waste your money into
computing units that you realy don't need.

## Challenges and Design Principles

- Challenges that come with Cloud native application development.

### Design Principles of Cloud Native Applications

- Design for performance: Responsive, concurrency, efficiency.
- Design for automation: Automate dev tasks and ops tasks.
- Design for resiliency: Fault tolerant and self healing.
- Design for elasticity: Dynamically sacale up and down and be reactive.
- Design for delivery: Short round trips and automated delivery.
- Design for diagnosability: Cluster-wide logs, traces, and metrics.

## Decomposition with Microservices

> Design > Build > Run

- Design: Design components
  - Complexity unit
  - Data integrity unit
  - Feature unit
  - Decoupled unit
- Build: Dev components
  - Planning unit
  - Knowledge unit
  - Development unit
  - Integration unit
- Run: Ops components
  - Release unit
  - Deployment unit
  - Runtime unit
  - Scaling unit

### The anatomy of an Ops Component

- Contraints (mostly technology driven):
  - Application does not use kernel space APIs.
  - Application does not listen on random ports.
  - Application does not require an exotic operatin system.
  - Used endpoints (DNS name, IPs, Ports) can be configured.

### Microservice Decomposition Trade-Offs

- Dev Components: System > Subsystem > Components > Services
- Ops Components: Monolith > Macroservices > Microservices > Nanoservices
- Decomposition trade-offs
  - +1 points:
    - More flexible to scale
    - Runtime isolation 
    - Independent releases and deployments
    - Higher resource utilization
  - -1 points:
    - Distribution debt: Latency
    - Increased infrastructure complexity
    - Increased integration complexity
    - Increased troubleshooting complexity

### Cloud Native Stack

Cloud Native App > Applications > Containers > Resources

    Go Lang      > Kubernetes   > Kubernetes >  crio, docker, podman

- Viewed Cloud native applications promise superior reliability and
  arbitrary scalability.
- Covered the important design principles for the development of
  Cloud native applications.
- Discussed certain decomposition trade-offs when using microservices.
- Explored the Cloud native stack addresses the inherent challenges and
  complexities.
- Viewed Go which is a suitable Cloud Native application platform.
