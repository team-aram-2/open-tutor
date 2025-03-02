<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	//import messagesData from '$lib/mock/messages_mock.json';
	import { onMount } from 'svelte';

	import type { MessageItem } from '$lib/types/types';
	import { PUBLIC_API_HOST } from '$env/static/public';

	let messagesData;
	let messages: MessageItem[] = [];
	let messageContent = '';
	const fetchMessages = async () => {
		try {
			messages = [];
			const res = await fetch(
				PUBLIC_API_HOST + '/conversation/messages/5425b439-4e2b-4a2b-8808-76c8fb406a13'
			);
			messagesData = await res.json();
			console.log(messagesData);
			for (let item of messagesData.messages) {
				let newItem: MessageItem = {
					conversationId: item.conversationId,
					messageId: item.id,
					originId: item.originId,
					messageContent: item.message,
					messageAttachments: item.messageAttachments,
					userId: item.originId,
					sentOn: new Date(item.sentOn).getTime()
				};
				messages = [...messages, newItem];
			}
			// Sort messages by timestamp
			messages.sort((a, b) => {
				return a.sentOn - b.sentOn;
			});
			console.log(messages);
		} catch (err) {
			console.log('fetch failed.');
		}
	};

	const sendMessage = async () => {
		if (messageContent.trim()) {
			try {
				const res = await fetch(PUBLIC_API_HOST + '/message', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						originId: '50df85b9-cb07-45f5-a7e1-95b40cb5b2fb',
						conversationId: '5425b439-4e2b-4a2b-8808-76c8fb406a13',
						message: messageContent
					})
				});
				const response = await res.json();
				console.log(response);
				if (res.status === 201) {
					console.log('Message sent');
					fetchMessages();
					messageContent = '';
				} else {
					console.log('Failed to send message');
				}
			} catch (err) {
				console.log('Error sending message:', err);
			}
		}
	};
	const handleKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Enter' && messageContent.trim()) {
			event.preventDefault();
			sendMessage();
		}
	};
	onMount(async () => {
		fetchMessages(); // Now fetchMessages is called on component mount
	});
</script>

<div class="messagecontainer">
	<div style="padding-bottom: 30px"></div>
	{#each messages as message}
		<!-- conversationId={message.conversationId}
  messageId={message.messageId}
  messageAttachments={message.messageAttachments}
  sentOn={message.sentOn} -->
		<Message
			originId={message.originId}
			messageContent={message.messageContent}
			userId="50df85b9-cb07-45f5-a7e1-95b40cb5b2fb"
		></Message>
	{/each}
	<!-- Sort messages by timestamp -->
	<!-- Load attachments? -->
</div>

<div class="textboxcontainer">
	<textarea class="textbox" bind:value={messageContent} on:keydown={handleKeydown}></textarea>
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
		left: 0;
		top: 0;
		z-index: 0;
	}
	.textboxcontainer {
		position: fixed;
		bottom: 0;
		right: 0;
		width: calc(100% - 375px);
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
