<script lang="ts">
	import Star from '../star.svelte';

	const categories = ['Professionalism', 'Knowledge', 'Communication', 'Punctuality', 'Overall'];
	let comment = '';
	export let submitRatings = (ratings: Record<string, number>, comment?: string) => {};
	export let onRatingSubmit = () => {};
	console.log(comment);

	$: ratingSubmissions = {} as Record<string, number>;

	$: canSubmit = categories.every((category) => Object.keys(ratingSubmissions).includes(category));
	$: submitting = false;
	const setRating = (cat: string, starCount: number) => {
		ratingSubmissions[cat] = starCount;
	};

	const submit = () => {
		if (!canSubmit) {
			return;
		}

		submitting = true;
		try {
			submitRatings(ratingSubmissions, comment);
			onRatingSubmit();
		} catch (err) {
			console.error(`Failed to run submitRatings: ${err}`);
		} finally {
			submitting = false;
		}
	};
</script>

<div class="absolute flex justify-center items-center w-full h-full bg-black/30 z-50">
	<div class="bg-[#fdf8ed] w-96 rounded-lg select-none">
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
				bind:value={comment}
			/>

			<!-- Submit -->
			<button
				class="mb-4 mx-auto rounded-md text-white bg-sky-500 px-3 py-2 enabled:cursor-pointer disabled:opacity-50"
				on:click={submit}
				disabled={!canSubmit}
			>
				Submit Ratings</button
			>
		</div>
	</div>
</div>

<style>
</style>
