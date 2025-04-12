<!-- src/routes/+layout.svelte -->

<!-- This will contain all shared layouts across the site -->

<script>
	// Shared logic across pages goes here
	import '../app.css';
	import Sidebar from '$lib/components/dashboard/sidebar.svelte';
	import { autoLogin } from '$lib/scripts/auth';

	import { onMount } from 'svelte';
	import { fontSize, sidebar_width } from '$lib/stores';

	// Load page settings
	onMount(() => {
		autoLogin();
		fontSize.subscribe((value) => {
			// Set global font size
			document.documentElement.style.setProperty('--font-size', String(value));
		});
	});
	console.log('Sidebar width: ' + $sidebar_width);
</script>

<!-- Page Content -->
<div class="dashboard-layout">
	<Sidebar></Sidebar>

	<main style="width: calc(100% - {$sidebar_width});">
		<slot />
		<!-- Where page content will be rendered -->
	</main>
</div>

<!-- Style for the page goes here -->
<style>
	/* Apply these globally */
	* {
		font-family: 'Inter', sans-serif;
	}

	.dashboard-layout {
		display: flex;
		flex-direction: row;
		position: fixed;

		top: 0;
		bottom: 0;
		left: 0;
		right: 0;

		min-height: 100;

		background: #231d34;
	}
</style>
