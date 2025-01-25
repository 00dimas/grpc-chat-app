import React, { useState, useEffect } from 'react';
import { grpc } from '@improbable-eng/grpc-web'; // Import library gRPC-Web
import { ChatServiceClient } from '../generated/ChatServiceClientPb.ts';
import { MessageRequest, MessageResponse } from '../generated/chat_pb.js'; // Import message types

const Chat: React.FC = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [newMessage, setNewMessage] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);

  // Fungsi untuk mengirim pesan
  const sendMessage = () => {
    const request = new MessageRequest();
    request.setUserId('user1'); // Ganti dengan user ID dinamis jika perlu
    request.setGroupId('group1'); // Ganti dengan group ID dinamis jika perlu
    request.setText(newMessage);

    grpc.invoke(ChatServiceClient.SendMessage, {
      request,
      host: 'http://localhost:50051', // URL gRPC server
      onMessage: (response: MessageResponse) => {
        console.log('Message sent:', response.toObject());
        setMessages((prevMessages) => [...prevMessages, newMessage]); // Tambahkan pesan baru ke state
        setNewMessage(''); // Reset input pesan
      },
      onEnd: (code: grpc.Code, msg: string) => {
        if (code !== grpc.Code.OK) {
          console.error('Error sending message:', msg);
        }
      },
    });
  };

  // Fungsi untuk mengambil pesan (polling)
  const fetchMessages = () => {
    const request = new MessageRequest();
    request.setGroupId('group1');

    grpc.invoke(ChatServiceClient.GetMessages, {
      request,
      host: 'http://localhost:50051', // URL gRPC server
      onMessage: (response: MessageResponse) => {
        const text = response.getText();
        if (text) {
          setMessages((prevMessages) => {
            const newMessages = [...prevMessages, text];
            return Array.from(new Set(newMessages)); // Hindari duplikasi
          });
        }
      },
      onEnd: (code: grpc.Code, msg: string) => {
        if (code !== grpc.Code.OK) {
          console.error('Error fetching messages:', msg);
        }
      },
    });
  };

  useEffect(() => {
    setLoading(true);

    // Set interval untuk polling pesan
    const intervalId = setInterval(() => {
      fetchMessages();
    }, 5000); // Polling setiap 5 detik

    // Cleanup untuk menghentikan polling
    return () => clearInterval(intervalId);
  }, []);

  return (
    <div>
      <h1>Chat</h1>
      {loading ? (
        <p>Loading...</p>
      ) : (
        <div>
          {messages.map((msg, index) => (
            <div key={index}>{msg}</div>
          ))}
        </div>
      )}
      <input
        type="text"
        value={newMessage}
        onChange={(e) => setNewMessage(e.target.value)}
        placeholder="Type your message here..."
      />
      <button onClick={sendMessage} disabled={!newMessage.trim()}>
        Send
      </button>
    </div>
  );
};

export default Chat;
