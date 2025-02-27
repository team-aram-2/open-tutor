<script lang="ts">
	import Hamburger from './hamburger.svelte';
	let selectedItem = 'view';
	let collapsed = false;

	const setSelectedItem = (item: string) => {
		selectedItem = item;
	};

	// Handle toggle command given by hamburger button
	const handleToggle = (event: CustomEvent<{ open: boolean }>) => {
		collapsed = event.detail.open; // Syncs with hamburger toggle's output
		console.log(collapsed);
	};
</script>

<div class="sidebar flex h-screen" style="border-top-right-radius: 25px;" class:collapsed>
	<div
		class="title-container"
		style="display: inline-flex; flex-direction: row; flex-wrap: nowrap; border-top-right-radius: 25px; width: 100%;"
	>
		<h2 class="sidebar-title">Student</h2>

		<div
			class="hamburger mt-3"
			style="align-self:stretch; display:flex; justify-content: flex-end; flex-direction: row; flex-grow: 1; border-top-right-radius: 25px; margin-right: 15px;"
		>
			<Hamburger open={collapsed} on:toggle={handleToggle}></Hamburger>
		</div>
	</div>

	<nav class="sidebar-items">
		<a
			href="/my_people/student"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'view'}
			on:click={() => setSelectedItem('view')}
		>
			<p>View Tutors</p>
		</a>
		<!-- TODO: MOVE THIS HREF BACK TO THE APPOINTMENTS <a> tag -->
		<a
			href="#/"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'apt'}
			on:click={() => setSelectedItem('apt')}
		>
			<p>Appointments</p>
		</a>
		<a
			href="/messages/student"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'msg'}
			on:click={() => setSelectedItem('msg')}
		>
			<p>Messages</p>
		</a>
		<a
			href="#/"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'pym'}
			on:click={() => setSelectedItem('pym')}
		>
			<p>Payments</p>
		</a>
	</nav>
</div>

<style>
	.sidebar {
		top: 0;
		bottom: 0;
		left: 0;

		padding-right: 0px;

		border-top-right-radius: 25px;
		border-bottom-right-radius: 25px;

		min-width: 375px;
		width: auto;

		display: flex;
		flex-direction: column;

		background-color: #453a69;

		z-index: 100;
	}
	.sidebar.collapsed {
		width: 90px;
		min-width: 90px;
	}
	.sidebar.collapsed .sidebar-title {
		display: none;
	}
	.sidebar.collapsed .sidebar-items a p {
		font-size: 0;
	}
	.sidebar.collapsed .sidebar-items a.selected-sidebar-item * {
		font-size: 0;
	}

	.sidebar-title {
		display: flex;
		flex-direction: column;

		position: relative;
		top: 0;
		left: 0;

		padding-top: 0;
		padding-left: 15px;
		margin-top: 15px;

		max-height: 50px;
		margin-bottom: 0;

		font-size: 70px;
		font-weight: bold;
		color: white;
	}

	.sidebar-items {
		display: flex;
		flex-direction: column;

		padding-left: 0px;
		padding-top: 50px;
		padding-bottom: 50px;

		margin: 0;

		position: relative;
		top: 0;

		font-size: 40px;
		font-weight: bold;
		line-height: 1.75em;
	}
	.sidebar-items p {
		padding: 15px 0px 15px 15px;
		/* padding: 15px; */
		color: white;
	}

	.selected-sidebar-item {
		background-color: #7261a8;
		padding-left: 20px;
		margin-right: 20px;

		border-bottom-right-radius: 20px;
		border-top-right-radius: 20px;
	}
	.no-decoration {
		text-decoration: none;
		color: inherit;
	}
</style>
