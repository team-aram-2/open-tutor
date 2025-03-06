<script>
	import { createEventDispatcher } from 'svelte';

	// Export props for external customization.
	export let open = true; // Whether the menu is open (X) or closed (hamburger)
	export let size = 24; // Button size in pixels
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
	style="--size: {size}px; --color: {lineColor};"
>
	<div class="inner">
		<span></span>
		<span></span>
		<span></span>
	</div>
</button>

<style>
	.hamburger {
		display: flex;
		justify-content: center;
		align-items: center;

		width: calc(var(--size) + 30px);
		height: calc(var(--size) + 30px);
		margin: 0;
		padding: 0;

		background: var(--yellow-light);
		border-radius: calc(var(--size) + 30px);
		cursor: pointer;
	}
	.hamburger div {
		display: inline-flex;
		flex-direction: column;
		justify-content: space-around;

		width: var(--size);
		height: var(--size);
		margin: 10px;
		padding: 0;
	}
	.hamburger span {
		display: block;
		width: 100%;
		height: 5px;
		background: var(--color);
		border-radius: 5px;
		transition:
			transform 0.3s ease,
			opacity 0.3s ease,
			width 0.3s ease,
			margin 0.3s ease;
	}
	/* Transform first line to form the top part of the X */
	.hamburger.open span:nth-child(1) {
		transform: translateY(8px) rotate(45deg) scale(1.1, 1.1);
	}
	/* Hide the middle line */
	.hamburger.open span:nth-child(2) {
		/* opacity: 0; */
		margin-left: calc(var(--size) / 2);
		width: 0;
	}
	/* Transform third line to form the bottom part of the X */
	.hamburger.open span:nth-child(3) {
		transform: translateY(-8px) rotate(-45deg) scale(1.1, 1.1);
	}
</style>
