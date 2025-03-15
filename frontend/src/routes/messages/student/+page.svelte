<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	import messagesData from '$lib/mock/messages_mock.json';

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
	// END MESSAGE LOAD

	//
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
	<div class="messagesbutton mt-auto mr-0 mb-[20px] ml-[10px]">
		<!-- margin: auto 0px 20px 10px; -->

		<svg
			width="100%"
			height="100%"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
			style="margin: calc(var(--font-size) * .1)"
		>
			<!-- Paperclip Path -->
			<path
				d="M21.2 11.05
				l-9.19 9.19
				a4.7 4.7 0 0 1 -7 -7
				l9.19 -9.19
				a3.2 3.2 0 1 1 4.53 4.53
				l-9.19 9.19
				a1.5 1.5 0 0 1 -2.1 -2.1
				l8.5 -8.5"
				stroke="var(--yellow-dark)"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
	</div>
	<textarea class="textbox"> </textarea>
	<div class="messagesbutton mt-[20px] mr-[10px] ml-0px mb-[20px]">
		<svg
			width="100%"
			height="100%"
			viewBox="0 0 24 24"
			fill="none"
			xmlns="http://www.w3.org/2000/svg"
		>
			<!-- Vertical line -->
			<line
				x1="12"
				y1="19"
				x2="12"
				y2="5"
				stroke="var(--yellow-dark)"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
			<!-- Head of arrow -->
			<polyline
				points="5 12 12 5 19 12"
				fill="none"
				stroke="var(--yellow-dark)"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			/>
		</svg>
	</div>
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

	.messagesbutton {
		display: flex;
		align-items: center;
		justify-items: center;

		width: calc(2 * var(--font-size));
		height: calc(100% - 40px);

		border-radius: var(--font-size);

		background-color: var(--yellow-light);
	}
	:hover.messagesbutton {
		transition: background-color 210ms ease-in-out;

		background-color: var(--purple-neutral);
	}
	.messagesbutton:hover svg :is(path, line, polyline) {
		transition: stroke var(--purple-light) 200ms ease-in-out;

		stroke: var(--purple-very-light);
	}
</style>
