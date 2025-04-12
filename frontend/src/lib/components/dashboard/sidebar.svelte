<script lang="ts">
	import Hamburger from './hamburger.svelte';
	import CalendarIcon from './sidebar-icons/calendar_icon.svelte';
	import SpeechBubbleIcon from './sidebar-icons/speechBubble_icon.svelte';
	import CreditCardIcon from './sidebar-icons/creditCard_icon.svelte';
	import GearIcon from './sidebar-icons/gear_icon.svelte';
	import PersonHeadIcon from './sidebar-icons/personHead_icon.svelte';
	import { logged_in, sidebar_width } from '$lib/stores';
	import { onMount } from 'svelte';

	let selectedItem = 'view';
	let collapsed = false;
	let resizeObserver: ResizeObserver;
	let sidebarRef: HTMLElement | null = null;

	const setSelectedItem = (item: string) => {
		selectedItem = item;
	};

	// Handle toggle command given by hamburger button
	const handleToggle = (event: CustomEvent<{ open: boolean }>) => {
		collapsed = event.detail.open; // Syncs with hamburger toggle's output
		console.log(collapsed);
	};

	onMount(() => {
		// Initialize ResizeObserver to observe changes in size
		// For all entries, update sidebar_width with entry.contentRect.width
		resizeObserver = new ResizeObserver((entries) => {
			for (let entry of entries) {
				sidebar_width.set(String(entry.contentRect.width) + 'px');
			}
		});

		// Begin observing the element referenced by elementref
		if (sidebarRef) {
			resizeObserver.observe(sidebarRef);
		}

		return () => {
			// Clean up observer when sidebar is destroyed
			if (resizeObserver && sidebarRef) {
				resizeObserver.unobserve(sidebarRef);
			}
			if (resizeObserver) {
				resizeObserver.disconnect();
			}
		};
	});
</script>

<aside
	class="sidebar flex h-screen"
	style="border-top-right-radius: 25px;"
	class:collapsed
	aria-label="Sidebar"
	bind:this={sidebarRef}
>
	<!-- Sidebar Title -->
	<div class="title-container">
		<h2 class="sidebar-title">Student</h2>

		<div class="hamburger">
			<Hamburger open={collapsed} on:toggle={handleToggle}></Hamburger>
		</div>
	</div>

	<!-- Sidebar Items -->
	<nav class="sidebar-items" aria-label="Sidebar Navigation">
		{#if $logged_in}
			<a
				href="/my_people/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'view'}
				on:click={() => setSelectedItem('view')}
				aria-label="View Tutors"
				aria-current={selectedItem === 'view' ? 'page' : undefined}
			>
				<p class="sidebar-item-text">View Tutors</p>
				<div class="collapsed-sidebar-item-icon">
					<!-- Person Head -->
					<PersonHeadIcon />
				</div>
			</a>
			<!-- TODO: MOVE THIS HREF BACK TO THE APPOINTMENTS <a> tag -->
			<a
				href="/meetings/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'mtng'}
				on:click={() => setSelectedItem('mtng')}
				aria-label="Meetings"
				aria-current={selectedItem === 'mtng' ? 'page' : undefined}
			>
				<p class="sidebar-item-text">Meetings</p>
				<div class="collapsed-sidebar-item-icon">
					<CalendarIcon />
				</div>
			</a>
			<a
				href="/messages/student"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'msg'}
				on:click={() => setSelectedItem('msg')}
				aria-label="Messages"
				aria-current={selectedItem === 'msg' ? 'page' : undefined}
			>
				<p class="sidebar-item-text">Messages</p>
				<div class="collapsed-sidebar-item-icon">
					<SpeechBubbleIcon />
				</div>
			</a>
			<a
				href="#/"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'pym'}
				on:click={() => setSelectedItem('pym')}
				aria-label="Payments"
				aria-current={selectedItem === 'pym' ? 'page' : undefined}
			>
				<p class="sidebar-item-text">Payments</p>
				<div class="collapsed-sidebar-item-icon">
					<CreditCardIcon />
				</div>
			</a>
			<a
				href="/settings_page"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'set'}
				on:click={() => setSelectedItem('set')}
				aria-label="Settings"
				aria-current={selectedItem === 'set' ? 'page' : undefined}
			>
				<p class="sidebar-item-text">Settings</p>
				<div class="collapsed-sidebar-item-icon">
					<GearIcon />
				</div>
			</a>
		{:else}
			<a
				href="/login"
				class="no-decoration"
				class:selected-sidebar-item={selectedItem === 'log'}
				on:click={() => setSelectedItem('log')}
				aria-label="Login"
				aria-current={selectedItem === 'log' ? 'page' : undefined}
			>
				<p>Login</p>
			</a>
		{/if}
	</nav>
