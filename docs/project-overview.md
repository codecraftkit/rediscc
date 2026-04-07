# Project Overview — rediscc

## What is it

rediscc is a software tool that allows applications to connect and communicate with in-memory data storage systems in a straightforward way. It acts as an intermediary that simplifies the most common operations: storing data, retrieving it, deleting it, and sending notifications between different parts of a system.

## What is it for

It solves the problem of having to write repetitive and complex code every time an application needs to interact with an in-memory data store. Its main functionalities are:

- **Simplified connection**: Configure the connection to the data store with a single instruction instead of multiple steps.
- **Basic operations**: Store, read, delete, and search data directly.
- **Messaging**: Send messages between different parts of the system in real time.
- **Debug mode**: Enable detailed logs to diagnose problems without modifying the code.

## How it works

1. The application using rediscc establishes a connection by specifying the data store address and the database number to use.
2. The system verifies that the connection is successful before proceeding.
3. Once connected, the application can perform operations:
   - **Store**: saves a piece of data with a name (key) and optionally an expiration time.
   - **Retrieve**: recovers a previously stored piece of data using its name.
   - **Delete**: removes a specific piece of data from the store.
   - **Search**: finds all data whose name matches a search pattern.
   - **Notify**: sends a message to a channel so that other subscribed systems can receive it.

## Who uses it

- **Application developers**: They integrate rediscc into their projects to handle temporary data, caches, or message queues without dealing with the complexity of direct connection to the store.
- **Backend teams**: They use the library to maintain consistency in how multiple services access the same data store.

## Security

- The connection to the data store is made through a secure address that can include authentication credentials.
- The system validates that the connection is successful before allowing any operation, preventing silent errors.
- Debug data is only displayed when explicitly enabled, preventing accidental exposure of sensitive information in production environments.

## Integration with other systems

rediscc is designed to integrate within any application that needs temporary in-memory data storage. It connects to compatible storage servers and exposes a standard contract that allows applications to change the underlying implementation without modifying the rest of the code. It also supports real-time messaging, allowing different services to communicate with each other through notification channels.
