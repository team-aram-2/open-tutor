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

<style>
	.hamburger {
		display: flex;
		justify-content: center;
		align-items: center;

		width: calc(var(--font-size) * 1.75);
		height: calc(var(--font-size) * 1.75);
		margin: 0;
		padding: 0;

		background: var(--yellow-light);
		border-radius: calc(var(--font-size) * 1.75);
		cursor: pointer;
	}
	.hamburger div {
		display: inline-flex;
		flex-direction: column;
		gap: calc(var(--font-size) * 0.01);
		justify-content: space-around;

		width: var(--font-size);
		height: var(--font-size);
		padding: 0;
		margin: 0px;
	}
	.hamburger span {
		display: block;
		width: 100%;
		height: calc(var(--font-size) / 5);
		background: var(--color);
		border-radius: calc(var(--font-size) / 5);
		transition:
			transform 0.3s ease,
			opacity 0.3s ease,
			width 0.3s ease,
			margin 0.3s ease;
	}
	/* Transform first line to form the top part of the X */
	.hamburger.open span:nth-child(1) {
		transform: translateY(calc(var(--font-size) / 3)) rotate(45deg) scale(1.1, 1.1);
	}
	/* Hide the middle line */
	.hamburger.open span:nth-child(2) {
		/* opacity: 0; */
		margin-left: calc(var(--font-size) / 2);
		width: 0;
	}
	/* Transform third line to form the bottom part of the X */
	.hamburger.open span:nth-child(3) {
		transform: translateY(calc(-1 * var(--font-size) / 3)) rotate(-45deg) scale(1.1, 1.1);
	}
</style>
