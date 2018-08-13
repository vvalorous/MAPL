# MAPL (Manageable Access-control Policy Language)
MAPL is a new language for access control rules, designed for a microservices environment.
It is designed to be intuitive, rich and expressive, as well as simple and straightforward.
MAPL makes it easier to declare and maintain access control rules. The language enables fine-grained control of traffic, with a resource based control model that takes into account the principals, action, resources on the principals, and conditions on traffic attributes, similar to AWS’s IAM policy model.

The MAPL rules have the following syntax:

`<sender, receiver, resource, operation> : <conditions> : <decision>`

where, basically, a rule gives a decision of wheteher the sender (client) may do the operaion on the resource of the receiver (server) when the conditions apply.
The language is described thoroughly in [MAPL](docs/MAPL.md)

# MAPL engine

Given a list of rules and message attributes, the MAPL engine gives a decision whether to allow, allow and alert or block the communication.
The engine is documented in [MAPL Engine](docs/MAPL_ENGINE.md)
The MAPL engine can be used in service meshes, API gateways and IAM solutions.
A demonstration of the use of the MAPL engine for service-to-service authorization in Istio using a gRPC adapter for Istio’s mixer service can be found in [insert link]

# Status of the MAPL Project
This project is still under active development.
If you have any questions about MAPL or how to use the MAPL engine, please contact MAPL@octarinesec.com

# Engine Roadmap
- Currently the engine supports service to service communication. We need to support ingress and egress by IP or CIDR
- Support Istio's source.labels and destination.labels
- Suppert HTTP request headers
- Add patterns to the wildcard use (*,?) similar to fnmatch [https://www.gnu.org/software/libc/manual/html_node/Wildcard-Matching.html]? Or allow input of regex in rules?
- Add message attributes to be used in the conditions
- Allow for definitions of service groups for reducing the number of rules (for example, use source.labels and destination.labels in the conditions)