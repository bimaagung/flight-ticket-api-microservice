## Flight Ticket Microservice
Mini project of making an airplane ticket sales application using the microservice concept with event-driven communication between services.

<br>

## 🔍 Entity Relationship Diagram

<div align="center">
  
  <img src="https://raw.githubusercontent.com/bimaagung/flight-ticket-api-microservice/master/erdflightticket.png" width="600">

</div>

## 🔍 Event-Driven Architecture

<div align="center">
  
  <img src="https://raw.githubusercontent.com/bimaagung/flight-ticket-api-microservice/master/microservice.png" width="600">

</div>
<br>
This project is designed using a microservices architecture, which is an approach to software development where a single application is built as a collection of smaller services, each running independently and communicating with each other through APIs or messaging systems. This architecture allows for loose coupling between the services, making them more flexible, scalable, and easier to maintain.

The project consists of six services, each serving a specific function:

### Track Service
The `Track Service` is responsible for tracking flights, managing flight schedules, and providing information on flight status. It is built using the Gin framework, which is a lightweight web framework for Go, and uses a PostgreSQL database to store and manage flight information.

### Airplane Service
The `Airplane Service` is responsible for managing airplanes, including information on airplane models, maintenance schedules, and availability. It is also built using the Gin framework and uses a MySQL database to store and manage airplane information.

### Ticket Service
The `Ticket Service` is responsible for handling ticket sales, booking, and cancellation. It is also built using the Gin framework and uses a PostgreSQL database to manage ticket-related data.

### Order Service
The `Order Service` is responsible for managing orders, including order placement, payment processing, and fulfillment. It is built using the NestJS framework, which is a progressive Node.js framework for building efficient, reliable, and scalable server-side applications, and uses a MongoDB database to manage order-related data.

### Payment Service
The `Payment Service` is responsible for handling payment processing and managing payment information. It is also built using the NestJS framework and uses a MongoDB database to manage payment-related data.

### Authentication Service
The `Authentication Service` is responsible for handling user authentication, including user registration, login, and logout. It is built using the NestJS framework and uses a PostgreSQL database to manage user-related data.

All of these services communicate with each other using an event-driven architecture, which is facilitated by the RabbitMQ message broker. This allows for loose coupling between the services and enables them to operate independently. When an event occurs in one service, such as the creation of an order, the relevant information is published to the message broker, which can then be consumed by other services that need that information.

To deploy these services, Docker containers can be used. Docker is a popular containerization platform that allows developers to package an application and all of its dependencies into a single container that can run on any operating system or infrastructure. Using Docker containers, each service can be built and deployed separately, and can be easily scaled up or down depending on demand.

For orchestration and management of these containers, Kubernetes can be used. Kubernetes is an open-source container orchestration platform that automates the deployment, scaling, and management of containerized applications. With Kubernetes, developers can easily deploy and manage containers across multiple hosts, scale containers up or down based on resource demands, and manage load balancing and networking between containers.

Overall, this architecture provides scalability, fault tolerance, and flexibility, making it a suitable approach for complex projects with many independent components. By using Docker containers and Kubernetes, developers can easily deploy and manage these services, making it easier to scale the application as the business needs grow.

<br>

## 💻 Technology Stack

- GO with framework Gin
- Typescript with framework Nest.JS
- PostgreSQL
- MySQL
- Elasticsearch
- Redis
- TypeORM
- RabbitMQ
- Viper
- Docker
