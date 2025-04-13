<script lang="ts">
	import { sidebar_width } from '$lib/stores';

	export let meeting;
	export let onSubmitRating = () => {};

	const startAtDate = new Date(meeting.startAt);
	const endAtDate = new Date(meeting.endAt);
</script>

<div class="meeting-card">
	<div class="time-date-info">
		<div class="month">
			<p>{startAtDate.toLocaleString('default', { month: 'long' })}</p>
		</div>
		<div class="day">
			<p>{startAtDate.getDate()}</p>
		</div>
		<div class="time">
			<p>
				{startAtDate.toLocaleTimeString('en-US', { hour: '2-digit', minute: '2-digit' })} - {endAtDate.toLocaleTimeString(
					'en-US',
					{ hour: '2-digit', minute: '2-digit' }
				)}
			</p>
		</div>
	</div>
	<div class="tutor-info-block">
		<div class="tutor-name">{meeting.userName}</div>
		<div></div>
	</div>
	<div class="flex flex-col gap-1 my-auto">
		<div class="link-info" style="width: {sidebar_width}px;">
			<a class="zoom-link" href={meeting.zoomHostLink ?? meeting.zoomJoinLink}>Zoom Link</a>
		</div>
		{#if !meeting.zoomHostLink}
			<div class="link-info" style="width: {sidebar_width}px;">
				<button class="zoom-link" on:click={onSubmitRating}>Submit Rating</button>
			</div>
		{/if}
	</div>
</div>

<style>
	@import './meeting-card.css';
</style>
