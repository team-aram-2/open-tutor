<script lang="ts">
	export let startAt = 0;
	export let endAt = 0;
</script>

<div class="meeting-card">
	<div class="time-date-info w-1/5"></div>
	<div class="tutor-info-block w-3/5"></div>
	<div class="link-info w-1/5"></div>
</div>

<style>
	.meeting-card {
		display: flex;
		flex-direction: row;
		width: calc(100% - 40px);
		height: calc(var(--font-size) * 5);

		margin: 10px 20px 0px 20px;

		background-color: var(--yellow-very-light);

		border: 5px solid var(--yellow-neutral);
		border-radius: 25px;
	}
	div {
		border: 5px solid blue;
	}
</style>
