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

### [Part 3 - Designing our Frontend](https://tutorialedge.net/projects/chat-system-in-go-and-react/part-3-designing-our-frontend/)

Need to install `node-sass`, since ReactJS projects do not automatically come with the ability to handle sass files.
Note: Node Sass is no longer supported. Please use `sass` or `sass-embedded` instead.

```shell
yarn add node-sass
```

## Terminology

1. [duplex communication][1] - a system where communication can occur in both directions between two devices or parties

## References

1. [What is duplex communication?][1]

[1]: https://www.pubnub.com/learn/glossary/duplex-communication/
