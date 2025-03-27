<!-- src/routes/+layout.svelte -->

<!-- This will contain all shared layouts across the site -->

<script>
	// Shared logic across pages goes here
	import '../app.css';
	import Sidebar from '$lib/components/dashboard/sidebar.svelte';
	import FPSidebar from '$lib/components/dashboard/front-page-sidebar.svelte'

	import { onMount } from 'svelte';

	onMount(() => {});

	let isLoggedIn = false;
  
	// Simulate login action
	function toggleLogin() {
		isLoggedIn = !isLoggedIn;
	}
</script>

<!-- Page Content -->
<div class="dashboard-layout">
	<FPSidebar></FPSidebar>

	<main class="w-full">
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

		/* font:  */
	}

	/* * .content {
		padding-left: 0px;
		padding-top: 0px;
		width: auto;
		flex-grow: 1;
	} */

	/* .page-title {
        padding-left: 30px;
    } */
</style>
