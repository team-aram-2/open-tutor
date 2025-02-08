<script lang="ts">
	enum Tab {
		Login = 'LOGIN',
		Register = 'REGISTER'
	}

	$: selectedTab = Tab.Login;
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-1">
		<h1 class="text-center text-5xl">Open Tutor</h1>

		<div class="my-6">
			<div class="flex flex-row h-10">
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab.Login ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab.Login)}>Login</button
				>
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab.Register ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab.Register)}>Register</button
				>
			</div>
			<hr class="border-gray-600 border-dashed" />
		</div>

		<div class="flex flex-col gap-3 w-120 mx-auto">
			{#if selectedTab === Tab.Login}
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="email" placeholder="Email" />
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="password" placeholder="Password" />
				<input class="rounded-md bg-blue-600 py-2 cursor-pointer" type="submit" value="Login" />
			{:else if selectedTab === Tab.Register}
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="text" placeholder="First name" />
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="text" placeholder="Last name" />
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="email" placeholder="Email" />
				<input class="bg-gray-900 px-4 py-2 rounded-md" type="password" placeholder="Password" />
				<input class="rounded-md bg-blue-600 py-2 cursor-pointer" type="submit" value="Register" />
			{/if}
		</div>
	</div>
</div>
