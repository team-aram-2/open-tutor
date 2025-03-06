<script lang="ts">
	let value = 'normal';
</script>

<div class="font-size-buttons-container">
	<input type="radio" name="font-size" value="very-small" class="font-size-button" />
	<input type="radio" name="font-size" value="small" class="font-size-button" />
	<input type="radio" name="font-size" value="normal" class="font-size-button" />
	<input type="radio" name="font-size" value="large" class="font-size-button" />
	<input type="radio" name="font-size" value="very-large" class="font-size-button" />
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

		border-radius: 15px;
		background-color: var(--yellow-light);
	}

	.font-size-button:checked {
		background-color: var(--yellow-neutral);
		transition: 120ms background-color ease-in-out;
	}
</style>
