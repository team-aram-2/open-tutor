<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import Star from '../star.svelte';

	const categories = ['Professionalism', 'Knowledge', 'Communication', 'Punctuality', 'Overall'];

	$: ratingSubmissions = {} as Record<string, number>;

	const setRating = (cat: string, starCount: number) => {
		ratingSubmissions[cat] = starCount;
	};

	const submitRatings = () => {
		const reqPromise = fetch(`${PUBLIC_API_HOST}/meeting`, {
			method: 'POST',
			body: JSON.stringify(ratingSubmissions),
			credentials: 'include'
		});
	};
</script>

<div class="bg-[#fdf8ed] w-1/4 rounded-lg select-none">
	<div class="flex flex-col gap-4 w-full">
		<!-- Title -->
		<span class="mt-4 text-center text-semibold text-2xl">Submit Rating</span>

		<!-- Ratings -->
		<div class="mx-4 flex flex-col">
			{#each categories as cat}
				<div class="flex flex-row items-center justify-center gap-2">
					<span class="">{cat}</span>
					<div class="flex flex-row ml-auto">
						{#each Array.from({ length: 5 }, (_, i) => i + 1) as starCount}
							<Star
								filled={starCount <= (ratingSubmissions[cat] || 0)}
								on:click={() => setRating(cat, starCount)}
							/>
						{/each}
					</div>
				</div>
			{/each}
		</div>

		<textarea
			class="mx-4 p-2 h-36 bg-gray-300 rounded-md shadow-md resize-none"
			placeholder="Enter comment (optional)"
		/>

		<!-- Submit -->
		<button
			class="mb-4 mx-auto rounded-md text-white bg-sky-500 px-3 py-2 enabled:cursor-pointer disabled:opacity-50"
			on:click={submitRatings}
		>
			Submit Ratings</button
		>
	</div>
</div>

<style>
</style>
