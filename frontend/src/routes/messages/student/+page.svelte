<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	//import messagesData from '$lib/mock/messages_mock.json';
	import { onMount } from 'svelte';

	import type { MessageItem } from '$lib/types/types';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';

	$: current_id = $user_id;
	let messagesData;
	let messages: MessageItem[] = [];
	let messageContent = '';
	let conversationId = '';
	let conversations: string[] = [];
	let isInitialized = false;
	$: if ($user_id && !isInitialized) {
		isInitialized = true;
		loadData($user_id);
	}
	async function loadData(userId: string) {
		await fetchConversations(userId);
		if (conversationId) {
			fetchMessages();
		}
	}
	const fetchMessages = async () => {
		try {
			//messages = [];
			let tempMessages: MessageItem[] = [];
			const res = await fetch(PUBLIC_API_HOST + '/conversation/messages/' + conversationId, {
				credentials: 'include'
			});
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
				tempMessages = [...tempMessages, newItem];
			}
			// Sort messages by timestamp
			tempMessages.sort((a, b) => {
				return a.sentOn - b.sentOn;
			});
			if (messages.length != tempMessages.length) {
				messages = tempMessages;
			}
			console.log(messages);
		} catch (err) {
			console.log('fetch failed.', err);
		}
	};

	const sendMessage = async () => {
		if (messageContent.trim()) {
			try {
				const res = await fetch(PUBLIC_API_HOST + '/message', {
					method: 'POST',
					body: JSON.stringify({
						originId: current_id,
						conversationId: conversationId,
						message: messageContent
					}),
					credentials: 'include'
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
	const fetchConversations = async (userId: string) => {
		try {
			const res = await fetch(PUBLIC_API_HOST + '/conversation/user/' + userId, {
				credentials: 'include'
			});
			conversations = await res.json();
			console.log(conversations);
			conversationId = conversations[0];
		} catch (err) {
			console.log('Error in the process of fetching messages:', err);
		}
	};
	const handleKeydown = (event: KeyboardEvent) => {
		if (event.key === 'Enter' && messageContent.trim()) {
			event.preventDefault();
			sendMessage();
		}
	};
	onMount(() => {
		const interval = setInterval(() => {
			fetchMessages();
		}, 1000);
		return () => {
			clearInterval(interval);
		};
	});
</script>

<div class="messagecontainer">
	<div style="padding-bottom: 30px"></div>
	{#each messages as message}
		<!-- conversationId={message.conversationId}
  messageId={message.messageId}
  messageAttachments={message.messageAttachments}
  sentOn={message.sentOn} -->
		<Message originId={message.originId} messageContent={message.messageContent} userId={current_id}
		></Message>
	{/each}
	<!-- Sort messages by timestamp -->
	<!-- Load attachments? -->
</div>

<div class="textboxcontainer">
	<textarea class="textbox" bind:value={messageContent} on:keydown={handleKeydown}></textarea>
	<button class="send-button" on:click={sendMessage} disabled={!messageContent.trim()}>Send</button>
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
		display: flex;
		align-items: center;
	}
	.textbox {
		position: relative;
		width: calc(100% - 80px);
		height: calc(100% - 20px);
		/* padding: 20px 20px 20px 20px; */
		margin: 10px 10px 10px 20px;
		box-sizing: border-box;
		resize: none;

		z-index: 101;

		font-size: xx-large;
		font-family: 'Inter', sans-serif;
		background-color: var(--yellow-very-light);
		border-radius: 10px;
		border-color: transparent;
	}
	.send-button {
		height: calc(100% - 20px);
		width: 60px;
		margin: 10px 20px 10px 0;
		border-radius: 10px;
		border: none;
		background-color: var(--yellow-very-light);
		font-family: 'Inter', sans-serif;
		font-size: large;
		cursor: pointer;
		transition: background-color 0.2s;
	}

	.send-button:hover {
		background-color: #e0e0e0;
	}

	.send-button:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}
</style>
