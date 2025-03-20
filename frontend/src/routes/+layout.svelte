<script>
	// Shared logic across pages goes here
	import { autoLogin } from '$lib/scripts/auth';
	import '../app.css';
	import Sidebar from '$lib/components/dashboard/sidebar.svelte';

	import { font_size, sidebar_width } from '$lib/stores';

	import { onMount } from 'svelte';

	onMount(() => {
		font_size.subscribe((value) => {
			// Set global font size
			document.documentElement.style.setProperty('--font-size', String(value));
		});
		autoLogin();
	});
</script>

<!-- Page Content -->
<div class="dashboard-layout">
	<div class="sidebar-container" style="width: min-content;">
		<Sidebar></Sidebar>
	</div>

	<!-- Where page content will be rendered -->
	<slot />
</div>

<style>
	/* Apply these globally */
	* {
		font-family: 'Inter', sans-serif;
	}

	.dashboard-layout {
		display: flex;
		flex-direction: row;
		position: fixed;
		flex-wrap: nowrap;

		overflow: hidden;

		top: 0;
		left: 0;

		min-height: 100;
		height: 100vh;
		width: 100vw;

		background: #231d34;
	}
</style>
