<script lang="ts">
	export const conversationId = -1;
	export const messageId = -1;
	export let originId;
	export let messageContent;
	export const messageAttachments = -1;
	export let userId;
	export const sentOn = -1;

	import '../../../app.css';

	const isOwnMessage = userId === originId;
</script>

<div class="message" class:own-message={isOwnMessage} class:other-message={!isOwnMessage}>
	{messageContent}
</div>

<style>
	.message {
		border-radius: 20px;
		padding: 20px 30px 20px 30px;
		margin: 0px 30px 40px 30px;
		font-size: calc(var(--font-size) * 0.825);
		max-width: 70%;
		display: inline-block;
		word-wrap: break-word;
	}
	.own-message {
		background-color: var(--yellow-light);
		margin-left: auto;
		margin-right: 20px;
	}
	.other-message {
		background-color: var(--purple-very-light);
		margin-right: auto;
		margin-left: 20px;
	}
</style>
