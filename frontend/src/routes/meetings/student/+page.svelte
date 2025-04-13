<script lang="ts">
	import MeetingCard from '$lib/components/meetings/meeting-card.svelte';
	import AddNewButton from '$lib/components/meetings/add-new-button.svelte';
	import { onMount } from 'svelte';

	import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';
	import RatingSubmission from '$lib/components/cards/rating-submission.svelte';

	// $: current_id = $user_id;
	$: meetings = [];
	// let meetingId = '';
	let isInitialized = false;
	$: if ($user_id && !isInitialized) {
		isInitialized = true;
		// loadData($user_id);
	}

	// async function loadData(userId: string) {
	// 	await fetchMeetings(userId);
	// }

	// const sendMessage = async () => {
	// 	if (messageContent.trim()) {
	// 		try {
	// 			const res = await fetch(PUBLIC_API_HOST + '/message', {
	// 				method: 'POST',
	// 				headers: {
	// 					'Content-Type': 'application/json'
	// 				},
	// 				body: JSON.stringify({
	// 					originId: current_id,
	// 					conversationId: conversationId,
	// 					message: messageContent
	// 				})
	// 			});
	// 			const response = await res.json();
	// 			console.log(response);
	// 			if (res.status === 201) {
	// 				console.log('Message sent');
	// 				fetchMessages();
	// 				messageContent = '';
	// 			} else {
	// 				console.log('Failed to send message');
	// 			}
	// 		} catch (err) {
	// 			console.log('Error sending message:', err);
	// 		}
	// 	}
	// };

	const fetchMeetings = async () => {
		try {
			const res = await fetch(PUBLIC_API_HOST + '/meetings', {
				method: 'GET',
				credentials: 'include'
			});
			meetings = await res.json();
		} catch (err) {
			console.error('Error in the process of fetching meetings:', err);
		}
	};

	onMount(async () => {
		fetchMeetings();
	});
</script>

<div class="meetings-container relative flex flex-col">
	<!-- !!! DO NOT PUT ANYTHING ABOVE THIS INSIDE THIS DIV IF YOU DO YOU WILL DIE !!! -->
	<!-- rating submission popover has to be the top element for layout to work -->
	<div class="absolute flex justify-center items-center w-full h-full bg-black/80 z-50">
		<RatingSubmission />
	</div>

	{#each meetings as meeting}
		<MeetingCard {meeting} />
	{/each}

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
