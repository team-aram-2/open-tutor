<script lang="ts">
	// export let startAt = 0;
	// export let endAt = 0;
	export let tutorName = 'Jon Bovi';
	export let zoomLink = 'https://www.youtube.com/watch?v=dQw4w9WgXcQ';

	import { sidebar_width } from '$lib/stores';
</script>

<div class="meeting-card">
	<div class="time-date-info">
		<div class="month">
			<p>September</p>
		</div>
		<div class="day">
			<p>29</p>
		</div>
		<div class="time">
			<p>12:30 - 12:00</p>
		</div>
	</div>
	<div class="tutor-info-block">
		<div class="tutor-name">{tutorName}</div>
		<div></div>
	</div>
	<div class="link-info" style="width: {sidebar_width}px;">
		<a class="zoom-link" href={zoomLink}>Zoom Link</a>
	</div>
</div>

<style>
	@import './meeting-card.css';
</style>
