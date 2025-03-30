<script lang="ts">
	import { fetchSkillsJSON, loadSkills, handleSubmit } from './tutor-search/tutor-search';
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	import type { SkillsItem } from '$lib/types/types';
	import SearchIcon from './tutor-search/search-icon.svelte';
	import AscIcon from './tutor-search/order-icons/asc-icon.svelte';
	import DescIcon from './tutor-search/order-icons/desc-icon.svelte';
	import UnsortedIcon from './tutor-search/unsorted-icon.svelte';
	import RatingIcon from './tutor-search/sort-icons/rating-icon.svelte';
	import TutorNameIcon from './tutor-search/sort-icons/tutor-name-icon.svelte';
	import InteractiveRatingBar from '../generic/interactive-rating-bar.svelte';
	// pageSize       (integer, req)
	// pageIndex      (integer, req)
	// minRating      (number, optional)
	// selectedSkills (array, optional)
	let selectedSkills: string[] = [];
	let order = 'asc'; // (string, optional) (can be "asc" XOR "desc")
	let sort = 'rating'; // (string, optional) (can be "rating XOR name")

	let availableSkills: string[] = [
		'Aboard',
		'About',
		'Above',
		'Across',
		'After',
		'Against',
		'Along',
		'Among',
		'Around',
		'At',
		'Before',
		'Behind',
		'Below',
		'Beneath',
		'Beside',
		'Between',
		'Beyond',
		'But',
		'By',
		'Down',
		'During',
		'Despite',
		'Except',
		'For',
		'From',
		'In',
		'Inside'
	];

	availableSkills;

	// On update of one of the selected skills, update the list
	function toggleSkill(skillName: string) {
		// TODO: Check that the skill is valid
		if (!availableSkills.includes(skillName)) {
			console.error('Invalid skill');
			return;
		}

		// If selectedSkills is already selected, filter it out of selectedSkills. Otherwise, add it to selectedSkills
		if (selectedSkills.includes(skillName)) {
			selectedSkills = selectedSkills.filter((value) => {
				return !(value == skillName);
			});
		} else {
			selectedSkills = [...selectedSkills, skillName];
		}
	}

	function toggleSort() {
		if (sort == 'rating') {
			sort = 'name';
		} else {
			sort = 'rating';
		}
	}
	function toggleOrder() {
		if (order == 'asc') {
			order = 'desc';
		} else if (order == 'desc') {
			order = 'unsorted';
		} else {
			order = 'asc';
		}
	}

	onMount(async () => {
		console.log('onmount started');
		try {
			// const skills: SkillsItem[] = await loadSkills();
			// availableSkills = skills.map((skill) => skill.title);
			console.log('Skills finished loading!');
		} catch (error) {
			console.error('error loading skills:', error);
		}
	});
</script>

<div class="search-container">
	<div class="skills-container">
		{#if availableSkills.length > 0}
			{#each availableSkills as skill}
				<button
					id={skill}
					class="skill"
					class:selected={selectedSkills.includes(skill)}
					on:click={() => {
						toggleSkill(skill);
					}}>{skill}</button
				>
			{/each}
		{:else}
			<p>Loading skills...</p>
		{/if}
	</div>
	<div class="sorting-container">
		{#if sort == 'name'}
			<div class="tutor-name-lookup">
				<label for="tutor-name-search-box" class="tutor-name-search-box-label"
					>Search By Tutor Name</label
				>
				<textarea
					class="tutor-name-search-box"
					name="tutor-name-search-box"
					maxlength="50"
					placeholder="Enter Name..."
				></textarea>
			</div>
		{:else if sort == 'rating'}
			<div class="tutor-rating-lookup">
				<label for="tutor-name-search-box" class="tutor-name-search-box-label"
					>Search By Tutor Rating</label
				>
				<div class="rating-bar" style="margin: 0px 10px 0px 10px;">
					<InteractiveRatingBar />
				</div>
			</div>
		{/if}

		<div class="sort-and-order-container">
			<button class="filter-by-toggle" on:click={toggleSort}>
				{#if sort == 'name'}
					<TutorNameIcon />
				{:else if sort == 'rating'}
					<RatingIcon />
				{/if}
			</button>
			<button class="sort-toggle" on:click={toggleOrder}>
				{#if order == 'asc'}
					<AscIcon />
				{:else if order == 'desc'}
					<DescIcon />
				{:else}
					<UnsortedIcon />
				{/if}
			</button>
		</div>
	</div>
	<button class="search-button">
		<search>
			<SearchIcon />
		</search>
	</button>
</div>

<style>
	@import './tutor-search/tutor-search.css';
</style>
