# Tutorial Notes

## Running the Project

In the front end folder, `npm run start` opens up the webpage locally at [http://localhost:3000/](http://localhost:3000/). In the back end folder, `go run main.go` starts up the back end chat app.

## [Part 1 - Initial Setup](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-1-initial-setup/)

- `GO111MODULE=on` did not work, but this seems to achieve the same thing:

  ```[C++]
  go env -w GO111MODULE=on
  ```

## [Part 2 - Simple Communication](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-2-simple-communication/)

### WebSocket

#### WebSocket Protocol

- offers duplex communication from a non-trusted source to a server across a TCP socket connection
  - maintains a single TCP socket connection for sending and listening to messages on

|                                       Pros                                        |                                                                          Cons                                                                          |
| :-------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------------------------------------------------------------: |
|              Prevents continually polling the web server for updates              |                             Introduction of state makes scaling up the application across multiple instances more complex                              |
|             Prevents having to perform TCP handshakes for every poll              | Need to consider options such as storing state in message brokers, or in databases/memory caches that can scale in parallel with application instances |
| Reduces the amount of network overhead required for the any real-time application |                                                                                                                                                        |
|      Allows maintaining large amounts of clients on a single server instance      |                                                                                                                                                        |

## [Part 3 - Designing our Frontend](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-3-designing-our-frontend/)

Need to install `node-sass`, since ReactJS projects do not automatically come with the ability to handle sass files.
Note: Node Sass is no longer supported. Please use `sass` or `sass-embedded` instead.

```[shell]
yarn add node-sass
```

## [Part 4 - Handling Multiple Clients](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/)

- In order to relative call other files in the same project, the full import path based on the module name should be used.
  - Relative imports like `./<module>` are not supported in Go.
  - Code should be organized in packages.
  - The module name can be found in the "go.mod" file after the `module` keyword.

### How Multiple Clients are Handled

Whenever a new connection is made, it is added to a pool of existing connections where every time a message is sent, all connections in the pool receives the message.

## [Part 5 - Improving the Frontend](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-5-improved-frontend/)

- in SCSS, "&" is used for nesting, allowing the combination of multiple selectors into a single line of code

### Constructor

- where the state for a component should be set up
- is not necessary for class components to recieve props, so a component does not need to have one if nothing needs to be done in the constructor

### State

- accessed in a component with `this.state`
- `this.setState({ key: value })` - function that accepts an object with key-value pairs
  - CASE: given object key matches an existing key in the state => updates the value in the state to the new value matching the key that is provided in the given object
  - CASE: given object key does not already exist in the state => creates a new key in the state with the new value matching the key that is provided in the given object

### Misc.

- `.bind()` - JavaScript method that allows an object to borrow a method from another object
  - Use Case: Maintaining the correct this context in callbacks

## [Part 6 - Dockerizing your Backend](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-6-dockerizing-your-backend/)

Building a Docker image:

```shell
docker build -t backend .
```

- `-t` - allows the specification of a name and optionally a tag in the `name:tag` format
- `.` - specifies the path to the Dockerfile is in the current directory

Running the Docker image:

```shell
docker run -it -p 8080:8080 backend
```

- `-it` - combines the `-i/--interactive` and `-t/--tty` flags, allowing the container to be run interactively with a terminal session
  - `-i/--interactive` - keeps the standard input (STDIN) open, allowing the interaction with the container
  - `-t/--tty` - allocates a pseudo-TTY, which gives a terminal interface
- `-p 8080:8080` - tells Dockerto map ports between the host machine's 8080 port and the container's 8080 port
  - exposes port 8080 inside the container to port 8080 on the host machine, allowing access to services running on the port inside the container from the host machine
- `backend` - the name or tag ofd the docker image to run

### Adding Docker Benefits

- allows the deployment of the docker-based application to any server or platform that supports Docker
- allows the specification of the exact environment that the application needs in order to start up

## Terminology

1. [duplex communication][1] - a system where communication can occur in both directions between two devices or parties
2. [channels][2] - pipes that link between goroutines within Go based applications that allow communication and subsequent passing of values to and from variables
   - used to make high performance, high concurrent applications in Go with minimal challenge compared to other programming languages
3. `constructor` - object-oriented programming concept/method that creates an object
   - called automatically every time an object is created, allowing the object to set initial values for its attributes or perform other setup tasks
4. `props` - pieces of data passed into a child component from the parent
   - ex. `<SomeComponent prop="Some prop data" />`
5. `state` - data controlled within a component

## References

1. [What is duplex communication?][1]
2. [golang-standards Project Layout](https://github.com/golang-standards/project-layout)
3. [Managing module source](https://go.dev/doc/modules/managing-source)
4. [Go Channels Tutorial][2]
5. [this.state - How to Use State in React js](https://www.iamtimsmith.com/blog/this-state-how-to-use-state-in-react)
6. [docker build (legacy builder)](https://docs.docker.com/reference/cli/docker/build-legacy/)

[1]: https://www.pubnub.com/learn/glossary/duplex-communication/
[2]: https://tutorialedge.net/golang/go-channels-tutorial/
