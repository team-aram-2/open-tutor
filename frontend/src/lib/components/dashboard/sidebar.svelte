<script lang="ts">
	import Hamburger from './hamburger.svelte';
	import CalendarIcon from './sidebar-icons/calendar_icon.svelte';
	import SpeechBubbleIcon from './sidebar-icons/speechBubble_icon.svelte';
	import CreditCardIcon from './sidebar-icons/creditCard_icon.svelte';
	import GearIcon from './sidebar-icons/gear_icon.svelte';
	import PersonHeadIcon from './sidebar-icons/personHead_icon.svelte';
	// import { resizeObserver } from '$lib/scripts/observers';

	import { sidebar_width } from '$lib/stores';
	import { onMount } from 'svelte';

	let selectedItem = 'view';
	let collapsed = false;
	import { logged_in } from '$lib/stores';

	let targetSidebar: HTMLDivElement | null = null;

	const setSelectedItem = (item: string) => {
		selectedItem = item;
	};

	// Handle toggle command given by hamburger button
	const handleToggle = (event: CustomEvent<{ open: boolean }>) => {
		collapsed = event.detail.open; // Syncs with hamburger toggle's output
		console.log(collapsed);
	};

	onMount(() => {
		if (!window || !targetSidebar) return;

		let observer = new ResizeObserver((entries) => {
			for (const entry of entries) {
				sidebar_width.set(entry.contentRect.width);
			}
		});

		observer.observe(targetSidebar);

		return () => observer.disconnect();
	});
</script>

<!-- bind:this={targetSidebar} -->
<div class="sidebar flex h-screen" style="border-top-right-radius: 25px;" class:collapsed>
	<!-- Sidebar Title -->
	<div class="title-container">
		<h2 class="sidebar-title">Student</h2>

		<div class="hamburger">
			<Hamburger open={collapsed} on:toggle={handleToggle}></Hamburger>
		</div>
	</div>

	<!-- Sidebar Items -->
	<nav class="sidebar-items">
		{#if $logged_in}
			<a
				href="/my_people/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'view'}
				on:click={() => setSelectedItem('view')}
			>
				<p class="sidebar-item-text">View Tutors</p>

				<!-- Icon that appears when sidebar is collapsed -->
				<div class="collapsed-sidebar-item-icon">
					<!-- Person Head -->
					<PersonHeadIcon />
				</div>
			</a>
			<!-- TODO: MOVE THIS HREF BACK TO THE APPOINTMENTS <a> tag -->
			<a
				href="/meetings/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'mtgs'}
				on:click={() => setSelectedItem('mtgs')}
			>
				<p class="sidebar-item-text">Meetings</p>

				<!-- Icon that appears when sidebar is collapsed -->
				<div class="collapsed-sidebar-item-icon">
					<!-- Calendar -->
					<CalendarIcon />
				</div>
			</a>
			<a
				href="/messages/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'msg'}
				on:click={() => setSelectedItem('msg')}
			>
				<p class="sidebar-item-text">Messages</p>

				<!-- Icon that appears when sidebar is collapsed -->
				<div class="collapsed-sidebar-item-icon">
					<!-- Speech Bubble -->
					<SpeechBubbleIcon />
				</div>
			</a>
			<a
				href="#/"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'pym'}
				on:click={() => setSelectedItem('pym')}
			>
				<p class="sidebar-item-text">Payments</p>

				<!-- Icon that appears when sidebar is collapsed -->
				<div class="collapsed-sidebar-item-icon">
					<!-- Credit Card -->
					<CreditCardIcon />
				</div>
			</a>
		{:else}
			<a
				href="/login"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'login'}
				on:click={() => setSelectedItem('login')}
			>
				<p class="sidebar-item-text">Login</p>

				<!-- Icon that appears when sidebar is collapsed -->
				<div class="collapsed-sidebar-item-icon">
					<!--TODO ADD LOGIN SIDEBAR ITEM-->
				</div>
			</a>
		{/if}
		<a
			href="/settings_page"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'set'}
			on:click={() => setSelectedItem('set')}
		>
			<p class="sidebar-item-text">Settings</p>

			<!-- Icon that appears when sidebar is collapsed -->
			<div class="collapsed-sidebar-item-icon">
				<!-- Gear -->
				<GearIcon />
			</div>
		</a>
	</nav>
</div>

<style>
	@import './sidebar.css';
</style>
