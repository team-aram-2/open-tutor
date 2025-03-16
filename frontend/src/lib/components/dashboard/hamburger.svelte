<script>
	import { createEventDispatcher } from 'svelte';

	// Export props for external customization.
	export let open = true; // Whether the menu is open (X) or closed (hamburger)
	export const backgroundColor = '#FFF';
	export let lineColor = '#000'; // Line color
	export let label = 'Toggle menu';

	const dispatch = createEventDispatcher();

	function toggle() {
		open = !open;
		dispatch('toggle', { open });
	}
</script>

<button
	class="hamburger {open ? ' ' : 'open'}"
	on:click={toggle}
	aria-label={label}
	style="--size: var(--font-size)px; --color: {lineColor};"
>
	<div class="inner">
		<span></span>
		<span></span>
		<span></span>
	</div>
</button>

<style src="./sidebar.css"></style>
