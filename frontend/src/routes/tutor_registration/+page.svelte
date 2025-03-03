<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	$: submittingForm = false;
	$: submissionError = '';

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		submissionError = params.get('err') ?? '';

		// TODO: NEED STATE CHECK HERE FOR LOGGED_IN...
		// This should only be available to users who have logged in and are not already registered as a tutor.
		// Redirect to tutor dashboard if this page is reached by a valid tutor.
		// Redir to /login if the user isn't logged in.
	});
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-1 w-120 mx-auto">
		<h1 class="text-center text-5xl">Tutor Registration</h1>
		<form class="flex flex-col gap-3" method="POST" action="{PUBLIC_API_HOST}/tutor}">
			<input
				class="bg-gray-900 px-4 py-2 rounded-md"
				name="first_name"
				type="text"
				placeholder="First name"
			/>
			<input
				class="bg-gray-900 px-4 py-2 rounded-md"
				name="last_name"
				type="text"
				placeholder="Last name"
			/>
			<input
				class="bg-gray-900 px-4 py-2 rounded-md"
				name="email"
				type="email"
				placeholder="Email"
			/>
			<input
				class="bg-gray-900 px-4 py-2 rounded-md"
				name="password"
				type="password"
				placeholder="Password"
			/>
			<input
				class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50"
				type="submit"
				value="Register"
				disabled={submittingForm}
			/>
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>
