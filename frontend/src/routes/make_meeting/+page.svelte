<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	$: studentIdInput = '';
	$: startAtInput = '';
	$: endAtInput = '';
	$: submittingForm = false;
	$: submissionError = '';
	$: loginError = '';

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		loginError = params.get('err') ?? '';
	});

	const onSubmit = async (e: any) => {
		if (submittingForm) return;

		submittingForm = true;
		submissionError = '';

		e.preventDefault();

		const requestBody = {
			studentId: studentIdInput,
			startAtInput: startAtInput,
			endAtInput: endAtInput
		};
		const reqPromise = fetch(`${PUBLIC_API_HOST}/meetings`, {
			method: 'POST',
			body: JSON.stringify(requestBody)
		});

		const response = await reqPromise;
		if (!response || !response?.ok) {
			console.error(`failed to sign up/login`);
			return;
		}

		window.localStorage.setItem('SessionToken', response.headers.get('X-Session-Token')!);
		// window.location.href = '/';
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
					type="date"
					bind:value={startAtInput}
				/>
			</div>
			<div class="w-full">
				<label for="end_at" class="block text-gray-400">End at</label>
				<input
					class="bg-gray-900 px-4 py-2 w-full rounded-md"
					id="end_at"
					type="date"
					bind:value={endAtInput}
				/>
			</div>
			<input
				class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50 mt-3"
				type="submit"
				value="Schedule"
				disabled={submittingForm}
			/>
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>
