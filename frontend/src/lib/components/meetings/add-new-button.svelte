<script>
</script>

<div class="add-new-button">
	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
		<line
			x1="12"
			y1="5"
			x2="12"
			y2="19"
			stroke="var(--yellow-dark)"
			stroke-width="2"
			stroke-linecap="round"
		/>
		<line
			x1="5"
			y1="12"
			x2="19"
			y2="12"
			stroke="var(--yellow-dark)"
			stroke-width="2"
			stroke-linecap="round"
		/>
	</svg>
</div>

<style>
	.add-new-button {
		display: flex;
		align-items: center;
		justify-items: center;

		width: calc(2 * var(--font-size));
		height: calc(100% - 40px);

		border-radius: var(--font-size);

		background-color: var(--yellow-light);
	}
	:hover.add-new-button {
		transition: background-color 210ms ease-in-out;

		background-color: var(--purple-neutral);
	}
	.add-new-button:hover svg :is(path, line, polyline) {
		transition: stroke var(--purple-light) 200ms ease-in-out;

		stroke: var(--purple-very-light);
	}
</style>
