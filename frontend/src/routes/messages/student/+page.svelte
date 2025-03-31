<script lang="ts">
	import Message from '$lib/components/messaging/message.svelte';
	import messagesData from '$lib/mock/messages_mock.json';

	import Sendbutton from '$lib/components/messaging/sendbutton.svelte';
	import Attachimagebutton from '$lib/components/messaging/attachimagebutton.svelte';

	import type { MessageItem } from '$lib/types/types';
	import { onMount } from 'svelte';
	import autosize from 'autosize';
	import { font_size } from '$lib/stores';
	import { get } from 'svelte/store';

	let messages: MessageItem[] = [];
	onMount(() => {
		// Set up event listener to update the size of the textboxcontainer
		let textbox_container = document.getElementsByClassName('textboxcontainer')[0] as HTMLElement;

		// Set up autosize for the textbox
		let text_area: HTMLTextAreaElement | null =
			document.querySelector<HTMLTextAreaElement>('textarea');
		if (text_area) {
			autosize(text_area);

			text_area.addEventListener('autosize:resized', function () {
				if (textbox_container) {
					const style = window.getComputedStyle(text_area);
					const lineHeight = parseFloat(style.lineHeight);
					// Here we assume that if the scrollHeight is roughly equal to the line-height,
					// then there's only one line of text. Adjust the tolerance if needed.
					if (text_area.scrollHeight <= lineHeight + 2) {
						text_area.style.height = lineHeight + 'px';
					}
					textbox_container.style.height = text_area.offsetHeight + 40 + 'px';
				}
			});
			if (textbox_container) {
				textbox_container.style.height =
					String(window.getComputedStyle(document.body).getPropertyValue('--font-size') + 40) +
					'px';
			}
		}
		if (typeof document !== undefined) {
			(document.getElementsByClassName('textboxcontainer')[0] as HTMLElement).style.height =
				Number(Number(get(font_size).slice(0, -2)) * 1.5 + 40) + 'px';
		}
	});

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
</script>

<div class="containercontainer">
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
		<textarea rows="1" class="textbox"></textarea>
		<Sendbutton />
	</div>
</div>

<style>
	.containercontainer {
		flex: 1;
		max-height: 100%;
		min-width: 0;
		display: flex;
		flex-direction: column;
	}
	.messagecontainer {
		flex: 1;
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;

		min-width: 0;

		overflow-y: scroll;
		z-index: 0;
	}

	.textboxcontainer {
		display: flex;
		flex-wrap: nowrap;
		flex-direction: row;
		align-items: stretch;

		min-width: 0;
		min-height: 0;

		position: relative;
		bottom: 0;
		right: 0;

		z-index: 100;
		background-color: var(--yellow-neutral);
	}

	.textbox {
		flex: 1 1 auto;
		position: relative;

		max-height: calc(11.5 * var(--font-size));
		height: min-content;
		min-height: calc(1.5 * var(--font-size));
		width: auto;
		min-width: 0;

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
