<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	import { onMount } from 'svelte';

	import Sendbutton from '$lib/components/messaging/sendbutton.svelte';
	import Attachimagebutton from '$lib/components/messaging/attachimagebutton.svelte';

	import type { MessageItem } from '$lib/types/types';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';

	interface Conversation {
		id: string;
		name: string;
	}

	$: current_id = $user_id;
	let messagesData;
	let messages: MessageItem[] = [];
	let messageContent = '';
	let conversationId = '';
	let conversations: Conversation[] = [];
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
			if (messages != tempMessages) {
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
			conversationId = conversations[0].id;
		} catch (err) {
			console.log('Error in the process of fetching messages:', err);
		}
	};

	const handleConversationChange = (event: Event) => {
		const selectElement = event.target as HTMLSelectElement;
		conversationId = selectElement.value;
		fetchMessages();
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

<div class="conversation-selector">
	<select on:change={handleConversationChange} value={conversationId} class="conversation-dropdown">
		{#each conversations as conversation}
			<option value={conversation.id}>{conversation.name}</option>
		{/each}
	</select>
</div>

<div class="messagecontainer">
	<div style="padding-bottom: 30px"></div>
	{#each messages as message}
		<Message originId={message.originId} messageContent={message.messageContent} userId={current_id}
		></Message>
	{/each}
</div>

<div class="textboxcontainer">
	<Attachimagebutton />
	<textarea class="textbox" bind:value={messageContent} on:keydown={handleKeydown}></textarea>
	<Sendbutton />
	<button class="send-button" on:click={sendMessage} disabled={!messageContent.trim()}>
		Send
	</button>
</div>

<style>
	.conversation-selector {
		width: 100%;
		padding: 10px;
		background-color: var(--yellow-neutral);
	}

	.conversation-dropdown {
		width: 100%;
		padding: 10px;
		border-radius: 5px;
		border: 1px solid #ccc;
		font-family: 'Inter', sans-serif;
		font-size: var(--font-size);
	}

	.messagecontainer {
		width: 100%;
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;
		flex: 1 1 auto;
		overflow-y: scroll;
		height: 70%; /* Reduced to make room for dropdown */
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
		display: flex;
		align-items: center;
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
