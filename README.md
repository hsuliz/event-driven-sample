# event-driven-sample

Sample of event-driven architecture using Go, Kafka and MongoDB.

# Description

There are 3 microservices:

- `engine` calculates the determinant of the matrix
- `history` stores the result of the calculation in MongoDB
- `order` creates a random matrix

Communication between them is done by Kafka topics.

# Flow

Generally it's unidirectional communication where `order` creates matrix and sends it to `order-topic`. Topic is
consumed by `engine`, which calculates determinant and sends it to `history` by `history-topic` to save result in DB.