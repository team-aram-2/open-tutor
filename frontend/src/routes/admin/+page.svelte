<!-- src/routes/admin/+page.svelte -->
<script>
	import AdminDashboard from '$lib/components/admin/AdminDashboard.svelte';
	import { onDestroy, onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { user_role_mask, user_id } from '$lib/stores';
	import { hasRole, Role } from '$lib/scripts/auth';
	import { get } from 'svelte/store';

	let isAuthorized = false;
	let ready = false;
	let timedOut = false;

	const timeout = setTimeout(() => {
		timedOut = true;
		if (!ready) {
			console.warn('Timed out waiting for role mask. Redirecting...');
			goto('/?redirect=/admin');
		}
	}, 5000); // 5 seconds

	const unsubscribe = user_role_mask.subscribe((mask) => {
		if (mask === 0 || timedOut) return;

		clearTimeout(timeout);
		ready = true;

		if (!hasRole(mask, Role.Admin)) {
			console.log('Redirecting — not admin');
			goto('/?redirect=/admin');
		} else {
			isAuthorized = true;
		}
	});

	onDestroy(() => {
		unsubscribe();
		clearTimeout(timeout);
	});
</script>

{#if loading}
	<div class="loading-container">
		<p>Loading...</p>
	</div>
{:else if isAuthenticated}
	<AdminDashboard />
{:else}
	<div class="unauthorized">
		<h1>Unauthorized</h1>
		<p>You do not have permission to access this page.</p>
		<a href="/login?redirect=/admin">Go to Login</a>
	</div>
{/if}

<style>
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
</style>
