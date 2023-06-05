import React, { useEffect } from "react";

const Room = (props) => {
  useEffect(() => {
    if (props.match && props.match.params) {
      const ws = new WebSocket(
        `ws://localhost:8000/join?roomID=${props.match.params.roomID}`
      );

      ws.addEventListener("open", () => {
        console.log("SENDING....")
        ws.send(JSON.stringify({ join: "true" }));
      });

      ws.addEventListener("message",(e) => {
        console.log(e.data)
      })

    };
});

  return (
    <div>
      <video autoPlay controls={true}></video>
      <video autoPlay controls={true}></video>
    </div>
  );
};

export default Room;
