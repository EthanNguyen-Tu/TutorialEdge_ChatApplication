# Tutorial Notes

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

|                                          Pros                                           |                                                                                Cons                                                                                |
| :-------------------------------------------------------------------------------------: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------: |
|              Prevents continually polling the web server for <br/> updates              |                                Introduction of state makes scaling up the application <br/> across multiple instances more complex                                 |
|             Prevents having to perform TCP handshakes <br/> for every poll              | Need to consider options such as storing state in <br/> message brokers, or in databases/memory caches <br/> that can scale in parallel with application instances |
| Reduces the amount of network overhead required <br/> for the any real-time application |                                                                                                                                                                    |
|      Allows maintaining large amounts of clients <br/> on a single server instance      |                                                                                                                                                                    |

## Terminology

1. [duplex communication][1] - a system where communication can occur in both directions between two devices or parties

## References

1. [What is duplex communication?][1]

[1]: https://www.pubnub.com/learn/glossary/duplex-communication/
