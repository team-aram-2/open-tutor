<script lang="ts">
	import FontSize from '$lib/components/settings_page/font-size.svelte';
	import { PUBLIC_API_HOST } from '$env/static/public';

	$: isTutor = true;
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
</style>
