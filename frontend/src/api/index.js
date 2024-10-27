var socket = new WebSocket("ws://localhost:8080/ws");

// Connect to the WebSocket endpoint and listen for events
// Logs connection events, printing out any issues to the
// browser console
let connect = (cb) => {
  console.log("Attempting Connection...");

  socket.onopen = () => {
    console.log("Successfully Connected.");
  };

  socket.onmessage = (msg) => {
    console.log(msg);
    cb(msg);
  };

  socket.onclose = (event) => {
    console.log("Socket Closed Connection: ", event);
  };

  socket.onerror = (error) => {
    console.log("Socket Error: ", error);
  };
};

// Method to send messages to the backend from the
// frontend via the WebSocket connection while logging
// the message
let sendMsg = (msg) => {
  console.log("Sending Message: ", msg);
  socket.send(msg);
};

export { connect, sendMsg };
