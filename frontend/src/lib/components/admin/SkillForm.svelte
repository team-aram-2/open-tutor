<!-- src/lib/components/admin/skills/SkillForm.svelte -->
<script>
	import { createEventDispatcher } from 'svelte';

	export let skill = { title: '', description: '', category_id: '' };
	export let categories = [];

	const dispatch = createEventDispatcher();

	let localSkill = { ...skill };
	let errors = {};

	function validateForm() {
		errors = {};

		if (!localSkill.title.trim()) {
			errors.title = 'Title is required';
		}

		if (!localSkill.description.trim()) {
			errors.description = 'Description is required';
		}

		if (!localSkill.category_id) {
			errors.category_id = 'Category is required';
		}

		return Object.keys(errors).length === 0;
	}

	function handleSubmit() {
		if (validateForm()) {
			dispatch('submit', localSkill);
		}
	}

	function handleCancel() {
		dispatch('cancel');
	}
</script>

<div class="skill-form">
	<h2>{skill.id ? 'Edit' : 'Create'} Skill</h2>

	<div class="form-field">
		<label for="title">Title</label>
		<input type="text" id="title" bind:value={localSkill.title} class:error={errors.title} />
		{#if errors.title}
			<span class="error-text">{errors.title}</span>
		{/if}
	</div>

	<div class="form-field">
		<label for="description">Description</label>
		<textarea
			id="description"
			bind:value={localSkill.description}
			rows="5"
			class:error={errors.description}
		></textarea>
		{#if errors.description}
			<span class="error-text">{errors.description}</span>
		{/if}
	</div>

	<div class="form-field">
		<label for="category">Category</label>
		<select id="category" bind:value={localSkill.category_id} class:error={errors.category_id}>
			<option value="" disabled>Select a category</option>
			{#each categories as category}
				<option value={category.id}>{category.name}</option>
			{/each}
		</select>
		{#if errors.category_id}
			<span class="error-text">{errors.category_id}</span>
		{/if}
	</div>

	<div class="form-actions">
		<button type="button" class="cancel-button" on:click={handleCancel}>Cancel</button>
		<button type="button" class="submit-button" on:click={handleSubmit}>
			{skill.id ? 'Update' : 'Create'} Skill
		</button>
	</div>
</div>

<style>
	.skill-form {
		background-color: #f9f9f9;
		padding: 1.5rem;
		border-radius: 0.5rem;
		margin-bottom: 2rem;
		border: 1px solid #eee;
	}

	h2 {
		margin-top: 0;
		margin-bottom: 1.5rem;
	}

	.form-field {
		margin-bottom: 1rem;
	}

	label {
		display: block;
		margin-bottom: 0.5rem;
		font-weight: bold;
	}

	input,
	textarea,
	select {
		width: 100%;
		padding: 0.75rem;
		border: 1px solid #ddd;
		border-radius: 0.25rem;
		font-size: 1rem;
	}

	input.error,
	textarea.error,
	select.error {
		border-color: #f44336;
	}

	.error-text {
		color: #f44336;
		font-size: 0.875rem;
		margin-top: 0.25rem;
		display: block;
	}

	.form-actions {
		display: flex;
		justify-content: flex-end;
		gap: 1rem;
		margin-top: 1.5rem;
	}

	.submit-button {
		background-color: #4caf50;
		color: white;
		border: none;
		padding: 0.75rem 1.5rem;
		border-radius: 0.25rem;
		cursor: pointer;
		font-weight: bold;
	}

	.cancel-button {
		background-color: #f5f5f5;
		color: #333;
		border: 1px solid #ddd;
		padding: 0.75rem 1.5rem;
		border-radius: 0.25rem;
		cursor: pointer;
	}
</style>
