<!-- src/lib/components/admin/skills/SkillsManager.svelte -->
<script>
	import { onMount } from 'svelte';
	import SkillForm from '$lib/components/admin/SkillForm.svelte';

	let skills = [];
	let categories = [];
	let loading = true;
	let error = null;
	let editingSkill = null;
	let showForm = false;

	// Fetch all skills
	async function fetchSkills() {
		loading = true;
		try {
			const response = await fetch('/api/skills', {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('adminToken')}`
				}
			});

			if (!response.ok) {
				throw new Error(`Failed to fetch skills: ${response.status}`);
			}

			skills = await response.json();
		} catch (err) {
			error = err.message;
		} finally {
			loading = false;
		}
	}

	// Fetch categories for dropdown
	async function fetchCategories() {
		try {
			const response = await fetch('/api/categories', {
				headers: {
					Authorization: `Bearer ${localStorage.getItem('adminToken')}`
				}
			});

			if (!response.ok) {
				throw new Error(`Failed to fetch categories: ${response.status}`);
			}

			categories = await response.json();
		} catch (err) {
			console.error('Error fetching categories:', err);
		}
	}

	// Delete a skill
	async function deleteSkill(id) {
		if (!confirm('Are you sure you want to delete this skill?')) return;

		try {
			const response = await fetch(`/api/skills/${id}`, {
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${localStorage.getItem('adminToken')}`
				}
			});

			if (!response.ok) {
				throw new Error(`Failed to delete skill: ${response.status}`);
			}

			// Refresh the skills list
			fetchSkills();
		} catch (err) {
			error = err.message;
		}
	}

	// Edit a skill
	function editSkill(skill) {
		editingSkill = { ...skill };
		showForm = true;
	}

	// Create a new skill
	function createNewSkill() {
		editingSkill = {
			title: '',
			description: '',
			category_id: categories.length > 0 ? categories[0].id : ''
		};
		showForm = true;
	}

	// Handle form submission
	async function handleFormSubmit(event) {
		const formData = event.detail;

		try {
			let url = '/api/skills';
			let method = 'POST';

			if (formData.id) {
				url = `/api/skills/${formData.id}`;
				method = 'PUT';
			}

			const response = await fetch(url, {
				method,
				headers: {
					'Content-Type': 'application/json',
					Authorization: `Bearer ${localStorage.getItem('adminToken')}`
				},
				body: JSON.stringify(formData)
			});

			if (!response.ok) {
				throw new Error(`Failed to ${formData.id ? 'update' : 'create'} skill: ${response.status}`);
			}

			// Hide form and refresh skills
			showForm = false;
			editingSkill = null;
			fetchSkills();
		} catch (err) {
			error = err.message;
		}
	}

	// Cancel form
	function cancelForm() {
		showForm = false;
		editingSkill = null;
	}

	onMount(() => {
		fetchSkills();
		fetchCategories();
	});
</script>

<div class="skills-manager">
	<div class="header">
		<h1>Skills Management</h1>
		<button class="create-button" on:click={createNewSkill}>Create New Skill</button>
	</div>

	{#if error}
		<div class="error-message">
			{error}
		</div>
	{/if}

	{#if showForm}
		<SkillForm
			skill={editingSkill}
			{categories}
			on:submit={handleFormSubmit}
			on:cancel={cancelForm}
		/>
	{/if}

	{#if loading}
		<div class="loading">Loading skills...</div>
	{:else}
		<table class="skills-table">
			<thead>
				<tr>
					<th>Title</th>
					<th>Description</th>
					<th>Category</th>
					<th>Actions</th>
				</tr>
			</thead>
			<tbody>
				{#if skills.length === 0}
					<tr>
						<td colspan="4" class="no-skills">No skills found</td>
					</tr>
				{:else}
					{#each skills as skill}
						<tr>
							<td>{skill.title}</td>
							<td class="description-cell">{skill.description}</td>
							<td>
								{categories.find((c) => c.id === skill.category_id)?.name || 'Unknown'}
							</td>
							<td class="actions">
								<button class="edit-button" on:click={() => editSkill(skill)}>Edit</button>
								<button class="delete-button" on:click={() => deleteSkill(skill.id)}>Delete</button>
							</td>
						</tr>
					{/each}
				{/if}
			</tbody>
		</table>
	{/if}
</div>

<style>
	.skills-manager {
		padding: 2rem;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 2rem;
	}

	.create-button {
		background-color: #4caf50;
		color: white;
		border: none;
		padding: 0.75rem 1.5rem;
		border-radius: 0.25rem;
		cursor: pointer;
		font-weight: bold;
	}

	.error-message {
		background-color: #ffebee;
		color: #c62828;
		padding: 1rem;
		border-radius: 0.25rem;
		margin-bottom: 1rem;
	}

	.loading {
		padding: 2rem 0;
		text-align: center;
		color: #666;
	}

	.skills-table {
		width: 100%;
		border-collapse: collapse;
	}

	.skills-table th,
	.skills-table td {
		padding: 1rem;
		text-align: left;
		border-bottom: 1px solid #ddd;
	}

	.skills-table th {
		background-color: #f5f5f5;
		font-weight: bold;
	}

	.description-cell {
		max-width: 300px;
		white-space: nowrap;
		overflow: hidden;
		text-overflow: ellipsis;
	}

	.no-skills {
		text-align: center;
		color: #666;
		padding: 2rem 0;
	}

	.actions {
		display: flex;
		gap: 0.5rem;
	}

	.edit-button {
		background-color: #2196f3;
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 0.25rem;
		cursor: pointer;
	}

	.delete-button {
		background-color: #f44336;
		color: white;
		border: none;
		padding: 0.5rem 1rem;
		border-radius: 0.25rem;
		cursor: pointer;
	}
</style>
