<script>
	import Tutor from '$lib/components/cards/tutor-card.svelte';
	import tutorsData from '$lib/mock/my_tutors_mock.json';
</script>

<div>
	<div class="cardcontainer">
		{#each tutorsData.tutors as tutor}
			<Tutor
				name="{tutor.firstName} {tutor.lastName}"
				rating={tutor.overallRating}
				skills={tutor.skills}
				userId={tutor.userId}
			></Tutor>
		{/each}
	</div>
	<p class="bottomText">
		Not what you're looking for?<a href="_#" class="bottomTextLink">Search again</a>
	</p>
	<!-- TODO: TURN THIS INTO AN ACTUAL LINK -->
</div>

<style>
	.cardcontainer {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		gap: 50px;
		flex: 0 0 33%;
		padding: 30px 30px 0 30px;
	}
	.bottomText {
		display: flex;
		width: 100%;
		font-weight: bold;
		font-size: var(--font-size);
		color: var(--yellow-very-light);
		justify-content: center;
	}
	.bottomTextLink {
		color: var(--yellow-neutral);
		margin-left: 0.25em;
	}
</style>
