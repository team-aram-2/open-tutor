<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { redirect } from '@sveltejs/kit';
	import { onMount } from 'svelte';

	$: studentIdInput = '';
	$: startAtInput = '';
	$: endAtInput = '';
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

		const requestBody = {
			studentId: studentIdInput,
			startAt: new Date(startAtInput).toISOString(),
			endAt: new Date(endAtInput).toISOString()
		};
		const reqPromise = fetch(`${PUBLIC_API_HOST}/meeting`, {
			method: 'POST',
			body: JSON.stringify(requestBody),
			credentials: 'include'
		});

		const response = await reqPromise;
		if (!response || !response?.ok) {
			console.error(`failed to sign up/login`);
			return;
		}
		if (response.ok) {
			redirect(302, '/meetings');
		}

		window.localStorage.setItem('SessionToken', response.headers.get('X-Session-Token')!);
	};
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-4 w-120 mx-auto">
		<h1 class="text-center text-5xl">Schedule meeting with student</h1>

		<span class="text-red-600 text-center">{loginError}</span>

		<form class="flex flex-col gap-3" on:submit={onSubmit}>
			<div class="w-full">
				<label for="student_id" class="block text-gray-400">Student ID</label>
				<input
					class="bg-gray-900 px-4 py-2 w-full rounded-md"
					id="student_id"
					type="text"
					placeholder=""
					bind:value={studentIdInput}
				/>
			</div>
			<div class="w-full">
				<label for="start_at" class="block text-gray-400">Start at</label>
				<input
					class="bg-gray-900 px-4 py-2 w-full rounded-md"
					id="start_at"
					type="datetime-local"
					bind:value={startAtInput}
				/>
			</div>
			<div class="w-full">
				<label for="end_at" class="block text-gray-400">End at</label>
				<input
					class="bg-gray-900 px-4 py-2 w-full rounded-md"
					id="end_at"
					type="datetime-local"
					bind:value={endAtInput}
				/>
			</div>
			<input
				class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50 mt-3"
				type="submit"
				value="Schedule"
				disabled={false}
			/>
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>
