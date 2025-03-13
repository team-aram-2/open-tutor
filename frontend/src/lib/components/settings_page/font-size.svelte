<script lang="ts">
	import { fontSize } from '$lib/stores';
	import { get } from 'svelte/store';
	let size = get(fontSize);

	function setFontSize(value: string) {
		console.log('Setting font size to ' + value);
		fontSize.set(value);
		console.log();
	}

	let value = 'normal';
</script>

<div class="font-size-buttons-container">
	<p
		class="text-size-indicator"
		style="font-size: calc(max(calc(var(--font-size)/2), 20px); padding-right: 0px;"
	>
		A
	</p>
	<input
		type="radio"
		name="font-size"
		value="very-small"
		class="font-size-button rounded-l-2xl rounded-r-sm"
		on:click={() => {
			setFontSize('10px');
		}}
	/>
	<input
		type="radio"
		name="font-size"
		value="small"
		class="font-size-button rounded-sm"
		on:click={() => {
			setFontSize('25px');
		}}
	/>
	<input
		type="radio"
		name="font-size"
		value="normal"
		class="font-size-button rounded-sm"
		on:click={() => {
			setFontSize('45px');
		}}
	/>
	<input
		type="radio"
		name="font-size"
		value="large"
		class="font-size-button rounded-sm"
		on:click={() => {
			setFontSize('60px');
		}}
	/>
	<input
		type="radio"
		name="font-size"
		value="very-large"
		class="font-size-button rounded-r-2xl rounded-l-sm"
		on:click={() => {
			setFontSize('90px');
		}}
	/>
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
