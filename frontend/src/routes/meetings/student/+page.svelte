<script lang="ts">
	import MeetingCard from '$lib/components/meetings/meeting-card.svelte';
	import AddNewButton from '$lib/components/meetings/add-new-button.svelte';
	import { onMount } from 'svelte';
	//import { get } from 'svelte/store';
	import { sidebar_width } from '$lib/stores';

	//import type { MeetingItem } from '$lib/types/types';
	//import { PUBLIC_API_HOST } from '$env/static/public';
	import { user_id } from '$lib/stores';

	//$: current_id = $user_id;
	//let meetingId = '';
	//let meetings: string[] = [];
	let isInitialized = false;
	$: if ($user_id && !isInitialized) {
		isInitialized = true;
		// loadData($user_id);
	}
	/*async function loadData(userId: string) {
		await fetchMeetings(userId);
	}*/

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
	/*const fetchMeetings = async (userId: string) => {
		try {
			const res = await fetch(PUBLIC_API_HOST + '/meetings/' + userId);
			meetings = await res.json();
			// console.log(meetings);
			meetingId = meetings[0];
		} catch (err) {
			console.error('Error in the process of fetching meetings:', err);
		}
	};*/
	// const handleKeydown = (event: KeyboardEvent) => {
	// 	if (event.key === 'Enter' && messageContent.trim()) {
	// 		event.preventDefault();
	// 		sendMessage();
	// 	}
	// };
	onMount(async () => {});
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
		max-width: 100%;
		overflow-x: scroll;
		border: 2px solid chartreuse;
	}
</style>
