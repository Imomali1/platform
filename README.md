# platform
That platform of university!
***

## gRPC-based architecture
**1. External Clients:** 
> These are the clients that want to access your gRPC-based services. They might be mobile apps, web applications, or other systems that don't natively support gRPC.

**2. gRPC Gateway (Client):**
>The gRPC Gateway acts as a bridge between the external clients and the gRPC-based services. It provides an HTTP/JSON API that external clients can communicate with. The gRPC Gateway translates RESTful HTTP/JSON requests from external clients into gRPC requests that the server understands.

**3. gRPC-based Services (Server):**
> These are the actual gRPC services that perform the core functionality of your application. They are typically designed to be efficient and use gRPC for communication.

**4. Communication Flow:**
>* External clients make HTTP/JSON requests to the gRPC Gateway.
>* The gRPC Gateway translates these requests into gRPC calls and forwards them to the gRPC-based services.
>* The gRPC services process the requests and generate responses.
>* The gRPC Gateway translates the responses back into HTTP/JSON and sends them to the external clients.