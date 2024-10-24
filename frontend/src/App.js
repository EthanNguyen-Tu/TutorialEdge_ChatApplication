import React, { Component } from "react";
import "./App.css";
import { connect, sendMsg } from "./api/index.js";
import Header from "./components/Header/Header";

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
        <Header />
        <button onClick={this.send}>Press Me!</button>
      </div>
    );
  }
}

export default App;
