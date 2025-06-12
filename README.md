# event-driven-sample

Sample of event-driven architecture using Go, Kafka and MongoDB.

# Description

There are 3 microservices:

- `engine` calculates the determinant of the matrix
- `order` creates a random matrix and waits for calculation

Communication between them is done by Kafka topics.

# Flow

Generally both microservices talk to each other by Kafka topics.