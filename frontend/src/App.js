import "./App.css";
import ChatHistory from "./components/ChatHistory/ChatHistory.jsx";
import { connect, sendMsg } from "./api/index.js";
import Header from "./components/Header/Header";
import React, { Component } from "react";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      chatHistory: [],
    };
  }

  componentDidMount() {
    connect((msg) => {
      console.log("New Message!");
      this.setState((prevState) => ({
        chatHistory: [...this.state.chatHistory, msg],
      }));
      console.log(this.state);
    });
  }

  send() {
    console.log("Button has been pressed!");
    sendMsg("The User has pressed the button.");
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Press Me!</button>
      </div>
    );
  }
}

export default App;
