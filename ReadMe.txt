A simple chatroom made with Go RPC.
Multiple clients can send messages to one server, and the server keeps all chat history.

How it works

server.go → runs the RPC server and saves all messages.

client.go → connects to the server, sends messages, and shows the full chat history.

How to run

Open 3 terminals in the same folder.

Run the server:

go run server.go


Run the client:

go run client.go


Type messages and press Enter to send.
Type exit to quit.

Link to google drive:
https://drive.google.com/file/d/1GwJjyuith8fbVP_1KckCinxW7xaS7InQ/view?usp=drive_link