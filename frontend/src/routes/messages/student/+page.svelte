<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	import messagesData from '$lib/mock/messages_mock.json';

	import type { MessageItem } from '$lib/types/types';
	console.log(messagesData);
	let messages: MessageItem[] = [];
	for (let item of messagesData.messages) {
		let newItem: MessageItem = {
			conversationId: item.conversationId,
			messageId: item.id,
			originId: item.originId,
			messageContent: item.message,
			messageAttachments: item.messageAttachments,
			userId: 'userId',
			sentOn: new Date(item.sentOn).getTime()
		};
		messages.push(newItem);
	}

	// Sort messages by timestamp
	messages.sort((a, b) => {
		return a.sentOn - b.sentOn;
	});
	console.log(messages);
</script>

<div class="messagecontainer">
	<div style="padding-bottom: 30px"></div>
	{#each messages as message}
		<Message
			conversationId={message.conversationId}
			messageId={message.messageId}
			originId={message.originId}
			messageContent={message.messageContent}
			messageAttachments={message.messageAttachments}
			userId="userId"
			sentOn={message.sentOn}
		></Message>
	{/each}
	<!-- Sort messages by timestamp -->
	<!-- Load attachments? -->
</div>

<div class="textboxcontainer">
	<textarea class="textbox"> </textarea>
</div>

<style>
	.messagecontainer {
		width: auto;
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;
		flex: 1 1 auto;
		overflow-y: scroll;
		height: calc(100vh - 100px);
		z-index: 0;
	}
	.textboxcontainer {
		position: relative;
		bottom: 0;
		right: 0;
		height: 100px;
		min-width: max-content;
		z-index: 100;
		background-color: var(--yellow-neutral);
	}
	.textbox {
		position: relative;
		width: calc(100% - 40px);
		height: calc(100% - 20px);
		/* padding: 20px 20px 20px 20px; */
		margin: 10px 20px 10px 20px;
		box-sizing: border-box;
		resize: none;

		z-index: 101;

		font-size: xx-large;
		font-family: 'Inter', sans-serif;
		background-color: var(--yellow-very-light);
		border-radius: 10px;
		border-color: transparent;
	}
</style>
