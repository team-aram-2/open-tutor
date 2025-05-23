<script lang="ts">
	import TutorSearch from '$lib/components/cards/tutor-search.svelte';
	import Tutor from '$lib/components/cards/tutor-card.svelte';
	// import tutorsData from '$lib/mock/my_tutors_mock.json';
	import { onMount } from 'svelte';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import type { TutorItem } from '$lib/types/types.ts';

	// While waiting on Alex to finish up his changes to the messages page, I
	// hacked this together cause I couldn't figure out how in the world to get the cards
	// to be of a uniform width while also adapting to the current font size. Bon appetit. -Caleb

	// Variables to automatically update the width of the cards to accommodate font size changes
	let container: HTMLDivElement;
	let containerWidth = 0;
	const gap = 10;
	let maxChildrenPerRow = 3;
	let childrenPerRow = 1;
	let cardWidth = 100;
	let tutorData: Array<TutorItem> = [];
	// Calculates minimum width for children based off of the global font size
	const calculateChildMinWidth = (): number => {
		// NOTE: minimum width of a card (to accommodate content) is as follows: calc(var(--font-size) * 12);
		const style = getComputedStyle(document.documentElement);
		return parseInt(style.getPropertyValue('--font-size').slice(0, -2), 10) * 11;
	};
	// Calculates number of cards that can fit per row in the flexbox (max 3 per row)
	const updateLayout = (): void => {
		if (container) {
			containerWidth = container.clientWidth;
			console.log('container width: ' + String(containerWidth));
			const childMinWidth = calculateChildMinWidth();

			// Calculate number of children per row
			const numOfChildren = Math.floor((containerWidth + gap) / (childMinWidth + gap));
			console.log('num of children: ' + String(numOfChildren));
			childrenPerRow = Math.min(maxChildrenPerRow, numOfChildren || 1);

			// Evenly distribute children along row
			cardWidth = (containerWidth - gap * (childrenPerRow - 1)) / childrenPerRow - 1;
		}
	};

	/**
	 * Fetch all tutors from the backend.
	 * This endpoint requires query parameters: pageSize and pageIndex.
	 * Adjust these values as needed.
	 *
	 * @returns A promise resolving to the JSON containing the list of tutors.
	 */
	async function fetchTutors(): Promise<TutorItem[] | null> {
		console.log('fetching tutors');
		try {
			const response = await fetch(`${PUBLIC_API_HOST}/tutor?pageSize=13&pageIndex=0`, {
				method: 'GET',
				headers: {},

				credentials: 'include'
			});

			if (!response.ok) {
				console.error(`HTTP error, status: ${response.status}`);
				throw new Error(`HTTP error, status: ${response.status}`);
			}

			const tutors = await response.json();
			return tutors; // The JSON data is stored in this variable.
		} catch (error) {
			console.error('Error fetching tutors:', error);
			return null;
		}
	}

	onMount(() => {
		// Recalculate width of cards every time that the window size changes
		window.addEventListener('resize', updateLayout);

		fetchTutors().then((data) => {
			if (data) {
				console.log(data);
				tutorData = structuredClone(data);
				console.log('All tutors:');
				console.log(tutorData);
				tutorData = tutorData;
			}
		});
		updateLayout();

		return () => window.removeEventListener('resize', updateLayout);
	});
</script>

<div style="height: 100%; overflow-y: scroll;">
	<!-- Tutor Search Bar -->
	<TutorSearch />
	<div bind:this={container} class="cardcontainer" style="gap: {gap}px; row-gap: {gap * 1.125}px;">
		{#each tutorData as tutor}
			<Tutor
				name="{tutor.info.firstName} {tutor.info.lastName}"
				rating={5}
				skills={tutor.skills && tutor.skills.length > 0 ? tutor.skills : ['No skills registered']}
				userId={tutor.info.userId}
				width={String(cardWidth)}
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
		flex-wrap: wrap;
		flex-direction: row;
		align-content: flex-start;

		flex: 1;
		height: auto;

		margin-top: 10px;
		margin-left: 10px;
		margin-right: 10px;
	}

	.bottomText {
		display: flex;
		width: 100%;
		height: auto;
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
