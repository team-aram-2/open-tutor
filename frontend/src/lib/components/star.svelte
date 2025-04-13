<script lang="ts">
	import { createEventDispatcher } from 'svelte';

	export let filled = false;

	const dispatch = createEventDispatcher();
</script>

<button class="w-8 h-8 cursor-pointer" on:click={() => dispatch('click')}>
	<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100">
		<!-- CSS Variables -->
		<style>
			.star-filled {
				fill: var(--yellow-neutral, #ffd700);
				stroke: none;
			}

			.star-outlined {
				fill: none;
				stroke: var(--yellow-neutral, #ffd700);
				stroke-width: 2;
				/* Ensure the outline is visible */
				stroke-opacity: 1;
			}
		</style>

		{#if filled}
			<path
				class="star-filled"
				d="M50,10 L61.8,35.4 L89.4,39.1 L69.7,58.3 L74.4,85.8 L50,72.8 L25.6,85.8 L30.3,58.3 L10.6,39.1 L38.2,35.4 Z"
			/>
		{:else}
			<path
				d="M50,10 L61.8,35.4 L89.4,39.1 L69.7,58.3 L74.4,85.8 L50,72.8 L25.6,85.8 L30.3,58.3 L10.6,39.1 L38.2,35.4 Z"
				class="star-outlined"
			/>
		{/if}
	</svg>
</button>
