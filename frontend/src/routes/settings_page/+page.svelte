<script lang="ts">
	import FontSize from '$lib/components/settings_page/font-size.svelte';
	import { PUBLIC_API_HOST } from '$env/static/public';
</script>

<div class="setting-container">
	<div class="setting-categories"></div>
	<div class="setting-body">
		<!-- Font size -->
		<div class="setting-entry min-h-[60px]">
			<div class="setting-label">Font Size</div>
			<FontSize></FontSize>
		</div>
		<div class="setting-entry">
			<div class="setting-label">another setting</div>
		</div>
	</div>
</div>

<style>
	.setting-container {
		display: flex;
		flex-direction: column;
		flex-wrap: nowrap;

		height: 100%;
		width: 100%;
	}
	.setting-categories {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;

		height: 7em;

		border-top-left-radius: 25px;
		border-top-right-radius: 25px;

		background-color: var(--purple-very-dark);
	}
	.setting-body {
		display: flex;
		flex-direction: column;

		height: 100%;
	}

	.setting-entry {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;

		align-items: center;

		width: 100%;
		height: calc(var(--font-size) * 2);

		border-bottom: 5px solid var(--purple-very-dark);
		border-left: 5px solid var(--purple-very-dark);
		border-radius: 15px;

		background-color: var(--purple-dark);
	}
	.setting-label {
		flex: 1;

		text-align: left;

		margin-left: 3%;
		padding: 0;

		font-size: var(--font-size);
		font-weight: bolder;
		color: var(--yellow-very-light);
	}
</style>
