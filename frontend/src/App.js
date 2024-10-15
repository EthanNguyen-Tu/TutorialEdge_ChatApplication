import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api/index.js";

class App extends Component {
  constructor(props) {
    super(props);
    connect();
  }

  send() {
    console.log("Button hass been pressed!");
    sendMsg("The User has pressed the button.");
  }

  render() {
    return (
      <div className="App">
        <button onClick={this.send}>Press Me!</button>
      </div>
    );
  }
}

export default App;
