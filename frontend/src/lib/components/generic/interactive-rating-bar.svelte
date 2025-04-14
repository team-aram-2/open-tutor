<script lang="ts">
	const starLabels = [1, 2, 3, 4, 5];

	// function setRating(newRating: number) {} TODO: IMPLEMENT SENDING SEARCH INFO TO BACKEND
	function setRating() {}
</script>

<div class="star-container">
	{#each starLabels as star}
		<button
			on:click={() => {
				// setRating(star); TODO: IMPLEMENT SENDING SEARCH INFO TO BACKEND
				setRating();
			}}
			class="star-button"
			id={String(star)}
		>
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="auto" height="100%">
				<g transform="translate(2,2) scale(0.2)">
					<path
						class="star"
						d="M50 5 L63 40.9 L100 40.9 L69.5 63.1 L82.5 99 L50 76.8 L17.5 99 L30.5 63.1 L0 40.9 L37 40.9 Z"
					/>
				</g>
			</svg>
		</button>
	{/each}
</div>

<style>
	.star-container {
		display: flex;
		flex-direction: row;
		flex: 1 1 50%;
		justify-content: center;

		width: 100%;
	}

	.star {
		fill: var(--yellow-neutral);
		stroke: var(--yellow-very-dark);
		stroke-width: 3px;
	}
	.star-button {
		width: 20%;
		max-height: calc(2 * var(--font-size));
	}
	/* .star :hover {
		stroke: var(--purple-neutral);
	} */
</style>
