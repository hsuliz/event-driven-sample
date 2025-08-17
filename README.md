# event-driven-sample

Sample of event-driven architecture using Go, Kafka and MongoDB.

# Abstract

There are 2 microservices:

- `engine` calculates the determinant of the matrix
- `order` creates a random matrix and waits for calculation

Communication between them is done by Kafka topics.

# How to run
### Skaffold
1. Skaffold
    ```shell
    skaffold dev
    ```
### Local development

1. Setup infrastructure
    1. Compose up
        ```shell
        docker compose up -d
        ```
    2. Init Kafka topics
       ```shell
       docker-compose exec kafka kafka-topics --create --topic order-topic --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092
       docker-compose exec kafka kafka-topics --create --topic engine-topic --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server localhost:9092
       ```

       > For testing if everything OK, this commands can be used
       >
       > ```shell
        > docker-compose exec kafka kafka-console-producer --topic order-topic --bootstrap-server localhost:9092
        > docker-compose exec kafka kafka-console-consumer --topic order-topic --from-beginning --bootstrap-server localhost:9092

2. Download dependencies
    ```shell
    go mod download
    ```
3. Build engine service
    ```shell
    go build -o engine ./cmd/engine
    ```
4. Build order service
    ```shell
    go build -o order ./cmd/order
    ```
5. Run each one