</aside>

<style>
	.sidebar {
		display: flex;
		flex-direction: column;

		top: 0;
		bottom: 0;
		left: 0;

		padding-right: 0px;

		border-bottom-right-radius: calc(min(25px, (var(--font-size) / 1.25)));
		border-top-right-radius: calc(min(25px, (var(--font-size) / 1.25)));

		width: calc(11 * var(--font-size));
		max-width: 100%;

		background-color: #453a69;

		z-index: 100;
	}
	.sidebar.collapsed {
		width: calc(4 * var(--font-size));
	}

	.sidebar.collapsed .title-container {
		justify-content: center;
	}

	.sidebar.collapsed .sidebar-title {
		display: none;
		margin: 0;
		padding: 0;
	}
	.sidebar.collapsed .sidebar-item-text {
		display: none;
	}
	.sidebar.collapsed .collapsed-sidebar-item-icon {
		display: block;
		width: calc(var(--font-size) * 2.7);
		left: 0;
		margin-left: 7px;
	}
	.sidebar.collapsed .sidebar-items .sidebar.collapsed .sidebar-items a p {
		font-size: 0;
	}
	.sidebar.collapsed .sidebar-items a.selected-sidebar-item * {
		font-size: 0;
	}

	.title-container {
		display: flex;
		flex-direction: row;
		flex-wrap: nowrap;
		align-items: center;
		justify-content: space-between;

		width: 100%;
		height: calc(2.75 * var(--font-size));
		min-height: calc(var(--font-size) + 30px);

		border-top-right-radius: 25px;
		border-bottom: 3px solid transparent;
		border-image: linear-gradient(to right, var(--yellow-light), 99%, transparent) 1;
	}
	.sidebar-title {
		display: flex;
		flex-direction: row;

		position: relative;
		top: 0;
		left: 0;

		padding: 0;
		margin-left: calc(max(7px, (var(--font-size) / 8)));

		height: fit-content;
		margin-bottom: 0;

		font-size: calc(1.5 * var(--font-size));
		line-height: calc(1.5 * var(--font-size));
		font-weight: bold;
		color: white;
	}

	.sidebar-items {
		display: flex;
		flex-direction: column;

		overflow-y: scroll;

		height: auto;

		margin-left: 0px;

		position: relative;
		top: 0;
	}
	.sidebar-items p {
		margin-top: calc(0.75 * var(--font-size));
		margin-bottom: calc(0.75 * var(--font-size));
		margin-left: calc(min(20px, (0.75 * var(--font-size))));

		font-size: var(--font-size);
		font-weight: bold;
		line-height: var(--font-size);

		color: white;
	}

	.selected-sidebar-item {
		background-color: #7261a8;
		margin-right: calc(20px - max(0px, (20px - (var(--font-size) / 1.5))));

		border-bottom-right-radius: calc(min(20px, (var(--font-size) / 1.25)));
		border-top-right-radius: calc(min(20px, (var(--font-size) / 1.25)));
	}
	.selected-sidebar-item p {
		margin-left: 20px;
	}

	.collapsed-sidebar-item-icon {
		display: none;
	}

	.hamburger {
		display: flex;
		justify-content: center;
		flex-direction: column;

		height: 100%;
		margin-right: 10px;

		right: 0px;
		border-top-right-radius: 25px;
	}
	.no-decoration {
		text-decoration: none;
		color: inherit;
	}
</style>
