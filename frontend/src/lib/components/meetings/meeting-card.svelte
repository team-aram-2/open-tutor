<script lang="ts">
	// Exports for container properties
	export let tutorName = 'Jon Bovi';
	export let zoomLink = 'https://www.youtube.com/watch?v=dQw4w9WgXcQ';

	import { sidebar_width } from '$lib/stores';
</script>

<article class="meeting-card" aria-label="Upcoming meeting with {tutorName}">
	<div class="time-date-info">
		<div class="month"><p>September</p></div>
		<div class="day"><p>29</p></div>
		<div class="time"><p>12:30am - 12:00pm</p></div>
	</div>
	<div class="tutor-info-block">
		<div class="tutor-name">{tutorName}</div>
		<div></div>
	</div>
	<div class="link-info" style="width: {sidebar_width}px;">
		<a class="zoom-link" href={zoomLink} aria-label={'Join Zoom meeting with ' + tutorName}>
			Zoom Link
		</a>
	</div>
</article>

<style>
	@import './meeting-card.css';
</style>
