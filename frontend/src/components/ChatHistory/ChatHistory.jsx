import Message from "../Message";
import "./ChatHistory.scss";
import React, { Component } from "react";

class ChatHistory extends Component {
  render() {
    console.log(this.props.chatHistory);

    // STEP: Map each message in the chat history to a Message component
    const messages = this.props.chatHistory.map((msg) => (
      <Message message={msg.data} />
    ));

    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>
        {messages}
      </div>
    );
  }
}

export default ChatHistory;
