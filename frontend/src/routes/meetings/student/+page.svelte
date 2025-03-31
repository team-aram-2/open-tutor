<script lang="ts">
	import MeetingCard from '$lib/components/meetings/meeting-card.svelte';
	import AddNewButton from '$lib/components/meetings/add-new-button.svelte';
	import { onMount } from 'svelte';
	// import { get } from 'svelte/store';

	// import type { MeetingItem } from '$lib/types/types';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';

	// let _meetingId = '';
	let meetings: string[] = [];
	let isInitialized = false;
	$: if ($user_id && !isInitialized) {
		isInitialized = true;
	}

	const fetchMeetings = async () => {
		try {
			const res = await fetch(PUBLIC_API_HOST + '/meetings', {
				method: 'GET',
				credentials: 'include'
			});
			meetings = await res.json();
			console.log(meetings);
			// _meetingId = meetings[0];
		} catch (err) {
			console.error('Error in the process of fetching meetings:', err);
		}
	};

	onMount(async () => {
		fetchMeetings();
	});
</script>

<div class="meetings-container flex flex-col">
	<MeetingCard />

	<div class="fit-content fixed right-10 bottom-10">
		<a href="/make_meeting">
			<AddNewButton />
		</a>
	</div>
</div>

<style>
	.meetings-container {
		height: 100%;
		overflow-x: scroll;
		flex-grow: 1;
	}
</style>
