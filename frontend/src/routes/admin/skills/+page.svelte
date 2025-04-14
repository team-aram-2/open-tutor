<!-- src/routes/admin/skills/+page.svelte -->
<script>
	import SkillsManager from '$lib/components/admin/skills/SkillsManager.svelte';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { user_role_mask, user_id } from '$lib/stores';
	import { hasRole, Role } from '$lib/scripts/auth';

	let isAuthenticated = false;
	let loading = true;

	onMount(async () => {
		// Check if user is authenticated as admin
		console.log($user_role_mask, Role.Admin);
		if (!hasRole($user_role_mask, Role.Admin)) {
			// Redirect to login page
			console.log('You do not have permission to access this resource.');
			goto('/?redirect=/admin/skills');
			return;
		}
	});
</script>

<div class="admin-skills">
	{#if loading}
		<div class="loading-container">
			<p>Loading...</p>
		</div>
	{:else if isAuthenticated}
		<div class="breadcrumb">
			<a href="/admin">Admin Dashboard</a> &gt; Skills Management
		</div>
		<SkillsManager />
	{:else}
		<div class="unauthorized">
			<h1>Unauthorized</h1>
			<p>You do not have permission to access this page.</p>
			<a href="/login?redirect=/admin/skills">Go to Login</a>
		</div>
	{/if}
</div>

<style>
	.admin-skills {
		min-height: 100vh;
	}

	.loading-container {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100vh;
	}

	.unauthorized {
		text-align: center;
		padding: 3rem;
	}

	.breadcrumb {
		padding: 1rem 2rem;
		background-color: #f5f5f5;
		border-bottom: 1px solid #ddd;
	}

	.breadcrumb a {
		color: #2196f3;
		text-decoration: none;
	}

	.breadcrumb a:hover {
		text-decoration: underline;
	}
</style>
