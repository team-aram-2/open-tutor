<script lang="ts">
	import { PUBLIC_API_HOST } from '$env/static/public';
	import { onMount } from 'svelte';

	enum Tab {
		_Login = 'LOGIN',
		_Register = 'REGISTER'
	}

	$: selectedTab = Tab._Login;
	$: submittingForm = false;
	$: submissionError = '';
	$: loginError = '';

	onMount(() => {
		const params = new URLSearchParams(window.location.search);
		loginError = params.get('err') ?? '';
	});
</script>

<div class="w-full flex flex-col text-white">
	<div class="flex flex-col gap-1 w-120 mx-auto">
		<h1 class="text-center text-5xl">Open Tutor</h1>

		<div class="my-6">
			<div class="flex flex-row h-10">
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab._Login ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab._Login)}>Login</button
				>
				<button
					class={`flex-grow transition-colors cursor-pointer ${
						selectedTab === Tab._Register ? 'bg-sky-700/40' : 'bg-white/0 hover:bg-white/20'
					}`}
					on:click={() => (selectedTab = Tab._Register)}>Register</button
				>
			</div>
			<hr class="border-gray-600 border-dashed" />
		</div>

		<span class="text-red-600 text-center">{loginError}</span>

		<form
			class="flex flex-col gap-3"
			method="POST"
			action="{PUBLIC_API_HOST}/auth/{selectedTab.toLowerCase()}"
		>
			{#if selectedTab === Tab._Login}
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="email"
					type="email"
					placeholder="Email"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="password"
					type="password"
					placeholder="Password"
				/>
				<div
					class="flex justify-left items-center gap-4 p-4 my-2 border border-black-200 rounded-sm dark:border-gray-700"
				>
					<input
						id="remember-login-checkbox"
						class="bg-gray-800 border border-gray-700 rounded-md shadow-inner
						checked:bg-blue-600 checked:border-blue-500 transition-all duration-200
						hover:ring-2 hover:ring-blue-400 focus:ring-2 focus:ring-blue-400 cursor-pointer"
						name="rememberLogin"
						type="checkbox"
					/>
					<label for="remember-login-checkbox" class="w-2/3">Remember login for 30 Days.</label>
					<input
						class="w-full rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50"
						type="submit"
						value="Login"
						disabled={submittingForm}
					/>
				</div>
			{:else if selectedTab === Tab._Register}
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="first_name"
					type="text"
					placeholder="First name"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="last_name"
					type="text"
					placeholder="Last name"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="email"
					type="email"
					placeholder="Email"
				/>
				<input
					class="bg-gray-900 px-4 py-2 rounded-md"
					name="password"
					type="password"
					placeholder="Password"
				/>
				<input
					class="rounded-md bg-blue-600 py-2 enabled:cursor-pointer disabled:opacity-50"
					type="submit"
					value="Register"
					disabled={submittingForm}
				/>
			{/if}
		</form>
		<span class="text-red-600 text-center">{submissionError}</span>
	</div>
</div>

<style>
	input[id='remember-login-checkbox'] {
		transform: scale(1.5);
		margin: 10px;
	}
</style>
