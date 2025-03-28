<script lang="ts">
	const categories = ['Professionalism', 'Knowledge', 'Communication', 'Punctuality', 'Overall'];
</script>

<div class="bg-[#fdf8ed] w-1/4 h-1/4 rounded-lg">
	<div class="flex flex-col m-4">
		{#each categories as cat}
			<div class="flex flex-row items-center gap-2">
				<span>{cat}</span>
				<div class="flex flex-row ml-auto">
					{#each Array(5) as num}
						<svg class="w-8 h-8" viewBox="0 0 51 48">
							<path d="M25,1 L31,17 L48,17 L35,29 L39,45 L25,36 L11,45 L15,29 L2,17 L19,17 Z" />
						</svg>
					{/each}
				</div>
			</div>
		{/each}
	</div>
</div>

<style>
</style>
