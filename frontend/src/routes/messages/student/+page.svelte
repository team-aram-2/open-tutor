<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	import messagesData from '$lib/mock/messages_mock.json';

	import Sendbutton from '$lib/components/messaging/sendbutton.svelte';
	import Attachimagebutton from '$lib/components/messaging/attachimagebutton.svelte';

	import type { MessageItem } from '$lib/types/types';

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
		<!-- conversationId={message.conversationId}
  messageId={message.messageId}
  messageAttachments={message.messageAttachments}
  sentOn={message.sentOn} -->
		<Message originId={message.originId} messageContent={message.messageContent} userId="userId"
		></Message>
	{/each}
	<!-- Sort messages by timestamp -->
	<!-- Load attachments? -->
</div>

<div class="textboxcontainer">
	<Attachimagebutton />
	<textarea class="textbox"> </textarea>
	<Sendbutton />
</div>

<style>
	.messagecontainer {
		width: 100%;
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;
		flex: 1 1 auto;
		overflow-y: scroll;
		height: 80%;
		left: 0;
		top: 0;
		z-index: 0;
	}

	.textboxcontainer {
		display: flex;
		flex-wrap: nowrap;
		flex-direction: row;
		flex: 1 1 auto;
		align-items: center;
		position: relative;
		bottom: 0;
		right: 0;

		height: 20%;

		width: 100%;
		z-index: 100;
		background-color: var(--yellow-neutral);
	}

	.textbox {
		flex: 1 1 auto;
		position: relative;

		height: calc(100% - 40px);
		box-sizing: border-box;
		resize: none;

		margin: 20px 20px 20px 20px;

		z-index: 10;

		font-size: var(--font-size);
		font-family: 'Inter', sans-serif;
		background-color: var(--yellow-very-light);
		border-radius: 10px;
		border-color: transparent;
	}
</style>
