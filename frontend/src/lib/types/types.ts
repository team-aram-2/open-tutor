export interface MessageItem {
  conversationId: string;
  messageContent: string;
  messageId: string;
  originId: string;
  userId: string;
  sentOn: number;
  messageAttachments: string[];
}