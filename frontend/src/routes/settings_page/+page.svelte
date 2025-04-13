<script lang="ts">
	import FontSize from '$lib/components/settings_page/font-size.svelte';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { logged_in } from '$lib/stores';
	import { goto } from '$app/navigation';

	$: isTutor = true;

	// Function to clear the cookie and update the store
	function logout(): void {
		// Remove the cookie by setting an expiration date in the past.
		// Adjust the cookie name and path as needed.
		document.cookie = 'session_token=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/;';

		// Update the Svelte store to indicate the user is logged out.
		logged_in.set(false);

		// Optionally, you might also refresh the page or navigate to a login route.
		// For example:
		goto('/login');
	}
</script>

<div class="setting-container">
	<div class="setting-categories"></div>
	<div class="setting-body">
		<!-- Font size -->
		<div class="setting-entry min-h-[60px]">
			<div class="setting-label">Font Size</div>
			<FontSize></FontSize>
		</div>
		<div class="setting-entry">
			<div class="setting-label">Billing portal</div>
			<a
				href="{PUBLIC_API_HOST}/billing_portal"
				class="mr-4 px-4 py-1 bg-sky-400 rounded-md shadow-md cursor-pointer">Open billing portal</a
			>
		</div>

		<!-- Tutor settings -->
		{#if isTutor}
			<span class="mt-4 ml-4 text-lg text-white">Tutor Settings</span>
			<div class="setting-entry">
				<div class="setting-label">Verify identity</div>
				<a
					href="{PUBLIC_API_HOST}/tutor_id_verification"
					class="mr-4 px-4 py-1 bg-sky-400 rounded-md shadow-md cursor-pointer">Start</a
				>
			</div>
		{/if}
		{#if $logged_in}
			<div class="setting-entry">
				<button class="log-out" on:click={logout}>Log out</button>
			</div>
		{/if}
	</div>
</div>

<style>
	.setting-container {
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;

		height: 100%;
		width: 100%;
	}
	.setting-categories {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;

		height: 7em;

		border-top-left-radius: 25px;
		border-top-right-radius: 25px;

		background-color: var(--purple-very-dark);
	}
	.setting-body {
		display: flex;
		flex-direction: column;

		height: 100%;
	}

	.setting-entry {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;

		align-items: center;

		width: 100%;
		height: calc(var(--font-size) * 2);

		border-bottom: 5px solid var(--purple-very-dark);
		border-left: 5px solid var(--purple-very-dark);
		border-radius: 15px;

		background-color: var(--purple-dark);
	}
	.setting-label {
		flex: 1;

		text-align: left;

		margin-left: 3%;
		padding: 0;

		font-size: var(--font-size);
		font-weight: bolder;
		color: var(--yellow-very-light);
	}
	.log-out {
		width: fit-content;
		height: fit-content;
		font-size: calc(0.8 * var(--font-size));
		padding: 5px 15px 5px 15px;
		background-color: var(--yellow-light);
		border-radius: 15px;
	}
</style>
