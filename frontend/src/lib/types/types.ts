export interface MessageItem {
	conversationId: string;
	messageContent: string;
	messageId: string;
	originId: string;
	userId: string;
	sentOn: number;
	messageAttachments: string[];
}

export interface MeetingItem {
	id: string;
	tutorId: string;
	studentId: string;
	startAt: string;
	endAt: string;
	zoomJoinLink: string;
	zoomHostLink: string;
}
