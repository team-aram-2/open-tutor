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

export interface SkillsItem{
	id: string;
	category: string;
	title: string;
	description: string;
	questions: string[];
	tutorHasSkill: boolean;
}