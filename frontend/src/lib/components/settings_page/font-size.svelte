<script lang="ts">
	import { fontSize } from '$lib/stores';

	// Dictionary with labels and sizes
	let sizeLabels = new Map<string, string>([
		['very-small', '10px'],
		['small', '25px'],
		['normal', '45px'],
		['large', '60px'],
		['very-large', '90px']
	]);

	function setFontSize(value: string) {
		console.log('Setting font size to ' + value);
		fontSize.set(value);
		console.log();
	}
</script>

<div class="font-size-buttons-container">
	<p
		class="text-size-indicator"
		style="font-size: calc(max(calc(var(--font-size)/2), 20px); padding-right: 0px;"
	>
		A
	</p>
	{#each [...sizeLabels] as [key, value]}
		<input
			type="radio"
			name="font-size"
			value={key}
			class="font-size-button rounded-2xl"
			on:click={() => {
				setFontSize(value);
			}}
		/>
	{/each}
	<p
		class="text-size-indicator"
		style="font-size: calc(max(var(--font-size), 45px); padding-left: 0px;"
	>
		A
	</p>
</div>

<style>
	.font-size-buttons-container {
		display: flex;
		flex-direction: row;
		gap: 5px;

		margin-right: 20px;

		justify-content: space-around;
		align-content: center;

		height: 100%;
		width: 40%;
	}
	.font-size-button {
		appearance: none;

		flex: 1;

		margin-top: 5px;
		height: calc(100% - 10px);
		min-width: 25px;

		background-color: var(--yellow-light);
	}

	.font-size-button:checked {
		background-color: var(--yellow-neutral);
		transition: 60ms background-color ease-in-out;
	}

	.text-size-indicator {
		flex: 1;

		height: calc(100% - 10px);

		line-height: 100%;
		align-content: center;
		text-align: center;

		font-weight: bolder;

		color: var(--purple-dark);

		padding: 0;
		margin-top: 5px;
		background-color: var(--purple-very-light);
		border-radius: 15px;
	}
</style>
