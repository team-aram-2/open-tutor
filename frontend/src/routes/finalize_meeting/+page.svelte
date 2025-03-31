<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	$: meetingIdInput = '';
	$: submissionError = '';
	$: loginError = '';
	$: submittingForm = false;

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		loginError = params.get('err') ?? '';
	});

	const onSubmit = async (e: SubmitEvent) => {
		if (submittingForm) return;

		submittingForm = true;
		submissionError = '';

		e.preventDefault();
		const reqPromise = fetch(`${PUBLIC_API_HOST}/meeting/${meetingIdInput}/finalize`, {
			method: 'POST',
			credentials: 'include'
		});

		submittingForm = false;

		const response = await reqPromise;
		if (!response || !response?.ok) {
			console.error(`failed to finalize meeting`);
			return;
		}
	};
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-4 w-120 mx-auto">
		<h1 class="text-center text-5xl">Finalize meeting with student</h1>

		<span class="text-red-600 text-center">{loginError}</span>

		<form class="flex flex-col gap-3" on:submit={onSubmit}>
			<div class="w-full">
				<label for="meeting_id" class="block text-gray-400">Meeting ID</label>
				<input
					class="bg-gray-900 px-4 py-2 w-full rounded-md"
					id="meeting_id"
					type="text"
					placeholder=""
					bind:value={meetingIdInput}
				/>
			</div>
			<input
				class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50 mt-3"
				type="submit"
				value="Finalize meeting"
				disabled={false}
			/>
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>
