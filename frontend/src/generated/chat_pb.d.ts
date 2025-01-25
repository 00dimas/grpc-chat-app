import * as jspb from 'google-protobuf'



export class Message extends jspb.Message {
  getId(): string;
  setId(value: string): Message;

  getUserId(): string;
  setUserId(value: string): Message;

  getGroupId(): string;
  setGroupId(value: string): Message;

  getText(): string;
  setText(value: string): Message;

  getTimestamp(): string;
  setTimestamp(value: string): Message;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Message.AsObject;
  static toObject(includeInstance: boolean, msg: Message): Message.AsObject;
  static serializeBinaryToWriter(message: Message, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Message;
  static deserializeBinaryFromReader(message: Message, reader: jspb.BinaryReader): Message;
}

export namespace Message {
  export type AsObject = {
    id: string,
    userId: string,
    groupId: string,
    text: string,
    timestamp: string,
  }
}

export class MessageRequest extends jspb.Message {
  getUserId(): string;
  setUserId(value: string): MessageRequest;

  getGroupId(): string;
  setGroupId(value: string): MessageRequest;

  getText(): string;
  setText(value: string): MessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MessageRequest): MessageRequest.AsObject;
  static serializeBinaryToWriter(message: MessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageRequest;
  static deserializeBinaryFromReader(message: MessageRequest, reader: jspb.BinaryReader): MessageRequest;
}

export namespace MessageRequest {
  export type AsObject = {
    userId: string,
    groupId: string,
    text: string,
  }
}

export class MessageResponse extends jspb.Message {
  getStatus(): string;
  setStatus(value: string): MessageResponse;

  getMessageId(): string;
  setMessageId(value: string): MessageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: MessageResponse): MessageResponse.AsObject;
  static serializeBinaryToWriter(message: MessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageResponse;
  static deserializeBinaryFromReader(message: MessageResponse, reader: jspb.BinaryReader): MessageResponse;
}

export namespace MessageResponse {
  export type AsObject = {
    status: string,
    messageId: string,
  }
}

export class MessageHistoryRequest extends jspb.Message {
  getGroupId(): string;
  setGroupId(value: string): MessageHistoryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MessageHistoryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: MessageHistoryRequest): MessageHistoryRequest.AsObject;
  static serializeBinaryToWriter(message: MessageHistoryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MessageHistoryRequest;
  static deserializeBinaryFromReader(message: MessageHistoryRequest, reader: jspb.BinaryReader): MessageHistoryRequest;
}

export namespace MessageHistoryRequest {
  export type AsObject = {
    groupId: string,
  }
}

