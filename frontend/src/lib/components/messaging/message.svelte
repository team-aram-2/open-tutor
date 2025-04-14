<script lang="ts">
	export const conversationId = -1;
	export const messageId = -1;
	export let originId;
	export let messageContent;
	export const messageAttachments = -1;
	export let userId;
	export let sentOn = -1;
	import '../../../app.css';
	const isOwnMessage = userId === originId;

	function formatDate(unixTimestamp: number) {
		console.log(unixTimestamp);
		if (unixTimestamp === -1) return '';
		const date = new Date(unixTimestamp);
		return date.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
	}
</script>

<div class="message-container">
	<div class="message" class:own-message={isOwnMessage} class:other-message={!isOwnMessage}>
		{messageContent}
	</div>
	<div class="timestamp" class:own-timestamp={isOwnMessage} class:other-timestamp={!isOwnMessage}>
		{originId + ' / ' + formatDate(sentOn)}
	</div>
</div>

<style>
	.message-container {
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		width: 100%;
	}

	.timestamp {
		color: white;
		font-size: calc(var(--font-size) * 0.7);
		margin-bottom: 5px;
		opacity: 0.8;
	}

	.own-timestamp {
		align-self: flex-end;
		margin-right: 30px;
	}

	.other-timestamp {
		align-self: flex-start;
		margin-left: 30px;
	}

	.message {
		border-radius: 20px;
		padding: 20px 30px 20px 30px;
		margin: 40px 30px 0px 30px;
		font-size: calc(var(--font-size) * 0.825);
		max-width: 70%;
		display: inline-block;
		word-wrap: break-word;
	}

	.own-message {
		background-color: var(--yellow-light);
		margin-left: auto;
		margin-right: 20px;
		align-self: flex-end;
	}

	.other-message {
		background-color: var(--purple-very-light);
		margin-right: auto;
		margin-left: 20px;
		align-self: flex-start;
	}
</style>
