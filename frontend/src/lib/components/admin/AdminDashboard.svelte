<!-- src/lib/components/admin/AdminDashboard.svelte -->
<script>
	import { goto } from '$app/navigation';
</script>

<div class="admin-dashboard">
	<h1>Admin Dashboard</h1>

	<div class="admin-tools">
		<button class="tool-card" on:click={() => goto('/admin/skills')}>
			<h2>Skills Management</h2>
			<p>Create, view, update, and delete tutor skills</p>
		</button>
		<!-- Add more admin tools here as needed -->
	</div>
</div>

<style>
	.admin-dashboard {
		padding: 2rem;
	}

	.admin-tools {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 1.5rem;
		margin-top: 2rem;
	}

	.tool-card {
		padding: 1.5rem;
		border: 1px solid #ddd;
		border-radius: 0.5rem;
		cursor: pointer;
		transition:
			transform 0.2s,
			box-shadow 0.2s;
	}

	.tool-card:hover {
		transform: translateY(-3px);
		box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
	}

	h1 {
		margin-bottom: 1.5rem;
	}

	h2 {
		margin-top: 0;
		margin-bottom: 0.5rem;
	}

	p {
		margin: 0;
		color: #666;
	}
</style>
