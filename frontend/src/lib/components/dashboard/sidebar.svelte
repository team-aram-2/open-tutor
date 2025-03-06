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
	<!-- Sidebar Title -->
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

	<!-- Sidebar Items -->
	<nav class="sidebar-items">
		<a
			href="/my_people/student"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'view'}
			on:click={() => setSelectedItem('view')}
		>
			<p class="sidebar-item-text">View Tutors</p>

			<!-- Icon that appears when sidebar is collapsed -->
			<div class="collapsed-sidebar-item-icon">
				<svg viewBox="0 0 100 100" width="100" height="100" fill="gray">
					<ellipse cx="50" cy="33" rx="20" ry="18" fill="#ccc" />
					<path d="M10 85 Q50 50 90 85" fill="#ccc" />
				</svg>
			</div>
		</a>
		<!-- TODO: MOVE THIS HREF BACK TO THE APPOINTMENTS <a> tag -->
		<a
			href="#/"
			class="no-decoration"
			class:selected-sidebar-item={selectedItem === 'apt'}
			on:click={() => setSelectedItem('apt')}
		>
			<p class="sidebar-item-text">Appointments</p>

			<!-- Icon that appears when sidebar is collapsed -->
			<div class="collapsed-sidebar-item-icon">
				<!-- Calendar -->
				<svg viewBox="0 0 100 100" width="100" height="100" fill="gray">
					<!-- Calendar base -->
					<rect x="12" y="27" width="76" height="59" rx="5" fill="#ccc" />
					<!-- Calendar rings -->
					<rect x="10" y="15" width="80" height="15" rx="4" fill="#888" />
					<!-- Number 1 -->
					<line
						x1="41"
						y1="47"
						x2="52"
						y2="40"
						stroke="#888"
						stroke-width="6"
						stroke-linecap="round"
					/>
					<line
						x1="52"
						y1="40"
						x2="52"
						y2="72"
						stroke="#888"
						stroke-width="6"
						stroke-linecap="round"
					/>
					<line
						x1="41"
						y1="72"
						x2="61"
						y2="72"
						stroke="#888"
						stroke-width="6"
						stroke-linecap="round"
					/>
				</svg>
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
				<svg viewBox="0 0 100 100" width="100" height="100" fill="gray">
					<!-- Bubble body-->
					<rect x="10" y="15" width="80" height="50" rx="10" fill="#ccc" />
					<!-- Bubble tail -->
					<polygon points="30,65 40,85 50,65" fill="#ccc" />
				</svg>
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
				<svg viewBox="0 0 100 100" width="100" height="100" fill="gray">
					<!-- Credit card base -->
					<rect x="10" y="15" width="80" height="68" rx="6" fill="#ccc" />
					<!-- Credit card stripe -->
					<rect x="10" y="30" width="80" height="18" fill="#888" />
				</svg>
			</div>
		</a>
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
				<svg viewBox="0 0 100 100" width="100" height="100">
					<g transform="translate(50, 50) scale(1.25) translate(-50, -50)">
						<!-- Gear body -->
						<circle cx="50" cy="50" r="27" fill="#ccc" />

						<!-- Gear teeth -->
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(78,50) rotate(0)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(69,69) rotate(45)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(50,78) rotate(90)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(31,69) rotate(135)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(22,50) rotate(180)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(31,31) rotate(225)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(50,22) rotate(270)"
						/>
						<rect
							x="-6"
							y="-6"
							width="12"
							height="12"
							rx="3"
							fill="#ccc"
							transform="translate(69,31) rotate(315)"
						/>
					</g>
				</svg>
			</div>
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
		width: 140px;
		min-width: 140px;
	}
	.sidebar.collapsed .sidebar-title {
		display: none;
	}
	.sidebar.collapsed .sidebar-item-text {
		display: none;
	}
	.sidebar.collapsed .collapsed-sidebar-item-icon {
		display: block;
		width: 100px;
		left: 0;
		margin-left: 7px;
	}
	.sidebar.collapsed .sidebar-items .sidebar.collapsed .sidebar-items a p {
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
		margin-right: 20px;

		border-bottom-right-radius: 20px;
		border-top-right-radius: 20px;
	}
	.selected-sidebar-item p {
		margin-left: 20px;
	}

	.collapsed-sidebar-item-icon {
		display: none;
	}

	.no-decoration {
		text-decoration: none;
		color: inherit;
	}
</style>
